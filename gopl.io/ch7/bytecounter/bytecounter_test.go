package bytecounter

import (
	"testing"
)

func TestByteCounter_Write(t *testing.T) {
	inputs := []struct {
		input    []byte
		expected int
	}{
		{[]byte(""),      0},
		{[]byte("bar"),   3},
		{[]byte("hello"), 5},
		{[]byte("こんにちは"), 15},
	}
	var c ByteCounter
	for _, i := range inputs {
		result,_ := c.Write(i.input)

		if i.expected != result {
			t.Errorf("expected %v, but %v", i.expected, result)
		}
	}
}