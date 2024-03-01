## GoMage - Go Image Processing 
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://github.com/dhiemaz)

GoMage is a image colour adjustment CLI application build entirely using Go Programming Language.

### Features
* Image temperature adjustment

## System Requirement ##
* GO SDK +1.21
* Cobra CLI

### Project structure layout
```text
├── cmd/
    ├── action/
├── internal/
    ├── image/
    ├── io/
├── vendor/
├── LICENSE
├── go.mod
├── main.go
├── Makefile
├── README.md
```

### Project layout explanation

#### - cmd package
cmd package contains all the action functions.

#### - internal package
internal package contains all package used internally in Rates Services.

#### - internal/image package
internal/image package contains all functions for image processing (eg: temperature).

#### - internal/io package
internal/io package contains all i/o functions (eg: file).

### Setup
* Go SDK, GoMage is entirely using Go programming language, it's mean that for the code can be run, you must have
  Go SDK installed in your system. Please refer to [installation](https://go.dev/doc/install) instruction and match it
  with your OS.
* Git, Go tool is using git as version control system. Please refer to [installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
  instruction and match it with your OS.
* makefile, is used to build the project. Please refer to [installation](https://www.gnu.org/software/make/manual/make.html)

### Build and Running
If you use *nix OS (MacOs, Linux) and have `make` tool installed on your system, then you can easily build project using `make` command.

* MacOs and Linux using `make` command

```shell
$ make build
```
It will create executable file `gomage` in the root directory of the project. To run the application, you can use the following command.

```shell
$ ./gomage
```

* MacOs and Linux using `go` tool

```shell
$ go build -o gomage
```

It will create executable file `gomage` in the root directory of the project. To run the application, you can use the following command.

```shell
$ ./gomage
```

* Windows using `go` tool

For windows user and don't have `make` tool installed, you still can build project using `go` tool.

```shell
c:\> go build -o gomage
```
It will create executable file `gomage` in the root directory of the project. To run the application, you can use the following command.

```shell
c:\> gomage.exe
```


