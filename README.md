# tigera/api

This is the location of the Tigera API structures for use in code that wants to use the Tigera API.

## How to use

One way is to import the clientset directly and use it. See [examples/main.go](examples/main.go) for some example code.

## Adding new APIs

1. Add the new types to `pkg/apis/<apigroup>/types.go`

1. Add the new types to `pkg/apis/<apigroup>/<version>/types.go`

1. Update generated code, including clients, informers, etc.

   ```
   make .generate_files
   ```
