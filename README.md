# unifi-adopt

## Usage

```bash
Monitor your Ubiqiti WAPs to ensure that stay connected to your Unifi server.
This tool will read a config file (defaults to ~/.unifi-adopt) and query each WAP configured
by SSHing to it and set the inform-url to your configuration.

Usage:
  unifi-adopt [flags]
  unifi-adopt [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Prints the current version

Flags:
  -c, --config string   config file (default "$HOME/.unifi-adopt")
  -d, --debug           debug to see all network calls
  -h, --help            help for unifi-adopt
  -p, --pushover        send pushover messages on all actions
  -t, --toggle          Help message for toggle

Use "unifi-adopt [command] --help" for more information about a command.
```

### version

```bash
1.2.0-SNAPSHOT-00c308e, commit 00c308e, built at 2025-07-06T18:00:40Z
```
