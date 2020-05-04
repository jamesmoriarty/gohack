package gohack

import (
	"os/exec"
	"strings"
	"testing"
)

func withProcess(path string, f func()) {
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
	_, err := Instrument()

	got := err.Error()

	want := "Failed to get pid csgo.exe"

	if got != want {
		t.Errorf("%q; want %q", got, want)
	}
}

func TestStubProcessNoDLL(t *testing.T) {
	withProcess("test\\nodll\\csgo.exe", func() {
		_, err := Instrument()

		got := err.Error()

		want := "Failed to get module offset"

		if got != want {
			t.Errorf("%q; want %q", got, want)
		}
	})
}

func TestStubProcess(t *testing.T) {
	withProcess("test\\dll\\csgo.exe", func() {
		_, err := Instrument()

		got := err.Error()

		want := "Failed to get player offset"

		if got != want {
			t.Errorf("%q; want %q", got, want)
		}
	})
}
