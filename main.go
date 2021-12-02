package main

import (
	"adventofcode2021/day02"
	"fmt"
)

func main() {
	// TODO: reference the previous function
	sample_ans_day02, _ := day02.ProductDepthHorizontalPosition("./day02/sample_input_01.txt")
	ans_day02, _ := day02.ProductDepthHorizontalPosition("./day02/input.txt")
	fmt.Printf("Result from sample input: %d\n", sample_ans_day02)
	fmt.Printf("Result from real input: %d\n", ans_day02)
}
