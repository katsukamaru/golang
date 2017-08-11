package bytecounter

import (
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// 7.1
type WordCounter int

func (c *WordCounter) Count(word string) (int, error) {
	split := strings.Fields(word)
	return len(split), nil
}

type LineCounter int

func (c *LineCounter) Count(word string) (int, error) {
	if word == "" {
		return 0, nil
	}
	sp := strings.Split(word, "\n")
	return len(sp), nil
}
