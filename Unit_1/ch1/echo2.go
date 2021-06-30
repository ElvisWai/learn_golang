//Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//_下标，rg值
	// range (go) = enumerate (python)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "

	}
	fmt.Println(s)

}

//s := ""                是一条短变量声明，最简洁，但只能用在函数内部，而不能用于包变量。
//var s string           依赖于字符串的默认初始化零值机制，被初始化为""。
//var s = ""             用得很少，除非同时声明多个变量。
//var s string = ""		 显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了。
