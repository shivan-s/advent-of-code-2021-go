package main

import (
	"adventofcode2021/day02"
	"fmt"
)

func main() {
	// day 02
	// part 1
	sample_file := "./day02/sample_input_01.txt"
	file := "./day02/input.txt"
	sample_ans_day02_part1, _ := day02.ProductDepthHorizontalPosition(sample_file)
	ans_day02_part1, _ := day02.ProductDepthHorizontalPosition(file)
	fmt.Println("Day 02 - Part 1:")
	fmt.Printf("Result from sample input: %d\n", sample_ans_day02_part1)
	fmt.Printf("Result from real input: %d\n", ans_day02_part1)

	// part 2
	sample_ans_day02_part2, _ := day02.ProductDepthAimHorizontalPosition(sample_file)
	ans_day02_part2, _ := day02.ProductDepthAimHorizontalPosition(file)
	fmt.Println("Day 02 - Part 2:")
	fmt.Printf("Result from sample input: %d\n", sample_ans_day02_part2)
	fmt.Printf("Result from real input: %d\n", ans_day02_part2)
}
