package main

import (
	"C" // https://github.com/golang/go/issues/22192
	"fmt"
)

//export Sum
func Sum(arg1, arg2 int32) int32 {
	return arg1 + arg2
}

//export HelloWorld
func HelloWorld() {
	fmt.Println("HelloWorld")
}

func main() {

}
