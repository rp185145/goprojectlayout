# Go Project Layout Demo

A repository for demonstrating a basic Go project layout

## About

I quickly put much of the ideas in this repository together during a demo call to demonstrate how we organize a project that consists of multiple packages, APIs, and executables.

## Packages

This project implements 3 packages usable as APIs:

- `github.com/rp185145/goprojectlayout` (root package)
- `github.com/rp185145/goprojectlayout/emitter`
- `github.com/rp185145/goprojectlayout/idutils`

## Applications

### `hello` and `howdy`

Demonstrates how to have two separate applications call the same package/API (`emitter`) in the same module and repository.

### `uuid`

Demonstrates how to have a completely separate application calling a completely separate package (`idutils`) in the same repository as other applications.

### `demo`

Demonstrates a separate application calling an exported function in the root (`goprojectlayout`) package.

## Building

### `hello`

```
go build -o bin/hello ./cmd/hello/...
```

### `howdy`

```
go build -o bin/howdy ./cmd/howdy/...
```

### `uuid`

```
go build -o bin/uuid ./cmd/uuid/...
```

### `demo`

```
go build -o bin/demo ./cmd/demo/...
```

### `gplcli`

This application implements a simple user interface using Cobra. Because the packages are written in an API style, they can be combined, mixed, and re-used. This application contains the subcommand `hello`, `howdy`, `uuid`, and `demo` to match the functionality of the above applications. A `custom` subcommand was added to demonstrate how to use a command-line option.

The source layout for this isn't ideal. Normally each subcommand should probably be in its own source file. This code is mainly to show how to use Cobra to frame out a command-line application.

```
go build -o bin/gplcli ./cmd/gplcli/...
```

## Testing

Test functions and modules were created using the VSCode Go extension. Ctrl-Shift-P -> Go: Generate unit tests for function.

To test everything and generate coverage output. Coverage output goes to `coverage.txt`.

```
go test -coverprofile=coverage.txt ./...
```

To process coverage output to HTML. HTML output goes to `coverage.html`.

```
go tool cover -html=coverage.txt -o coverage.html
```

## API Code

Notice in these cases there's still a `go.mod` in the root and the module is designated as `github.com/rp185145/goprojectlayout` but the API packages are in `emitter/` and `idutils/`, therefore the imports are `github.com/rp185145/goprojectlayout/emitter` and `github.com/rp185145/goprojectlayout/idutils`, respectively.

Also notice that `emitter` can call exported functions in `idutils` **or** `idutils` can call exported functions in `emitter` but not both because Go will complain about circular references. In this demo, `emitter` calls a function in `idutils`.

## Modules

Create a new module with

```
go mod init github.com/rp185145/goprojectlayout
```

## Consuming this Package

Because this package has been committed to a public repository and especially because it has a [version tagged release](https://github.com/rp185145/goprojectlayout/releases) it can be consumed by other projects. See the corresponding public project, [Google Project Layout Consumer](https://github.com/rp185145/gcpconsumer) for an example.

## Godoc

These packages have been lightly commented to demonstrate the [`godoc` standard for documentation](https://pkg.go.dev/golang.org/x/tools/cmd/godoc?tab=doc). Install `godoc` (probably by using `go get golang.org/x/tools/cmd/godoc` on your OS), then in the root of the repository, execute something like:

```
godoc --http=":6060"
```

Browse to http://localhost:6060/pkg/github.com/rp185145/goprojectlayout/ and observe how the output matches in the comments and function examples.

## Useful Information and Packages

### Go Modules

https://blog.golang.org/using-go-modules

### Cobra

For command-line interfaces. This package is so good for implementing a CLI that I rarely write code without it.

https://github.com/spf13/cobra

### Viper

For configuration, especially of 12-factor apps. It can be integrated with Cobra. I've only experimented with this package.

https://github.com/spf13/viper
