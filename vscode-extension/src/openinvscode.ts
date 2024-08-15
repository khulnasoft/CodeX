import * as os from 'os';
import * as which from 'which';
import fetch from 'node-fetch';
import { exec } from 'child_process';
import * as FormData from 'form-data';
import { Uri, commands, window } from 'vscode';
import { chmod, mkdir, open, writeFile } from 'fs/promises';

type VmInfo = {
  /* eslint-disable @typescript-eslint/naming-convention */
  vm_id: string;
  private_key: string;
  username: string;
  working_directory: string;
  /* eslint-enable @typescript-eslint/naming-convention */
};

export async function handleOpenInVSCode(uri: Uri) {
  const queryParams = new URLSearchParams(uri.query);

  if (queryParams.has('vm_id') && queryParams.has('token')) {
    //Not yet supported for windows + WSL - will be added in future
    if (os.platform() !== 'win32') {
      window.showInformationMessage('Setting up codex');
      // getting ssh keys
      try {
        console.debug('Calling getVMInfo...');
        const response = await getVMInfo(queryParams.get('token'), queryParams.get('vm_id'));
        const res = await response.json() as VmInfo;
        console.debug('getVMInfo response: ' + res);
        // set ssh config
        console.debug('Calling setupSSHConfig...');
        await setupSSHConfig(res.vm_id, res.private_key);

        // connect to remote vm
        console.debug('Calling connectToRemote...');
        connectToRemote(res.username, res.vm_id, res.working_directory);
      } catch (err: any) {
        console.error(err);
        window.showInformationMessage('Failed to setup codex remote connection.');
      }
    } else {
      window.showErrorMessage('This function is not yet supported in Windows operating system.');
    }
  } else {
    window.showErrorMessage('Error parsing information for remote environment.');
    console.debug(queryParams.toString());
  };
}

async function getVMInfo(token: string | null, vmId: string | null): Promise<any> {
  // send post request to gateway to get vm info and ssh keys
  const gatewayHost = 'https://api.codex.sh/g/vm_info';
  const data = new FormData();
  data.append("vm_id", vmId);
  console.debug("calling codex to get vm_info...");
  const response = await fetch(gatewayHost, {
    method: 'post',
    body: data,
    headers: {
      authorization: `Bearer ${token}`
    }
  });
  console.debug("API Call to api.codex.sh response: " + response);
  return response;
}

async function setupCodexLauncher(): Promise<any> {
  // download codex launcher script
  const gatewayHost = 'https://releases.khulnasoft/codex';
  const response = await fetch(gatewayHost, {
    method: 'get',
  });
  const launcherPath = `${process.env['HOME']}/.config/codex/launcher.sh`;

  try {
    const launcherScript = await response.text();
    const launcherData = new Uint8Array(Buffer.from(launcherScript));
    const fileHandler = await open(launcherPath, 'w');
    await writeFile(fileHandler, launcherData, { flag: 'w' });
    await chmod(launcherPath, 0o711);
    await fileHandler.close();
  } catch (err: any) {
    console.error("error setting up launcher script" + err);
    throw (err);
  }
  return launcherPath;
}

async function setupSSHConfig(vmId: string, prKey: string) {
  const codexBinary = await which('codex', { nothrow: true });
  let codexPath = 'codex';
  if (codexBinary === null) {
    codexPath = await setupCodexLauncher();
  }
  // For testing change codex to path to a compiled codex binary or add --config
  exec(`${codexPath} generate ssh-config`, (error, stdout, stderr) => {
    if (error) {
      window.showErrorMessage('Failed to setup ssh config. Run VSCode in debug mode to see logs.');
      console.error(`Failed to setup ssh config: ${error}`);
      return;
    }
    console.debug(`stdout: ${stdout}`);
    console.debug(`stderr: ${stderr}`);
  });

  // save private key to file
  const prkeyDir = `${process.env['HOME']}/.config/codex/ssh/keys`;
  await ensureDir(prkeyDir);
  const prkeyPath = `${prkeyDir}/${vmId}.vm.codex-vms.internal`;
  try {
    const prKeydata = new Uint8Array(Buffer.from(prKey));
    const fileHandler = await open(prkeyPath, 'w');
    await writeFile(fileHandler, prKeydata, { flag: 'w' });
    await chmod(prkeyPath, 0o600);
    await fileHandler.close();
  } catch (err: any) {
    // When a request is aborted - err is an AbortError
    console.error('Failed to setup ssh config: ' + err);
    throw (err);
  }
}

function connectToRemote(username: string, vmId: string, workDir: string) {
  try {
    const host = `${username}@${vmId}.vm.codex-vms.internal`;
    const workspaceURI = `vscode-remote://ssh-remote+${host}${workDir}`;
    const uriToOpen = Uri.parse(workspaceURI);
    console.debug("uriToOpen: ", uriToOpen.toString());
    commands.executeCommand("vscode.openFolder", uriToOpen, false);
  } catch (err: any) {
    console.error('failed to connect to remote: ' + err);
  }
}

async function ensureDir(dir: string) {
  try {
    await mkdir(dir, {recursive: true, mode: 0o700});
  } catch (err: any) {
    if (err.code !== 'EEXIST') {
      console.error('Failed to setup ssh keys directory: ' + err);
      throw (err);
    }
  }
}
