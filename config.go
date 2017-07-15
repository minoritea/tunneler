package main

import (
	"fmt"
	"github.com/naoina/toml"
	"io/ioutil"
	"sync"
)

type BastionConfig struct {
	Host     string
	User     string
	CertPath string
	Tunnels  map[string]Tunnel
	Cascades BastionConfigs
}

type BastionConfigs map[string]BastionConfig

type Config struct {
	BastionConfigs
	errch chan error `toml:"-"`
}

func handleError(errch chan error) {
	for err := range errch {
		fmt.Printf("%+v\n", err)
	}
}

func LoadConfig(confpath string) (Config, error) {
	config := Config{make(BastionConfigs), make(chan error)}
	data, err := ioutil.ReadFile(confpath)
	if err != nil {
		return config, err
	}
	err = toml.Unmarshal(data, &config.BastionConfigs)
	if err != nil {
		return config, err
	}
	return config, nil
}

func (bc BastionConfig) start(wg *sync.WaitGroup, errch chan error) {
	defer wg.Done()
	b, err := NewBastion(bc, errch)
	if err != nil {
		errch <- err
		return
	}
	defer b.Close()
	b.Run(bc)
}

func (config Config) Execute() {
	go handleError(config.errch)
	wg := new(sync.WaitGroup)
	for _, bc := range config.BastionConfigs {
		wg.Add(1)
		go bc.start(wg, config.errch)
	}
	wg.Wait()
}
