package util

import (
	log "github.com/sirupsen/logrus"
)

func NeverExit(f func()) {
	defer func() {
		if v := recover(); v != nil {
			log.Fatal("Goroutine failed. Recovering...")

			go NeverExit(f)
		}
	}()

	f()
}
