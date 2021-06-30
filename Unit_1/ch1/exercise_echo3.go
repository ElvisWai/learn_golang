package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var rangeStart = time.Now()
	s, sep := "", ""
	//_下标，rg值
	// range (go) = enumerate (python)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(time.Since(rangeStart))

	var joinStart = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(joinStart))
}
