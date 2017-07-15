package main

import "net"

type Tunnel struct {
	LocalAddr  string
	RemoteAddr string
	callback   func(net.Addr) `toml:"-"`
}
