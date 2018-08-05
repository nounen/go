package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
)

// 浅尝辄止地来体验一下 Go 语言里的 goroutine 和 channel
// 普通版 和 goroutine 版本 费时比较
// `go run 1.6.1.fetch_go.go https://www.baidu.com https://www.oschina.net https://segmentfault.com/`
func main() {
	withoutGo()
	fmt.Println()
	byGo()

	/**
	0.28s      227  https://www.baidu.com
	0.25s   160580  https://www.oschina.net
	0.43s    83079  https://segmentfault.com/
	fetch 普通版共费时: 0.96s elapsed

	0.02s      227  https://www.baidu.com
	0.10s   160580  https://www.oschina.net
	0.22s    83143  https://segmentfault.com/
	fetch goroutine 版共费时: 0.22s elapsed
	 */
}

func byGo()  {
	start := time.Now()
	ch := make(chan string)
	args := os.Args[1:]

	for _, url := range args{
		// start a goroutine
		go fetchGo(url, ch)
	}

	for range args{
		// receive from channel ch
		fmt.Println(<-ch)
	}

	fmt.Printf("fetch goroutine 版共费时: %.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch goroutine 版本
func fetchGo(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		// send to channel ch
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func withoutGo()  {
	start := time.Now()
	args := os.Args[1:]

	for _, url := range args{
		fetchWithoutGo(url)
	}

	fmt.Printf("fetch 普通版共费时: %.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch 普通版本
func fetchWithoutGo(url string)  {
	start := time.Now()
	resp, _ := http.Get(url)

	nbytes, _ := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs  %7d  %s\n", secs, nbytes, url)
}
