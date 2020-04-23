package gohack

import (
	"os/exec"
	"strings"
	"testing"
)

func withEXE(path string, f func()) {
	s := []string{"cmd.exe", "/C", "start", path}

	cmd := exec.Command(s[0], s[1:]...)

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	ss := strings.Split(path, "\\")
	exe := ss[len(ss)-1]

	defer exec.Command("TASKKILL", "/F", "/IM", exe).Run()

	f()
}

func TestNoProcess(t *testing.T) {
	_, _, err := Instrument()

	got := err.Error()

	if got != "Failed to get pid csgo.exe" {
		t.Errorf(got)
	}
}

func TestStubProcessNoDLL(t *testing.T) {
	withEXE("test\\nodll\\csgo.exe", func() {
		_, _, err := Instrument()

		got := err.Error()

		if got != "Failed to get module address client_panorama.dll" {
			t.Errorf(got)
		}
	})
}

func TestStubProcess(t *testing.T) {
	withEXE("test\\dll\\csgo.exe", func() {
		_, _, err := Instrument()

		got := err.Error()

		if got != "Failed to get LocalPlayer address" {
			t.Errorf(got)
		}
	})
}
