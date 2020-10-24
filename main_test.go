package main

import (
	"os"
	"testing"
)

func TestGetInterval(t *testing.T) {
	// Env BAR is not set
	omitted := getInterval("BAR")
	if omitted != "60m" {
		t.Errorf("Expect 60m as default if BAR is not set")
	}

	// Env FOO is set to "10"
	os.Setenv("FOO", "10")
	interval := getInterval("FOO")
	if interval != "10m" {
		t.Errorf("Expect %vm, get %v", interval, interval)
	}

	// Env FOO is set to "10m"
	os.Setenv("FOO", "10m")
	interval = getInterval("FOO")
	if interval != "10m" {
		t.Errorf("Expect %v, get %v", os.Getenv("FOO"), interval)
	}
}
