package main

import (
	"testing"
)

func Test_basedir1(t *testing.T) {
	/* ------------------- */
	/*       Prepare       */
	/* ------------------- */
	candidates := []struct {
		input    string
		expected string
	}{
		{"foo/bar.go", "bar"},
		{"/foo/bar.go", "bar"},
		{"foo/bar.piyo.go", "bar.piyo"},
	}

	/* ------------------- */
	/*       Assert        */
	/* ------------------- */
	for _, v := range candidates {
		result := basedir(v.input)
		if result != v.expected {
			t.Errorf("Expected %v but %v", v.expected, result)
		}
	}
}
