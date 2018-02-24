package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: gray; fill:white; stroke-width: 0.7' width='%d'>", width, height)
}

func someFunc(input string) (string, error) {
	_, err := fmt.Println(input)
	return "", err
}
