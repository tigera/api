# tigera/api

This is the location of the Tigera API structures for use in code that wants to use the Tigera API.

**NOTE**: At the moment, this is a mirror of the canonical resource definitions which are located in libcalico, as well as
the apiserver. Changes to those resources need to be manually updated in this repository in order to be used. Eventually,
this will become the canonical source for API definitions.

## How to use

One way is to import the clientset directly and use it. See [examples/main.go](examples/main.go) for some example code.

## Adding new APIs

1. Add the new types to `pkg/apis/<apigroup>/types.go`

1. Add the new types to `pkg/apis/<apigroup>/<version>/types.go`

1. Update generated code, including clients, informers, etc.

   ```
   make .generate_files
   ```
