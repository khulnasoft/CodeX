---
title: Eclipse IDE 
---


## Java
This guide describes how to configure Eclipse to work with a codex Java environment.

### Setting up Codex shell
To create a codex shell make sure to have codex installed. If you don't have codex installed follow the installation guide first. Then run the following commands from the root of your project's repo:

1. `codex init` if you don't have a codex.json in the root directory of your project.
2. `codex add jdk` to make sure jdk gets installed in your codex shell.
3. `codex shell -- 'echo $JAVA_HOME'` to activate your codex shell temporarily to find the path to your java home. Copy and save the path. It should look something like:
    ```bash
    /nix/store/qaf9fysymdoj19qtyg7209s83lajz65b-zulu17.34.19-ca-jdk-17.0.3
    ```
4. Open Eclipse IDE and create a new Java project if you don't have already
5. From the top menu go to Run > Run Configurations > JRE and choose **Alternate JRE:**
6. Click on **Installed JREs...**  and click **Add...** in the window of Installed JREs.
7. Choose **Standard VM** as JRE Type and click Next.
8. Paste the value you copied in step 4 in **JRE HOME** and put an arbitrary name such as "codex-jre" in **JRE Name** and click Finish.
9. Click **Apply and Close** in Installed JREs window. Then close Run Configurations.

Now your project in Eclipse is setup to compile and run with the same Java that is installed in your codex shell. Next step is to run your Java code inside Codex.

### Setting up Eclipse Terminal

The following steps show how to run a Java application in a codex shell using the Eclipse terminal. Note that most of these steps are not exclusive to Eclipse and can also be used in any Linux or macOS terminal.

1. Press `ctrl + alt/opt + T` to open terminal window in Eclipse.
2. Navigate to the projects root directory using `cd` command.
3. Make sure `codex.json` is present in the root directory `ls | grep codex.json`
4. Run `codex shell` to activate codex shell in the terminal.
5. Use `javac` command to compile your Java project. As an example, if you have a simple hello world project and the directory structure such as: 
    ```bash
    my_java_project/
    -- src/
    -- -- main/
    -- -- -- hello.java
    ```
    You can use the following command to compile:
    to compile:
    ```bash
    javac my_java_project/src/main/hello.java
    ```
6. Use `java` command to run the compiled proect. For example, to run the sample project from above:
    ```bash
    cd src/
    java main/hello
    ```

If this guide is missing something, feel free to contribute by opening a [pull request](https://github.com/khulnasoft/codex/pulls) in Github.
