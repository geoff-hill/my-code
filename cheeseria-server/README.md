# Server

This subproject contains the backend server developed in Java. It uses JDK17, so you can code in a modern style, using new features like records.

It persists the cheese data to a JSON File, stored in a ```.cheeseria``` directory under your user home. 


## Building and Running the Server in a Terminal

The project uses Gradle as its build tool, so building the app is simple.

We use: 
- [Gradle Wrapper](https://docs.gradle.org/current/userguide/gradle_wrapper.html), which will download the correct version of gradle if needed, and then use it.
- [Gradle Java Toolchains](https://docs.gradle.org/current/userguide/toolchains.html#sec:auto_detection), which can auto-detect JDKs, and auto-provision one if needed.
 
You do not need to have gradle installed on your machine, although it won't cause problems if you do. You don't need JDK 17, although the initial build will be quicker if one can be detected.Ub fact, the only pre-requisite is that you have some version of Java (JRE/JDK) installed to bootstrap the gradle-wrapper. It has been tested on a JRE 8, so chances are you have a newer one than that.

So clone the repo from github, change to the server directory, and run:

### `./gradlew run`

This will work on any Unix/Linux/Macos terminal, and even in Windows Terminal with Powershell.  Should you really insist upon running it in an old-style `cmd` window, then you'll need to use `.\gradlew.bat run`.

Either way, the script will download and install the correct version of gradle (if needed), download and install a correctly versioned JDK17, run the build, and start the server. Suffice to say, this takes a while the first time. But subsequent builds will be very fast.

The gradle command will not terminate, as the build runs the server on the terminal window, and you will see the log output onscreen. To stop the server, just use `^C` or the kill technique of your choice.

The server backend listens on [http://localhost:9100](http://localhost:9100), and you can test it using the built-in [Swagger UI](http://localhost:9100/swagger-ui).

## Developing in an IDE

You can use your favourite IDE (or text editor) to make changes to the project).

IDEs can get confused if they get wrong versions of things into their cache. So **do the terminal install (above) first - this guarantees that you have a correct version of JDK17, and that the gradle wrapper has provisioned gradle correctly**.

The IDEs can throw some warning dialogs and error messages the first time if you are doing a fresh install of the IDE and Java, but it seems to work pretty reliably on Windows or MAC - sadly I don't have Linux at the moment to test, but that should be pretty easy too.

Here's some IDE specific help if you need it:
- [IntelliJ](./docs/IntelliJ_IDEA_Setup_Guide.md)
- [VSCode](./docs/VSCode_Setup_Guide.md)
- [Eclipse](./docs/Eclipse_Setup_Guide.md) 

## Code

The application uses features such as records from Java 17, so you can code in a modern java style. To put everyone on an equal footing, it does not use either of the popular mega-frameworks (spring-boot or Jakarta-EE). The libraries directly used are:

- [Javalin](http://javalin.io) - lightweight Java and Kotlin web framework - very quick to learn, what you see is what you get;
- [Jackson](https://github.com/FasterXML/jackson) - JSON processing (plus YAML and other formats);
- [SLF4J](https://www.slf4j.org/manual.html) - Simple Logging Facade for Java.

### Configuration and Data files

The application persists its configuration and data in a `.cheeseria` folder under your home directory. The folder and these files are created automatically if they don't exist, and any changes you make to them will not be overwritten by the running application.

This does mean that if you want to use the default values again, or they have been somehow corrupted, or you have changed the format of the `cheeseria.yml` then you just need to delete the file and restart the app. Same thing to reinitialise the "database", in the `cheeses.json` file

