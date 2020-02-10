package main

import (
	"os/exec"
	"testing"
)

func withCSGOExe(path string, f func()) {
	s := []string{"cmd.exe", "/C", "start", path}

	cmd := exec.Command(s[0], s[1:]...)

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	defer exec.Command("TASKKILL", "/F", "/IM", "csgo.exe").Run()

	f()
}

func TestNoProcess(t *testing.T) {
	_, _, err := instrument()

	got := err.Error()

	if got != "Failed to get pid csgo.exe" {
		t.Errorf(got)
	}
}

func TestStubProcessNoDLL(t *testing.T) {
	withCSGOExe("test\\nodll\\csgo.exe", func() {
		_, _, err := instrument()

		got := err.Error()

		if got != "Failed to get module address client_panorama.dll" {
			t.Errorf(got)
		}
	})
}

func TestStubProcess(t *testing.T) {
	withCSGOExe("test\\dll\\csgo.exe", func() {
		_, _, err := instrument()

		got := err.Error()

		if got != "Failed to get LocalPlayer address" {
			t.Errorf(got)
		}
	})
}
