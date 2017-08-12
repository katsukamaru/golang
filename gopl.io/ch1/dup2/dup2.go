package main

import (
	"bufio"
	"fmt"
	"os"
)

// ファイルから読み込むようにし、複数回同じ行が登場したら、行番号とファイル名も表示させる
func main() {
	counts := map[string]int{} // TODO katsukamaru make(map[string]int)との違いは?
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				// log.Fatal(err)  -> こうしてもいいけど、filesは複数のファイルを想定している。できるだけ実行して失敗したものをログに出す方がいい
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// defer file.Close() -> Possibly leak ... というwarnが出る。確かにloop内でdeferを書くよりは、このwarnを見てちゃんとf.Closeを呼ぶべき。
			countLines(file, counts)
			file.Close()
		}
	}

	for i, v := range counts {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, i)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
