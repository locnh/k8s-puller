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

func TestGetImages(t *testing.T) {
	// INPUT: Env IMAGES is not set"
	// EXPECTED OUTPUT, ok is false
	var expected []string
	result, ok := getImages("IMAGES")
	if ok {
		t.Errorf("The \"ok\" flag should be \"false\"")
	}
	if !testEqual(result, expected) {
		t.Errorf("Empty input test, Result is not as expected")
	}

	// INPUT: Env IMAGES is set to 1 image without tag "busybox"
	os.Setenv("IMAGES", "busybox")
	// EXPECTED OUTPUT
	expected = []string{"busybox"}
	if result, ok := getImages("IMAGES"); ok {
		if !testEqual(result, expected) {
			t.Errorf("Single image test, Result is not as expected")
		}
	} else {
		t.Errorf("IMAGES un-parsable")
	}

	// INPUT: Env IMAGES is set to multiple images with tag "busybox:latest,alpine:latest"
	os.Setenv("IMAGES", "busybox:latest,alpine:latest")
	// EXPECTED OUTPUT
	expected = []string{"busybox:latest", "alpine:latest"}
	if result, ok := getImages("IMAGES"); ok {
		if !testEqual(result, expected) {
			t.Errorf("Multiple image with tag, Result is not as expected")
		}
	} else {
		t.Errorf("IMAGES un-parsable")
	}
}

func testEqual(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
