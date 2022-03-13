![Continuous Integration](https://github.com/chris-crone/containerized-go-dev/workflows/Continuous%20Integration/badge.svg)

Containerized Go Development Environment
------------------------------------------------

This repository contains an example Go project with a containerized development
environment. The example project is a simple CLI tool that echos back its
inputs.

## Prerequisites

The only requirements to build and use this project are Docker and `make`. The
latter can easily be substituted with your scripting tool of choice.

You will also need to enable the BuildKit builder in the Docker CLI. This can be
done by setting `DOCKER_BUILDKIT=1` in your environment.

### macOS

* Install [Docker Desktop](https://www.docker.com/products/docker-desktop)
* Ensure that you have `make` (included with Xcode)
* Run `export DOCKER_BUILDKIT=1` in your terminal or add to your shell
  initialization scripts

### Windows

* Install [Docker Desktop](https://www.docker.com/products/docker-desktop)
* Ensure that you have `make`
* If using PowerShell, run `$env:DOCKER_BUILDKIT=1`
* If using command prompt, run `set DOCKER_BUILDKIT=1`

### Linux

* Install [Docker](https://docs.docker.com/engine/install/)
* Ensure that you have `make`
* Run `export DOCKER_BUILDKIT=1` in your terminal or add to your shell
  initialization scripts

## Getting started

Building the project will output a static binary in the bin/ folder. The
default platform is for Windows but this can be changed using the `PLATFORM` variable:
```console
$ make                        # build for your host OS
$ make PLATFORM=darwin/amd64  # build for macOS
$ make PLATFORM=windows/amd64 # build for Windows x86_64
$ make PLATFORM=linux/amd64   # build for Linux x86_64
$ make PLATFORM=linux/arm     # build for Linux ARM
```

You can also set the name of the application in the bin/ folder.
The default name of the application is set as `myApp` in Makefile.
```console
$ make NAME=example                         # build example for your host OS
$ make PLATFORM=windows/amd64 NAME=example  # build example.exe for Windows x86_64
$ make NAME=win/example                     # build example application in /bin/win folder
```

You can then run the binary, which is a simple echo binary, as follows:
```console
$ ./bin/myApp hello world!
hello world!
```

To run the unit tests run:
```console
$ make unit-test
```

To run the linter:
```console
$ make lint
```

There's then a helpful `test` alias for running both the linter and the unit
tests:
```console
$ make test
```

## Structure of project

### Dockerfile

The [Dockerfile](./Dockerfile) codifies all the tools needed for the project
and the commands that need to be run for building and testing it.

### Makefile

The [Makefile](./Makefile) is purely used to script the required `docker build`
commands as these can get quite long. You can replace this file with a scripting
language of your choice.

### CI

The CI is configured in the [ci.yaml file](./.github/workflows/ci.yaml). By
containerizing the toolchain, the CI relies on the toolchain we defined in the
Dockerfile and doesn't require any custom setup.

## Related Blog

[Containerize your local Go developer environment series](https://www.docker.com/blog/tag/go-env-series/)
## Read more

* [Docker build reference documentation](https://docs.docker.com/engine/reference/commandline/build/)
* [Experimental Dockerfile syntax](https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/experimental.md)

---

# Want to learn more about Golang?
If you have an experience in any programming language, then dive into two series video from Jake Wright on *Youtube*. First video is [Learn go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI). The second video is [Concurrency in Go](https://www.youtube.com/watch?v=LvgVSSpwND8). 

## Catchup with Golang
Golang works in the workspace. You can get the workspace by typing `go env GOPATH`.

The source code of the files are located in the `src` folder inside the workspace.

The first line of the golang file needs to be name of the package. Any executable application written in golang should have one package named main.

```go
package main
```

The next section is the import, where you import different packages. The most common package is `fmt` which executes input and outputs.

```go
import(
  "fmt"
)
```

*Note, there is no commas when listing packages.*

The main package should have a main function which is the entry point of the application. Here is the formula of writing the function that doesn't accept any arguments and doesn't return any value.

```go
func main() {
  fmt.Println("Hello world!")
}
```

To execute the application:
* `go <file name>.go` &ndash; executes the file.
* `go build` &ndash; compiles the binaries into the current folder.
* `go install` &ndash; compiles the binaries into `/bin/` folder.

The go will cashe the external dependencies into the `/package` folder.

### Variables

Declaration of variable:

```go

var x int   // uninitialized variable is 0.
var y int = 5 
z := x + y  // same as var z int = x + y

```

### Control statements

```go
num := 5

if num > 5 {
  fmt.Println("More than 5")
} else if num < 5 {
  fmt.Println("Less than 5")
} else {
  fmt.Println("It's five")
}
```


### Arrays, Slices, Maps

Arrays are one data type fixed length.

```go
var a [5]int
a[2] = 7

b := [5]int{0,0,7,0,0}    // identical to var b [5]int
                          // b[2] = 7
```

Slices are an abtraction on top of arrays without fixed length of the list.

```go
var a []int
a = append(a, 0)
a = append(a, 0)
a = append(a, 7)
a = append(a, 0)
a = append(a, 0)

b := []int{0,0,7,0,0}      // slice with 5 elements
```

Maps are key&ndash;value pairs. 

```go
vertices := make(map[string]int)

vertices["triangle"] = 3
vertices["square"] = 4

delete(vertices, "square")
```

### Loops

```go
for i := 0; i < 5; i++ {
  fmt.Println(i)
}

//------------------------------------
//
// loop over array
//
//------------------------------------ 
arr := []string{"orange", "coconut", "banana", "apple"}

for index, value := range arr {
  // fmt.Println puts a space between parameters
  fmt.Println("index:", index, "value:", value)
}

m := make(map[string]string)
m["a"] = "alpha"
m["b"] = "beta"

for key, value := range m {
  fmt.Println("index:", key, "value:", value)
}
```

### Functions

```go
package main

import(
  "fmt"
  "math"
  "errors"
)

func main() {
  res := sum(1, 4)

  // go doesn't have exceptions
  root, err := sqrt(res)

  if err == nil {
    fmt.Println("root:", root)
  } else {
    fmt.Println("err:", err)
  }
}

func sum(x int, y int) int {
  return x + y
}

func sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, errors.New("x must be positive")
  }

  return math.Sqrt(x), nil
}
```


### Structs

```go
package main 

import("fmt")

// structs are defined outside of functions
type person struct {
  name string 
  age int
}

func main() {
  medet := people{name: "Medet", age: 26}

  fmt.Println(medet.name)
}
```

