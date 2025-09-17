## gvm new

Create a new Go project with the current active version

### Synopsis

Create a new Go project initialized with the Go version currently set by gvm.

Example:
  gvm new myapp
This will create a folder 'myapp', initialize a Go module,
and set it up using the active Go version.

```
gvm new [project-name] [flags]
```

### Options

```
  -h, --help             help for new
  -m, --module string    Go module path (default is project name)
  -V, --version string   Go version(default in current version)
```

### Options inherited from parent commands

```
  -v, --verbose   verbose output
```

### SEE ALSO

* [gvm](gvm.md)	 - Go version manager for installing and switching between multiple Go versions

