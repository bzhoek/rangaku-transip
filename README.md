# Update TransIP DNS record with external IP

DynDNS like functionality with TransIP DNS API. You need to enable API access first.

    go run dns.go <username> <path to private key> <domain> <hostname>

# Build

  export GOPATH="$HOME/.go"
  mkdir $GOPATH
  go get # gives: no install location for directory, but installs just the same
  go build

`go get -u all` installs in default `~/go/` `GOPATH`.  

## Documentation

https://golang.org/doc/code.html

https://github.com/transip/gotransip and doc https://godoc.org/github.com/transip/gotransip
