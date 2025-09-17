## gvm list

List Go versions

### Synopsis

Display all Go versions.
Example:
  gvm list
    Show all Go versions installed locally.
  gvm list -r
    Show all available Go versions remotely.

```
gvm list [flags]
```

### Options

```
  -h, --help          help for list
  -r, --remote        List remote Go versions
  -t, --type string   Version type (default all): stable | unstable | archived  (default "all")
```

### Options inherited from parent commands

```
  -v, --verbose   verbose output
```

### SEE ALSO

* [gvm](gvm.md)	 - Go version manager for installing and switching between multiple Go versions

