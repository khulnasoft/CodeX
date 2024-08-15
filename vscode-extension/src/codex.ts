import { window, workspace, commands, ProgressLocation, Uri, ConfigurationTarget, env } from 'vscode';
import { writeFile, open } from 'fs/promises';
import { spawn, spawnSync } from 'node:child_process';


interface Message {
    status: string
}

export async function codexReopen() {
  if (process.platform === 'win32') {
    const seeDocs = 'See Codex docs';
    const result = await window.showErrorMessage(
      'This feature is not supported on your platform. \
      Please open VSCode from inside codex shell in WSL using the CLI.', seeDocs
    );
    if (result === seeDocs) {
      env.openExternal(Uri.parse('https://www.khulnasoft/codex/docs/ide_configuration/vscode/#windows-setup'));
      return;
    }
  }
  await window.withProgress({
    location: ProgressLocation.Notification,
    title: "Setting up your Codex environment. Please don't close vscode.",
    cancellable: true
  },
    async (progress, token) => {
      token.onCancellationRequested(() => {
        console.log("User canceled the long running operation");
      });

      const p = new Promise<void>(async (resolve, reject) => {
        if (workspace.workspaceFolders) {
          const workingDir = workspace.workspaceFolders[0].uri;
          const dotcodex = Uri.joinPath(workingDir, '/.codex');
          await logToFile(dotcodex, 'Installing codex packages');
          progress.report({ message: 'Installing codex packages...', increment: 25 });
          await setupDotCodex(workingDir, dotcodex);
          
          // setup required vscode settings
          await logToFile(dotcodex, 'Updating VSCode configurations');
          progress.report({ message: 'Updating configurations...', increment: 50 });
          updateVSCodeConf();

          // Calling CLI to compute codex env
          await logToFile(dotcodex, 'Calling "codex integrate" to setup environment');
          progress.report({ message: 'Calling Codex to setup environment...', increment: 80 });
          // To use a custom compiled codex when testing, change this to an absolute path.
          const codex = 'codex';
          // run codex integrate and then close this window
          const debugModeFlag = workspace.getConfiguration("codex").get("enableDebugMode");
          let child = spawn(codex, ['integrate', 'vscode', '--debugmode='+debugModeFlag], {
            cwd: workingDir.path,
            stdio: [0, 1, 2, 'ipc']
          });
          // if CLI closes before sending "finished" message
          child.on('close', (code: number) => {
            console.log("child process closed with exit code:", code);
            logToFile(dotcodex, 'child process closed with exit code: ' + code);
            window.showErrorMessage("Failed to setup codex environment.");
            reject();
          });
          // send config path to CLI
          child.send({ configDir: workingDir.path });
          // handle CLI finishing the env and sending  "finished"
          child.on('message', function (msg: Message, handle) {
            if (msg.status === "finished") {
              progress.report({ message: 'Finished setting up! Reloading the window...', increment: 100 });
              resolve();
              commands.executeCommand("workbench.action.closeWindow");
            }
            else {
              console.log(msg);
              logToFile(dotcodex, 'Failed to setup codex environment.' + String(msg));
              window.showErrorMessage("Failed to setup codex environment.");
              reject();
            }
          });
        }
      });
      return p;
    }
  );
}

async function setupDotCodex(workingDir: Uri, dotcodex: Uri) {
  try {
    // check if .codex exists
    await workspace.fs.stat(dotcodex);
  } catch (error) {
    //.codex doesn't exist
    // running codex shellenv to create it
    spawnSync('codex', ['shellenv'], {
      cwd: workingDir.path
    });
  }
}

function updateVSCodeConf() {
  if (process.platform === 'darwin') {
    const shell = process.env["SHELL"] ?? "/bin/zsh";
    const shellArgsMap = (shellType: string) => {
      switch (shellType) {
        case "fish":
          // We special case fish here because fish's `fish_add_path` function
          // tends to prepend to PATH by default, hence sourcing the fish config after
          // vscode reopens in codex environment, overwrites codex packages and 
          // might cause confusion for users as to why their system installed packages
          // show up when they type for example `which go` as opposed to the packages
          // installed by codex.
          return ["--no-config"];
        default:
          return [];
      }
    };
    const shellTypeSlices = shell.split("/");
    const shellType = shellTypeSlices[shellTypeSlices.length - 1];
    shellArgsMap(shellType);
    const codexCompatibleShell = {
      "codexCompatibleShell": {
        "path": shell,
        "args": shellArgsMap(shellType)
      }
    };

    workspace.getConfiguration().update(
      'terminal.integrated.profiles.osx',
      codexCompatibleShell,
      ConfigurationTarget.Workspace
    );
    workspace.getConfiguration().update(
      'terminal.integrated.defaultProfile.osx',
      'codexCompatibleShell',
      ConfigurationTarget.Workspace
    );
  }
}

async function logToFile(dotCodexPath: Uri, message: string) {
  // only print to log file if debug mode config is set to true
  if (workspace.getConfiguration("codex").get("enableDebugMode")){
    try {   
      const logFilePath = Uri.joinPath(dotCodexPath, 'extension.log');
      const timestamp = new Date().toUTCString();
      const fileHandler = await open(logFilePath.fsPath, 'a');
      const logData = new Uint8Array(Buffer.from(`[${timestamp}] ${message}\n`));
      await writeFile(fileHandler, logData, {flag: 'a'} );
      await fileHandler.close();
    } catch (error) {
      console.log("failed to write to extension.log file");
      console.error(error);
    }
  }
}
