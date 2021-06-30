package main

import "fmt"

func main() {
	fmt.Println("Hello,World!(你好，世界！)")
}

//main包比较特殊。它定义了一个独立可执行的程序，而不是一个库。
//在main里的main 函数 也很特殊，它是整个程序执行时的入口（译注：C系语言差不多都这样）。
//main函数所做的事情就是程序做的。
