# Setting Up Example Projects For Java

## Java with Maven
Maven is an all-in-one CI-CD tool for building testing and deploying Java projects. To setup a sample project with Java and Maven in codex follow the steps below:

1. Create a dummy folder: `dummy/` and call `codex init` inside it. Then add the nix-pkg: `codex add jdk` and `codex add maven`.
    - Replace `jdk` with the version of JDK you want. Get the exact nix-pkg name from `search.nixos.org`.
2. Then do `codex shell` to get a shell with that `jdk` nix pkg.
3. Then do: `mvn archetype:generate -DgroupId=com.codex.mavenapp -DartifactId=codex-maven-app -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false`
    - In the generated `pom.xml` file, replace java version in `<maven.compiler.source>` with the specific version you are testing for.
4. `mvn package` should compile the package and create a `target/` directory.
5. `java -cp target/codex-maven-app-1.0-SNAPSHOT.jar com.codex.mavenapp.App` should print "Hello World!".
6. Add `target/` to `.gitignore`.

## Java with Gradle
To test a sample Gradle app with codex, follow the steps below:

1. Create a dummy folder: `dummy/` and call `codex init` inside it. Then add these packages: `codex add jdk` and `codex add gradle`.
    - Replace `jdk` with the version of JDK you want. Get the exact nix-pkg name from `search.nixos.org`.
2. Then do `codex shell` to get a shell with that `jdk` nix pkg.
3. Then do: `gradle init`
    - In the generated `gradle.build` file, put the following text block:
        ```gradle
        apply plugin: 'java'
        apply plugin: 'application'
        sourceCompatibility = 17
        targetCompatibility = 17
        mainClassName = 'hello.HelloWorld'
        jar {
            manifest {
                attributes 'Main-Class': 'hello.HelloWorld'
            }
        }
        ```
4. `gradle build` should compile the package and create a `build/` directory that contains an executable jar file.
5. `gradle run` should print "Hello World!".
6. Add `build/` to `.gitignore`.
