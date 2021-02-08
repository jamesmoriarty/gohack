package gohack

import (
	"fmt"
	"strconv"
)

func ToHexString(ptr uintptr) string {
	s := fmt.Sprintf("%d", ptr)
	n, _ := strconv.Atoi(s)
	h := fmt.Sprintf("0x%x", n)

	return h
}
