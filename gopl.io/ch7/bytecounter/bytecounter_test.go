package bytecounter

import (
	"testing"
)

func TestByteCounter_Write(t *testing.T) {
	candidates := []struct {
		input    []byte
		expected int
	}{
		{[]byte(""), 0},
		{[]byte("bar"), 3},
		{[]byte("hello"), 5},
		{[]byte("こんにちは"), 15},
	}

	var c ByteCounter
	for _, i := range candidates {
		result, _ := c.Write(i.input)

		if i.expected != result {
			t.Errorf("expected %v, but %v", i.expected, result)
		}
	}
}

func TestWordCounter_Count(t *testing.T) {
	candidates := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"foo bar", 2},
		{"foobar", 1},
		{"foo  bar", 2},
		{"foo  bar fuga", 3},
		{"foo  bar fuga  piyo", 4},
		{"foo  bar fuga  piyo  ", 4},
	}
	var c WordCounter
	for _, i := range candidates {
		result, _ := c.Count(i.input)

		if i.expected != result {
			t.Errorf("expected %v, but %v", i.expected, result)
		}
	}
}

func TestLineCounter_Count(t *testing.T) {
	candidates := []struct {
		input    string
		expected int
	}{
		{"foo bar", 1},
		{"foobar\n", 2},
		{"", 0},
	}

	var c LineCounter
	for _, i := range candidates {
		result, _ := c.Count(i.input)

		if i.expected != result {
			t.Errorf("expected %v, but %v", i.expected, result)
		}
	}

}
