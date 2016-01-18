package eqemuconfig

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig()
	if err != nil {
		t.Fatalf("Error loading config: %s", err.Error())
	}
	t.Fatal(c)
}
