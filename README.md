# SpringX - Spring Boot Project Initializer CLI

A powerful command-line tool for generating Spring Boot projects quickly with customizable configurations.

## Features

- 🚀 Generate Spring Boot projects instantly
- 📦 Support for popular dependencies (Web, JPA, Security, etc.)
- 🔧 Choose between Maven or Gradle build tools
- ☕ Configurable Java versions
- 📝 Auto-generated project structure with best practices
- 🎯 Simple and intuitive CLI interface

## Installation

```bash
# Clone the repository
git clone https://github.com/subrotokumar/springx.git
cd springx

# Build the binary
go build -o springx .

# Optional: Move to PATH
sudo mv springx /usr/local/bin/
```

## Usage

### Basic Usage

Generate a simple Spring Boot project with default settings:

```bash
springx init my-app
```

This creates a new Spring Boot project with:
- Maven as build tool
- Java 17
- Spring Web dependency
- Standard project structure

### Advanced Usage

Customize your project with flags:

```bash
# Generate with Gradle and Java 21
springx init my-api --build=gradle --java=21

# Add multiple dependencies
springx init my-web --deps=web,data-jpa,security

# Custom group ID and package
springx init my-service --group=com.mycompany --package=com.mycompany.myservice

# Specify output directory
springx init my-app --output=./projects/my-app
```

### List Available Dependencies

View all supported Spring Boot dependencies:

```bash
springx init --list
```

Available dependencies include:
- `web` - Spring Web (REST APIs)
- `data-jpa` - Spring Data JPA (Database persistence)
- `security` - Spring Security (Authentication & Authorization)
- `actuator` - Spring Boot Actuator (Monitoring)
- `validation` - Validation (Bean validation)
- `thymeleaf` - Thymeleaf (Template engine)
- `devtools` - Spring Boot DevTools (Hot reload)
- `test` - Spring Boot Starter Test (Testing framework)

## Command Options

### `springx init [project-name] [flags]`

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--group` | `-g` | Group ID | `com.example` |
| `--artifact` | `-a` | Artifact ID | Project name |
| `--version` | `-v` | Project version | `0.0.1-SNAPSHOT` |
| `--description` | `-d` | Project description | Auto-generated |
| `--package` | `-p` | Base package name | Auto-generated |
| `--java` | `-j` | Java version | `17` |
| `--build` | `-b` | Build tool (maven/gradle) | `maven` |
| `--deps` | | Comma-separated dependencies | `web` |
| `--output` | `-o` | Output directory | `./<project-name>` |
| `--list` | `-l` | List available dependencies | - |

## Examples

### Create a REST API with Database

```bash
springx init rest-api --deps=web,data-jpa,validation
```

### Create a Secure Web Application

```bash
springx init secure-web --deps=web,security,thymeleaf --build=gradle
```

### Create a Microservice with Monitoring

```bash
springx init user-service \
  --deps=web,data-jpa,actuator \
  --group=com.mycompany \
  --java=21
```

## Generated Project Structure

```
my-app/
├── src/
│   ├── main/
│   │   ├── java/
│   │   │   └── com/example/myapp/
│   │   │       └── MyAppApplication.java
│   │   └── resources/
│   │       ├── application.properties
│   │       ├── static/
│   │       └── templates/
│   └── test/
│       └── java/
│           └── com/example/myapp/
│               └── MyAppApplicationTests.java
├── pom.xml (or build.gradle)
└── README.md
```

## Running Generated Projects

### Maven Projects

```bash
cd my-app
./mvnw spring-boot:run
```

### Gradle Projects

```bash
cd my-app
./gradlew bootRun
```

## Requirements

- Go 1.19 or higher (for building from source)
- Java 17+ (for running generated projects)
- Maven 3.6+ or Gradle 7.0+ (for building generated projects)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

See LICENSE file for details.

## Author

Subroto Kumar <subrotokumar@outlook.in>

## Related Projects

- [Spring Initializr](https://start.spring.io/) - Official Spring Boot project generator
- [Spring Boot](https://spring.io/projects/spring-boot) - Spring Boot framework
