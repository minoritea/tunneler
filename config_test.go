package main

import "testing"

func TestLoadConfig(t *testing.T) {
	_, err := LoadConfig("test.toml")
	if err != nil {
		t.Error(err)
	}
}
