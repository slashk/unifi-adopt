# unifi-adopt


## Usage

```bash
Monitor your Ubiqiti WAPs to ensure that stay connected to your Unifi server.
This tool will read a config file (defaults to ~/.unifi-adopt) and query each WAP configured 
by SSHing to it and checking it's config. If it is not connected

Usage:
  unifi-adopt [flags]
  unifi-adopt [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Prints the current version

Flags:
      --config string   config file (default is $HOME/.unifi-adopt.yaml)
  -d, --debug           debug to see all network calls
  -h, --help            help for unifi-adopt
  -t, --toggle          Help message for toggle

Use "unifi-adopt [command] --help" for more information about a command.
```


### version

```bash
0.0.0-SNAPSHOT-none, commit none, built at 2022-11-25T19:43:43Z
```
