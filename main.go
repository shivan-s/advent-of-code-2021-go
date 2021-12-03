package main

import (
	"adventofcode2021/day02"
	"adventofcode2021/day03"
	"fmt"
)

func main_day02() {
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

func main_day03() {
	// day03
	// part 1
	sample_file_day03 := "./day03/sample_input.txt"
	file_day03 := "./day03/input.txt"
	sample_ans_day03_part1 := day03.CalculatePowerConsumption(sample_file_day03)
	ans_day03_part1 := day03.CalculatePowerConsumption(file_day03)
	fmt.Println("Day 03 - Part 3:")
	fmt.Printf("Result from sample input: %d\n", sample_ans_day03_part1)
	fmt.Printf("Result from real input: %d\n", ans_day03_part1)

	// part 2
}

func main() {
	main_day03()
}
