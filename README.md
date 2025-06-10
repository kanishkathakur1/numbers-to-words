# Numbers to Words CLI, in Go (take home task)

This is a command line tool that converts a given number into its English word representation. The application is written in Golang, and can be compiled into a binary and executed that way. 

Currently, the bin folder has a binary that works for MacOS on ARM-based chips (M1, M2, M3 etc.). If you are using a different operating system and hardware, please remove this manually or with the **``` make clean ```** command and, follow the set-up/build steps to build the application for your system.

## Prerequisites

- Go (version 1.24 used for the project)
- Make (optional)
- Docker (optional)

## Building and testing the application

### Using Makefile (recommended)

To build the application using the make tool, and run all the unit tests, run the following command in the root of the project:

```sh
make
```

To run just the unit tests, use the following command:

```sh
make test
```

To just build the application, use the following command:

```sh
make build
```

To clean up the binaries, using the follwing command:

```sh
make clean
```

### Using Docker

This application can also be run using docker, which provides an isolated and consistent environment. The Dockerfile is provided in the root of the directory, so simply run the following command to build the image (make sure that docker is running first):

```sh
docker build -t numbers-to-words-app .
```
This will create an image called numbers-to-words-app. To run the application use the follwing command:

```sh
docker run --rm numbers-to-words-app [arg]
```
Replace the **'[arg]'** with your number that you want to pass as the argument. The --rm flag cleans up the container after use.


### Using Go (without make tool)

To run build the binary, without using the Makefile, run the following commands in the terminal

```sh
go build -o bin/numbers-to-words ./src
```
