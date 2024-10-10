# DNS Resolver
---
Simple DNS resolver written in Go.

# Installation
---
Clone the repository
```bash
git clone git@github.com:FranChesK0/dns-resolver.git
```

Build app
```bash
make
```

# Usage
---
```bash
./dns-resolver google.com
```

It is also possible to resolve multiple domains
```bash
./dns-resolver google.com youtube.com
```

You can specify name server that is using for resolving (default is `77.240.157.30`)
```bash
./dns-resolver google.com --name-server 204.106.240.53
```