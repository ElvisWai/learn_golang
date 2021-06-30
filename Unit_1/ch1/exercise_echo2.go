// exercise_echo2 print index and vlaue for command-line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for index, value := range os.Args[1:] {
		s += (sep + strconv.Itoa(index) + "  " + value)
		sep = "\n"
	}

	fmt.Println(s)

}

// int到string
//string:=strconv.Itoa(int)

//string到int
//int,err:=strconv.Atoi(string)
