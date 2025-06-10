# Numbers to Words CLI

A command-line tool, written in Go, that converts a given number (from 0 to 100,000) into its British English word representation, including correct commas and "and" separators.


## Features

- Converts numbers from 0 to 100,000.
- Follows British English conventions (e.g., "one hundred and one", "two thousand, five hundred).
- Provides multiple ways to build and run: Make, Docker, or native Go commands.
- Comrehensive unit tests.

## Prerequisites

- **Go**: Version 1.22+
- **Make**: Recommended for easy building and testing.
- **Docker**: Required for containerised method.

## Getting Started

This project uses a Makefile to simplify building and testing. This is the recommended way to get started.

#### 1. Clone the repository (or unzip the submission folder)
#### 2. Build the application
From the project root, run the make command. This will run the tests and compile the binary into the ``` /bin ``` directory.

```sh
make
```

## Usage

Once the application is built, you can run it from the command line. The executable will be located at ./bin/numbers-to-words (the executable will end with .exe, in windows).

#### Command Syntax

```sh
./bin/numbers-to-words [number]
```

#### Examples

```sh
./bin/numbers-to-words 42
# Output: forty-two

./bin/numbers-to-words 12345
# Output: twelve thousand, three hundred and forty-five

./bin/numbers-to-words 99009
# Output: ninety-nine thousand and nine
```

To clean up the binary, use the following command:

```sh
make clean
```

## Alternative Build & Run Methods

### Using Docker
The application can be built and run as a self-contained Docker image.

#### 1. Build the image:

```sh
docker build -t numbers-to-words-app .
```
#### 2. Run the container:
The --rm flag automatically removes the container after it exits.

```sh
docker run --rm numbers-to-words-app 54321
# Output: fifty-four thousand, three hundred and twenty-one
```

### Using Go Commands Directly

If you don't have make, you can use the native Go toolchain.

#### 1. Run the unit tests

```sh
go test -v ./src/... 
```

#### 2. Build the binary:

```sh
go build -o bin/numbers-to-words ./src
```

### Notes for Reviewer (Design Decisions)

A few technical decisions were made during the development of this tool:

- **Makefile:** A cross-platform ```Makefile``` is provided to automate builds and tests. It automatically handles the ```.exe``` extension on Windows.
- **Docker:** A multi-stage ```Dockerfile``` is used to create a minimal and secure final image, separating the build environment from the runtime environment.
- **Testing:** The project includes comprehensive unit testing (```parse_num_test.go```) to validate the core conversion logic.
- **Code Structure:** The core conversion logic (```parse_num.go```) is decoupled from the command-line interface (```main.go```), making the logic reusable and easy to test in isolation.
- **Dependency-Free:** To maximise simplicity and long-term stability, the project relies solely on the Go standard library, avoiding external dependencies.


## Project Structure

The project follows a standard Go application layout:

* `bin/`: This directory contains the compiled application binaries.
* `src/`: This directory holds all the Go source code for the project.
    * `main.go`: The main entry point for the application.
    * `parse_num.go`: Source file containing number parsing logic.
    * `parse_num_test.go`: Unit tests for the number parsing functionality.
* `Dockerfile`: Defines the instructions for building the application into a Docker container.
* `go.mod`: The Go modules file that manages the project's dependencies.
* `Makefile`: Contains helper commands for building, testing, and running the project.
* `README.md`: Project documentation.