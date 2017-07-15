package main

import (
	"github.com/naoina/toml"
	"io/ioutil"
	"net"
)

type BastionConfig struct {
	Host     string
	Port     string
	User     string
	CertPath string
	Tunnels  map[string]Tunnel
	Cascades map[string]BastionConfig
}

type Tunnel struct {
	LocalHost  string
	LocalPort  string
	RemoteHost string
	RemotePort string
	callback   func(net.Addr) `toml:"-"`
}

func LoadConfig(confpath string) (map[string]BastionConfig, error) {
	config := make(map[string]BastionConfig)
	data, err := ioutil.ReadFile(confpath)
	if err != nil {
		return config, err
	}
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
