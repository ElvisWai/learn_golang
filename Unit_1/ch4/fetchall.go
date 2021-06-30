// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// go这个关键字创建一个goroutine，并且让函数在这个goroutine异步执行http.Get方法
		// main函数本身也运行在一个goroutine中，而go function则表示创建一个新的goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
	//（译注：可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据），
	// 因为我们需要这个方法返回的字节数，但是又不想要其内容。
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs   %7d   %s", secs, nbytes, url)
}
