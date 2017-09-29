# Tunneler
Tunneler is a ssh port forwarder with TOML configuration management.

## How to install
```bash
# We need a Go language compiler and tools (1.8 or later).
$ go get -u github.com/minoritea/tunneler
```

## How to use
```bash
$ tunneler -c config.toml
```

## Configuration
Tunneler uses [TOML](https://github.com/toml-lang/toml) for configuration file format.

[https://github.com/toml-lang/toml](https://github.com/toml-lang/toml)

```toml
# Each top-level keys corresponds each servers which we want to have them forward connections via SSH.
# And each entries must includes connection configurations.
# We can name each entries at will.
# For example, a configuration for a server which named `bastion` is below.
[bastion]
# the server's IP and port
host = "192.168.1.1"
port = "22"

# the login user's name
user = "remote_user" 

# Tunneler currently supports cert file authentication using PEM format.
# Cert files must be placed in our local machine.
cert_path = "/home/local_user/.ssh/cert.pem"

# `server`.tunnels is a table of settings for port forwarding targets.
# We can name each keys at will.
[bastion.tunnels.postgres]
# the target server's IP and port
remote_host = "192.168.10.1"
remote_port = "5432"
# the forwarded local port
local_port  = "5432"

# We can adds multiple entries.
```

### Multi hop SSH tunneling
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

# We can set multi stage intermediate servers.
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

### Resolving hostnames on hosts
Tunneler can resolve hostnames on each forwarding servers.
```toml
[bastion]
# ...

[bastion.cascades.foo_example]
host = "foo.example.org"
port = "22"
user = "foo_user"
cert_path = "/home/local_user/.ssh/cert.pem"
# If we want to resolve hostnames on a forwarding server,
# enable `resolve_on_host`.
# In this example, the host `foo.example.org` is resolved on `bastion`.
resolve_on_host = true

[bastion.cascades.foo_example.tunnels.bar_example]
host = "bar.example.org"
port = "80"
user = "bar_user"
cert_path = "/home/local_user/.ssh/cert.pem"
# We can also resolve hostnames in tunnels.
# Resolving process will run on the server which forwards the tunnel.
# In this example, the host `bar.example.org` is resolved on `bastion`.
resolve_on_host = true
```

## LICENSE
MIT License (see the attached file: LICENSE)

Copyright (c) 2017 Minori Tokuda
