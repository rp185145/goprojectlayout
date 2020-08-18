# Building and Testing

## Build Script Guidelines

See `scripts/buildall.sh` for a script that:

- Enforce proper code formatting
- Passes the `vet` tests
- Executes unit tests with the race detector enabled
- Builds executables for Linux and Windows
- Executes the applications as a sanity check

The first four items seem to be common for many populate applications. You may find some useful information in the script for creating your own build scripts for use on build servers.

## Building

The source files for the individual executables are found in the `cmd/` subdirectory, which in turn contains a subdirectory for each executable.

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

The source file layout for this isn't ideal. Normally each subcommand should probably be in its own source file. This code is mainly to show how to use Cobra to frame out a command-line application. Depending on feedback, I may restructure this in the future if any users find that useful.

```
go build -o bin/gplcli ./cmd/gplcli/...
```

## Testing

Test functions and modules were created using the VSCode Go extension. Use "Ctrl-Shift-P -> Go: Generate unit tests for function".

To test everything and generate coverage output. Coverage output goes to `coverage.txt`.

```
go test -coverprofile=coverage.txt ./...
```

To process coverage output to HTML. HTML output goes to `coverage.html`.

```
go tool cover -html=coverage.txt -o coverage.html
```
