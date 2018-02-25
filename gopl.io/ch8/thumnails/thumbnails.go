package thumnails

import "os"

func openFile(filename string) error {
	_, e := os.Open(filename)
	return e
}
