package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// URLを引数で受取り、内容を表示する
func main() {
	const URL_PREFIX = "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, URL_PREFIX) {
			url = URL_PREFIX + url
		}

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("%s", err)
			continue
		}

		b, err := ioutil.ReadAll(resp.Body)
		s := resp.Status                            // Bodyは可変長なのでio.Readerで受ける必要があるが、statusは決まっているのでstringで問題ない。
		resp.Body.Close()                           // () をつけずにresp.Body.Close　を実行したらerror : ./fetch.go:21: resp.Body.Close evaluated but not used
		fmt.Printf("status: %s \nbody  : %s", s, b) // []byteだがstringにcastされる
	}
}

// TODO katsukamaru implement 1.7
