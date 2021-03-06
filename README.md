# Go Project Layout Demo

A repository for demonstrating a basic Go project layout

## About

Many of the ideas in this repository were coded live during a demo call demonstrating how I typically organize my Go repositories. By request, this project contains multiple APIs implemented with packages, several executables, simple unit testing, and a basic Cobra implementation.

Go through this README and the code to see how I work. If you know some better organization techniques and would like to share them, let me know with a message or even a pull request.

## Building

See the [Building and Testing](BUILDING.md) document.

## Packages

This project implements 3 packages usable as APIs:

- `github.com/rp185145/goprojectlayout` (root package)
- `github.com/rp185145/goprojectlayout/emitter`
- `github.com/rp185145/goprojectlayout/idutils`

## Applications

### `hello` and `howdy`

Demonstrates how to have two separate applications call the same package/API (`emitter`) in the same module and repository. Note that both are package `main` but since they are in separate directories and don't reference each other, they don't know about each other.

### `uuid`

Demonstrates how to have a completely separate application calling a completely separate package (`idutils`) in the same repository as other applications.

### `demo`

Demonstrates a separate application calling an exported function in the root (`goprojectlayout`) package.

### gplcli

A simple implementation using Cobra to handle the command-line interface processing. This executable contains several subcommands corresponding to the above applications, meaning it contains all the above functionality in a single executable. Execute `gplcli` to see the available subcommands. From there, the `--help` option can be used to determine the options for the subcommands. Only `gplcli custom` contains a command-line option. It can be discovered with `gplcli custom --help`. It is the subcommands, help, and option processing that are handled by Cobra.

## API Code

Notice there's a `go.mod` in the root and the module is designated as `github.com/rp185145/goprojectlayout` but API packages are also contained in `emitter/` and `idutils/`, therefore the imports are `github.com/rp185145/goprojectlayout/emitter` and `github.com/rp185145/goprojectlayout/idutils`, respectively. It is customary to have matching directory and package names.

Realize that `emitter` can call exported functions in `idutils` **or** `idutils` can call exported functions in `emitter` but not both because Go will complain about circular references. For this demo, `emitter` calls a function in `idutils` to show APIs can call each other.

## Modules

[Go Modules](https://blog.golang.org/using-go-modules) became the default method for managing dependencies in Go 1.13. Don't use any other legacy dependency management method.

Create a new module with

```
go mod init github.com/rp185145/goprojectlayout
```

## Consuming this Package

Because this package has been committed to a public repository and especially because it has a [version tagged release](https://github.com/rp185145/goprojectlayout/releases) it can be consumed by other projects. See the corresponding public project, [Google Project Layout Consumer](https://github.com/rp185145/gplconsumer) for an example.

## Godoc

These packages have been lightly commented to demonstrate the [`godoc` standard for documentation](https://pkg.go.dev/golang.org/x/tools/cmd/godoc?tab=doc). Install `godoc` (probably by using `go get golang.org/x/tools/cmd/godoc` on your OS), then in the root of the repository, execute something like:

```
godoc --http=":6060"
```

Browse to http://localhost:6060/pkg/github.com/rp185145/goprojectlayout/ and observe how the output matches in the comments and function examples.

Because I've given this project a permissive license, `pkg.go.dev` also consumes and presents this documentation. You can find it at https://pkg.go.dev/github.com/rp185145/goprojectlayout.

## Useful Information and Packages

### Go Modules

The use of Go Modules for dependency management was introduced in Go 1.11 (August 2018) and became the default in 1.13 (September 2019). Keep this in mind when referencing older documentation.

https://blog.golang.org/using-go-modules

### Cobra

For command-line interfaces. This package is so good for implementing a CLI that I rarely write code without it.

https://github.com/spf13/cobra

### Viper

For configuration, especially of 12-factor apps. It can be integrated with Cobra. I've only experimented with this package.

https://github.com/spf13/viper

## Enhancements

There are useful additions that can be made. One is to illustrate the idea of logging using [Go's standard log package](https://golang.org/pkg/log/). Maybe this will be implemented one day or a PR will be filed for a _simple and straightforward_ logging implementation.