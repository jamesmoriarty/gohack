package main

 // https://github.com/golang/go/issues/22192
import (
	"fmt"
	"math"
	"syscall"
	"time"
)

var (
	dll  = syscall.MustLoadDLL("client_panorama.dll")
	proc = dll.MustFindProc("Sum")
)

func main() {
	fmt.Printf("dll=%v proc=%v", dll, proc)
	proc.Call(1, 2)
	<-time.After(time.Duration(math.MaxInt64))

}
