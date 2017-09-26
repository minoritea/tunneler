# Tunneler
Tunneler is a ssh port forwarder with TOML configuration management.

## How to install
```bash
# you need a Go language compiler and tools (1.8 or later).
$ go get -u github.com/minoritea/tunneler
```

## How to use
```bash
$ tunneler -c config.toml
```

## Configuration
Tunneler uses TOML for configuration file format.

```toml
# Each top-level keys are server names which forwards connections via SSH.
# And each entries must includes connection configurations.
# For example, a configuration for a server which named `bastion` is below.
[bastion]
# the server's IP and port
host = "192.168.1.1"
port = "22"
# the login user's name
user = "remote_user" 
# Tunneler currently supports cert file authentication using PEM format.
# Cert files must be placed in your local machine.
cert_path = "/home/local_user/.ssh/cert.pem"

# `server`.tunnels is a table of settings for port forwarding targets.
[bastion.tunnels.postgres]
# the target server's IP and port
remote_host = "192.168.10.1"
remote_port = "5432"
# the forwarded local port
local_port  = "5432"

# You can adds multiple entries.
```

## Multi hop SSH tunneling
Tunneler also supports multi hop SSH tunneling.

```toml
[bastion]
host = "192.168.1.1"
port = "22"
user = "remote_user" 
cert_path = "/home/local_user/.ssh/cert.pem"

# `server`.cascades is a table of settings for intermediate servers.
[bastion.cascades.server1]
host = "192.168.100.1"
port = "22"
user = "remote_user" 
cert_path = "/home/local_user/.ssh/cert.pem"

# You can set multi stage intermediate servers.
[bastion.cascades.server1.cascades.server2]
host = "192.168.100.2"
port = "22"
user = "remote_user" 
cert_path = "/home/local_user/.ssh/cert.pem"

[bastion.cascades.server1.cascades.server2.tunnels.postgres]
remote_host = "192.168.100.10"
remote_port = "5432"
local_port  = "5432"
```

## LICENSE
MIT License (see the attached file: LICENSE)

Copyright (c) 2017 Minori Tokuda
