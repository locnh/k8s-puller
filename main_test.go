package main

import (
	"os"
	"testing"
)

func TestGetInterval(t *testing.T) {
	// Env BAR is not set
	omitted := getConfigInterval("BAR")
	if omitted != "60m" {
		t.Errorf("Expect 60m as default if BAR is not set")
	}

	// Env FOO is set to "10"
	os.Setenv("FOO", "10")
	interval := getConfigInterval("FOO")
	if interval != "10m" {
		t.Errorf("Expect %vm, get %v", interval, interval)
	}

	// Env FOO is set to "10m"
	os.Setenv("FOO", "10m")
	interval = getConfigInterval("FOO")
	if interval != "10m" {
		t.Errorf("Expect %v, get %v", os.Getenv("FOO"), interval)
	}
}

func TestGetJSONLog(t *testing.T) {
	// Env JSONLOG is not set
	result := getConfigLogFormat("JSONLOG")
	if result != false {
		t.Errorf("Expect FALSE as default if JSONLOG is not set")
	}

	// Env JSONLOG is set to "1"
	os.Setenv("JSONLOG", "1")
	result = getConfigLogFormat("JSONLOG")
	if result != true {
		t.Errorf("Expect TRUE as JSONLOG is set to \"1\"")
	}

	// Env JSONLOG is set to "true"
	os.Setenv("JSONLOG", "true")
	result = getConfigLogFormat("JSONLOG")
	if result != true {
		t.Errorf("Expect TRUE as JSONLOG is set to \"true\"")
	}

	// Env JSONLOG set to anything except "1" or "true"
	os.Setenv("JSONLOG", "whatever")
	result = getConfigLogFormat("JSONLOG")
	if result != false {
		t.Errorf("Expect FALSE as JSONLOG is not set to \"1\" or \"true\"")
	}
}
