// The module 'vscode' contains the VS Code extensibility API
import { workspace, window, commands, Uri, ExtensionContext } from 'vscode';
import { posix } from 'path';

import { handleOpenInVSCode } from './openinvscode';
import { codexReopen } from './codex';

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: ExtensionContext) {
	// This line of code will only be executed once when your extension is activated
	initialCheckCodexJSON(context);
	// Creating file watchers to watch for events on codex.json
	const fswatcher = workspace.createFileSystemWatcher("**/codex.json", false, false, false);

	fswatcher.onDidDelete(e => {
		commands.executeCommand('setContext', 'codex.configFileExists', false);
		context.workspaceState.update("configFileExists", false);
	});
	fswatcher.onDidCreate(e => {
		commands.executeCommand('setContext', 'codex.configFileExists', true);
		context.workspaceState.update("configFileExists", true);
	});
	fswatcher.onDidChange(e => initialCheckCodexJSON(context));

	// Check for codex.json when a new folder is opened
	workspace.onDidChangeWorkspaceFolders(async (e) => initialCheckCodexJSON(context));

	// run codex shell when terminal is opened
	window.onDidOpenTerminal(async (e) => {
		if (workspace.getConfiguration("codex").get("autoShellOnTerminal")
			&& e.name !== "CodexTerminal"
			&& context.workspaceState.get("configFileExists")
		) {
			await runInTerminal('codex shell', true);
		}
	});

	// open in vscode URI handler
	const handleVSCodeUri = window.registerUriHandler({ handleUri: handleOpenInVSCode });

	const codexAdd = commands.registerCommand('codex.add', async () => {
		const result = await window.showInputBox({
			value: '',
			placeHolder: 'Package to add to codex. E.g., python39',
		});
		await runInTerminal(`codex add ${result}`, false);
	});

	const codexRun = commands.registerCommand('codex.run', async () => {
		const items = await getCodexScripts();
		if (items.length > 0) {
			const result = await window.showQuickPick(items);
			await runInTerminal(`codex run ${result}`, true);
		} else {
			window.showInformationMessage("No scripts found in codex.json");
		}
	});

	const codexShell = commands.registerCommand('codex.khulnasoft.comell', async () => {
		// todo: add support for --config path to codex.json
		await runInTerminal('codex shell', true);
	});

	const codexRemove = commands.registerCommand('codex.remove', async () => {
		const items = await getCodexPackages();
		if (items.length > 0) {
			const result = await window.showQuickPick(items);
			await runInTerminal(`codex rm ${result}`, false);
		} else {
			window.showInformationMessage("No packages found in codex.json");
		}
	});

	const codexInit = commands.registerCommand('codex.init', async () => {
		await runInTerminal('codex init', false);
		commands.executeCommand('setContext', 'codex.configFileExists', true);
	});

	const codexInstall = commands.registerCommand('codex.install', async () => {
		await runInTerminal('codex install', true);
	});

	const codexUpdate = commands.registerCommand('codex.update', async () => {
		await runInTerminal('codex update', true);
	});

	const codexSearch = commands.registerCommand('codex.search', async () => {
		const result = await window.showInputBox({ placeHolder: "Name or a subset of a name of a package to search" });
		await runInTerminal(`codex search ${result}`, true);
	});

	const setupDevcontainer = commands.registerCommand('codex.setupDevContainer', async () => {
		await runInTerminal('codex generate devcontainer', true);
	});

	const generateDockerfile = commands.registerCommand('codex.generateDockerfile', async () => {
		await runInTerminal('codex generate dockerfile', true);
	});

	const reopen = commands.registerCommand('codex.reopen', async () => {
		await codexReopen();
	});

	context.subscriptions.push(reopen);
	context.subscriptions.push(codexAdd);
	context.subscriptions.push(codexRun);
	context.subscriptions.push(codexInit);
	context.subscriptions.push(codexInstall);
	context.subscriptions.push(codexSearch);
	context.subscriptions.push(codexUpdate);
	context.subscriptions.push(codexRemove);
	context.subscriptions.push(codexShell);
	context.subscriptions.push(setupDevcontainer);
	context.subscriptions.push(generateDockerfile);
	context.subscriptions.push(handleVSCodeUri);
}

async function initialCheckCodexJSON(context: ExtensionContext) {
	// check if there is a workspace folder open
	if (workspace.workspaceFolders) {
		const workspaceUri = workspace.workspaceFolders[0].uri;
		try {
			// check if the folder has codex.json in it
			await workspace.fs.stat(Uri.joinPath(workspaceUri, "codex.json"));
			// codex.json exists setcontext for codex commands to be available
			commands.executeCommand('setContext', 'codex.configFileExists', true);
			context.workspaceState.update("configFileExists", true);
		} catch (err) {
			console.log(err);
			// codex.json does not exist
			commands.executeCommand('setContext', 'codex.configFileExists', false);
			context.workspaceState.update("configFileExists", false);
			console.log("codex.json does not exist");
		}
	}
}

async function runInTerminal(cmd: string, showTerminal: boolean) {
	// check if a terminal is open
	if ((<any>window).terminals.length === 0) {
		const terminalName = 'CodexTerminal';
		const terminal = window.createTerminal({ name: terminalName });
		if (showTerminal) {
			terminal.show();
		}
		terminal.sendText(cmd, true);
	} else {
		// A terminal is open
		// run the given cmd in terminal
		await commands.executeCommand('workbench.action.terminal.sendSequence', {
			'text': `${cmd}\r\n`
		});
	}
}

async function getCodexScripts(): Promise<string[]> {
	try {
		if (!workspace.workspaceFolders) {
			window.showInformationMessage('No folder or workspace opened');
			return [];
		}
		const workspaceUri = workspace.workspaceFolders[0].uri;
		const codexJson = await readCodexJson(workspaceUri);
		return Object.keys(codexJson['shell']['scripts']);
	} catch (error) {
		console.error('Error processing codex.json - ', error);
		return [];
	}
}

async function getCodexPackages(): Promise<string[]> {
	try {
		if (!workspace.workspaceFolders) {
			window.showInformationMessage('No folder or workspace opened');
			return [];
		}
		const workspaceUri = workspace.workspaceFolders[0].uri;
		const codexJson = await readCodexJson(workspaceUri);
		return codexJson['packages'];
	} catch (error) {
		console.error('Error processing codex.json - ', error);
		return [];
	}
}

async function readCodexJson(workspaceUri: Uri) {
	const fileUri = workspaceUri.with({ path: posix.join(workspaceUri.path, 'codex.json') });
	const readData = await workspace.fs.readFile(fileUri);
	const readStr = Buffer.from(readData).toString('utf8');
	const codexJsonData = JSON.parse(readStr);
	return codexJsonData;
}

// This method is called when your extension is deactivated
export function deactivate() { }
