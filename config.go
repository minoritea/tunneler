package main

import (
	"fmt"
	"github.com/naoina/toml"
	"io/ioutil"
	"net"
)

type BastionConfig struct {
	Host          string
	Port          string
	User          string
	CertPath      string
	Tunnels       map[string]Tunnel
	Cascades      map[string]BastionConfig
	ResolveOnHost bool
}

type Tunnel struct {
	LocalHost     string
	LocalPort     string
	RemoteHost    string
	RemotePort    string
	ResolveOnHost bool
	callback      func(net.Addr) `toml:"-"`
}

func LoadConfig(confpath string, printconf bool) (map[string]BastionConfig, error) {
	config := make(map[string]BastionConfig)
	data, err := ioutil.ReadFile(confpath)
	if err != nil {
		return config, err
	}
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	if printconf {
		fmt.Println(string(data))
	}
	return config, nil
}
