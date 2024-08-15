---
title: Running Scripts
---

Scripts are shell commands that can be defined in your codex.json file. They can be executed by using the  `codex run` command. Scripts started with `codex run` are launched in a interactive `codex shell` that terminates once the script finishes, or is interrupted by CTRL-C.

Scripts will run after your packages finish installing, and after your `init_hook` completes.

## Configuring scripts

Scripts can be added in your `codex.json`. Scripts require a unique name, and a command or list of commands to run:

```json
"shell": {
    "init_hook": "echo \"Hello \"",
    "scripts": {
        "echo_once": "echo \"World\"",
        "echo_twice": [
            "echo \"World\"",
            "echo \"Again\""
        ]
    }
}
```

## Running your scripts

To run a script, use `codex run <script_name>`. This will start your shell, run your `init_hook`, and then run the script:

```bash
$ codex run echo_once
Installing nix packages. This may take a while... done.
Starting a codex shell...
Hello
World

$ codex run echo_twice
Installing nix packages. This may take a while... done.
Starting a codex shell...
Hello
World
Again
```

Your codex shell will exit once the last line of your script has finished running, or when you interrupt the script with CTRL-C (or a SIGINT signal).

## Running a One-off Command

You can use `codex run` to run any command in your Codex shell, even if you have not defined it as a script. For example, you can run the command below to print "Hello World" in your Codex shell:

```bash
codex run echo "Hello World"
```

You can also run commands that use flags as normal. For example:

```bash
codex run lsof -i :80
```

Note that if you want to pass flags to `codex run`, you should pass them before the command. For example:

```bash
# Run `lsof -i :80` in your codex shell in quiet mode
codex run -q lsof -i :80
```

## Run Scripts with Custom Environment Variables

You can use the `--env` flag to set custom environment variables in your Codex shell. For example, the following command will set the `MY_VAR` environment variable to `my_value` when running the `echo` command:

```bash
codex run --env MY_VAR=my_value echo $MY_VAR
```

You can also load environment variables from a .env file using the `--env-file` flag. For example, the following command will load the environment variables from the `.env.codex` file in your current directory:

```bash
codex run --env-file .env.codex echo $MY_VAR
```

## Tips on using Scripts

1. Since `init_hook` runs every time you start your shell, you should primarily use it for setting environment variables and aliases. For longer running tasks like database setup, you can create and run a Codex script
2. You can use Codex scripts to start and manage long running background processes and daemons.
   1. For example -- If you are working on a LAMP stack project, you can use scripts to start MySQL and Apache in separate shells and monitor their logs. Once you are done developing, you can use CTRL-C to exit the processes and shells
3. If a script feels too long to put it directly in `codex.json`, you can save it as a shell script in your project, and then invoke it in your `codex scripts`.
4. For more ideas, see the LAMP stack example in our [Codex examples repo](https://github.com/khulnasoft/codex/tree/main/examples/stacks/lapp-stack).
