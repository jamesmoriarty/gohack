package main

import (
	"testing"
)

func TestNoProcess(t *testing.T) {
	_, _, err := instrument()

	got := err.Error()

	if got != "Failed to get pid csgo.exe" {
		t.Errorf(got)
	}
}
