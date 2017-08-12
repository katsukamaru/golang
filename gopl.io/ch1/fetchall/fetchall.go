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
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	s := resp.Status
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // TODO katsumaru %vとはなんぞや
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s %s", secs, nbytes, s, url)
}

// 1.10
// よくわからない。どうすればキャッシュされているということが言えるのか。
// リクエストが完全に同じであること？statusCodeは200で、キャッシュされているような様子はない。

// 1.11
// timeoutの設定を入れればOK
//for range os.Args[1:] {
//	select {
//	case <-ch:
//		fmt.Println(<-ch)
//	case <-time.After(time.Second * 3):  // この実装ではだめ。channelをロックしてしまう。
//		fmt.Println("timeout")
//	}
//}
// これで大丈夫かと思ったが、実際にはchannelをロックしてしまう。この項は9章で再度勉強するのでスルーする。
//
