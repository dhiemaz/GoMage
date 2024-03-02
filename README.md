## GoMage - Go Image Processing 
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://github.com/dhiemaz)

GoMage is a image colour adjustment CLI application build entirely using Go Programming Language.

### Features
* Image temperature adjustment

## System Requirement ##
* Go SDK +1.21
* [Cobra CLI](https://github.com/spf13/cobra)
* [go-colorful](https://github.com/lucasb-eyer/go-colorful)

### Project structure layout
```text
├── cmd/
    ├── action/
├── internal/
    ├── image/
    ├── io/
├── vendor/
├── bin/
├── LICENSE
├── go.mod
├── main.go
├── Makefile
├── README.md
```

### Project layout explanation

#### - bin package
bin contains build executable file .

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

### Build
If you use *nix OS (MacOs, Linux) and have `make` tool installed on your system, then you can easily build project using `make` command.

* MacOs and Linux using `make` command

```shell
$ make build
```
It will create executable file `gomage` inside `/bin` directory of the project. To run the application, you can use the following command.

```shell
$ ./bin/gomage
```

* MacOs and Linux using `go` tool

```shell
$ go build -o gomage
```

It will create executable file `gomage` in the root directory of the project. 

* Windows using `go` tool

For windows user and don't have `make` tool installed, you still can build project using `go` tool.

```shell
c:\> go build -o gomage
```
It will create executable file `gomage` in the root directory of the project. 

### Running GoMage

Running GoMage is very simple, you can run the application using the following command. if you use `make` tool, the executable file will be created inside `/bin` directory, you can use the following command.

```shell
$ ./bin/gomage
```

it will show the following output.

```shell
Go Image manipulation command line.

Usage:
  GoMage [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  temperature adjust image temperature.

Flags:
  -h, --help   help for GoMage

Use "GoMage [command] --help" for more information about a command.
```

### Running GoMage temperature command

For now the only feture available is temperature adjustment. You can run the following command to adjust the temperature of the image.

```shell
$ ./bin/gomage temperature --input <input_image> --output <output_image> --temperature <temperature>
```

it will show the following output.

```shell

 ________          _____                         
 /  _____/  ____   /     \ _____     ____   ____  
/   \  ___ /  _ \ /  \ /  \\__  \   / ___\_/ __ \ 
\    \_\  (  <_> )    Y    \/ __ \_/ /_/  >  ___/ 
 \______  /\____/\____|__  (____  /\___  / \___  >
        \/               \/     \//_____/      \/ 
Error: required flag(s) "input image", "output image" not set
Usage:
  GoMage temperature [flags]

Flags:
  -h, --help                  help for temperature
  -i, --input image string    input image file (mandatory).
  -o, --output image string   output image file (mandatory).
  -t, --temperature float     temperature value (optional), use positive number for warmer and negative number for cooler. if not set it will set to default (0).
```

* input_image: is the path of the input image.
* output_image: is the path of the output image.
* temperature: is the temperature value that you want to adjust. Use positive number for warmer and negative number for cooler. if not set it will set to default (0).

Example:

```shell
$ ./bin/gomage temperature --input ./testdata/sample.jpg --output ./testdata/sample_output.jpg --temperature 10
```
