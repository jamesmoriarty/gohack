package main

import (
	"math"
	"time"
)

func main() {
	<-time.After(time.Duration(math.MaxInt64))
}
