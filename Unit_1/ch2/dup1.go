// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 控制循环退出
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// map存储了键/值（key/value）的集合，对集合元素，提供常数时间的存、取或测试操作。
// input.Scan()，即读入下一行，并移除行末的换行符。Scan函数在读到一行时返回true，不再有输入时返回false。

// ===================  Printf转换  =======================
// %d          十进制整数
// %x, %o, %b  十六进制，八进制，二进制整数。
// %f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
// %t          布尔：true或false
// %c          字符（rune） (Unicode码点)
// %s          字符串
// %q          带双引号的字符串"abc"或带单引号的字符'c'
// %v          变量的自然形式（natural format）
// %T          变量的类型
// %%          字面上的百分号标志（无操作数）
