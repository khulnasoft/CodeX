# Java

In addition to installing the JDK, you'll need to install either the Maven or Gradle build systems in your shell.

In both cases, you'll want to first activate `codex shell` before generating your Maven or Gradle projects, so that the tools use the right version of the JDK for creating your project.

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/java)

## Adding the JDK to your project

`codex add jdk binutils`, or in your `codex.json`

```json
  "packages": [
    "jdk@latest",
    "binutils@latest"
  ],

```

This will install the latest version of the JDK. To find other installable versions of the JDK, run `codex search jdk`.

Other distributions of the JDK (such as OracleJDK and Eclipse Temurin) are available in Nixpkgs, and can be found using [NixPkg Search](https://search.nixos.org/packages?channel=22.05&from=0&size=50&sort=relevance&type=packages&query=jdk#)

## Maven

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/java/maven/hello-world)

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/open/templates/java-maven)

Maven is an all-in-one CI-CD tool for building testing and deploying Java projects. To setup a sample project with Java and Maven in codex follow the steps below:

1. Create a dummy folder: `dummy/` and call `codex init` inside it. Then add the nix-pkg: `codex add jdk` and `codex add maven`.
    - Replace `jdk` with the version of JDK you want. Get the exact nix-pkg name from `search.nixos.org`.
2. Then do `codex shell` to get a shell with that `jdk` nix pkg.
3. Then do: `mvn archetype:generate -DgroupId=com.codex.mavenapp -DartifactId=codex-maven-app -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false`
    - In the generated `pom.xml` file, replace java version in `<maven.compiler.source>` with the specific version you are testing for.
4. `mvn package` should compile the package and create a `target/` directory.
5. `java -cp target/codex-maven-app-1.0-SNAPSHOT.jar com.codex.mavenapp.App` should print "Hello World!".
6. Add `target/` to `.gitignore`.

An example `codex.json` would look like the following:

```json
{
  "packages": [
    "maven",
    "jdk",
    "binutils"
  ],
  "shell": {
    "init_hook": null
  }
}
```
