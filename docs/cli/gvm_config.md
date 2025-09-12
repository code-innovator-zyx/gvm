## gvm config

Manage gvm configuration

### Synopsis

View and modify gvm configuration.

Examples:
  gvm config list           # show all config
  gvm config set mirrors https://mirrors.aliyun.com/golang/
  gvm config get mirrors    # show a specific config
  gvm config unset mirrors  # remove a config item

### Options

```
  -h, --help   help for config
```

### Options inherited from parent commands

```
  -v, --verbose   verbose output
```

### SEE ALSO

* [gvm](gvm.md)	 - Go version manager for installing and switching between multiple Go versions
* [gvm config get](gvm_config_get.md)	 - Get a configuration value
* [gvm config list](gvm_config_list.md)	 - List all configuration values
* [gvm config set](gvm_config_set.md)	 - Set a configuration value
* [gvm config unset](gvm_config_unset.md)	 - Remove a configuration value

