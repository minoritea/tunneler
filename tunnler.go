package main

import (
	"fmt"
	"sync"
	"time"
)

func handleError(errch chan error) {
	for err := range errch {
		fmt.Printf("%+v\n", err)
	}
}

func execute(config map[string]BastionConfig) {
	errch := make(chan error)
	go handleError(errch)
	wg := new(sync.WaitGroup)
	for _, bc := range config {
		wg.Add(1)
		go start(bc, wg, errch)
	}
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
}

func start(bc BastionConfig, wg *sync.WaitGroup, errch chan error) {
	defer wg.Done()
	b, err := NewBastion(bc, errch)
	if err != nil {
		errch <- err
		return
	}
	defer b.Close()
	b.Up()
}
