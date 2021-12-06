package main

import (
	"adventofcode2021/day02"
	"adventofcode2021/day03"
	"adventofcode2021/day04"
	"adventofcode2021/day06"
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
	fmt.Println("Day 03 - Part 1:")
	fmt.Printf("Result from sample input: %d\n", sample_ans_day03_part1)
	fmt.Printf("Result from real input: %d\n", ans_day03_part1)

	// part 2
	sample_ans_day03_part2 := day03.CalculateOxygenCO2(sample_file_day03)
	ans_day03_part2 := day03.CalculateOxygenCO2(file_day03)
	fmt.Println("Day 03 - Part 2:")
	fmt.Printf("Result from sample input: %d\n", sample_ans_day03_part2)
	fmt.Printf("Result from real input: %d\n", ans_day03_part2)
}

func main_day04() {
	// day04
	// part 1
	sample_file_day04_draw := "./day04/sample_input_draw.txt"
	sample_file_day04 := "./day04/sample_input.txt"
	file_day04_draw := "./day04/input_draw.txt"
	file_day04 := "./day04/input.txt"

	sample_ans := day04.WinnerBingo(sample_file_day04, sample_file_day04_draw)
	ans := day04.WinnerBingo(file_day04, file_day04_draw)

	fmt.Println("Day 04 - Part 1:")
	fmt.Printf("Result from sample input: %v\n", sample_ans)
	fmt.Printf("Result from input: %v\n", ans)

	// part 2
}

// func main_day05() {
// 	// day05
// 	sample_file := "sample_input.txt"
// 	file := "input.txt"
// 	// part 1
// 	// TODO: create functions
// }

func main_day06() {
	// day06
	sample_file := "./day06/sample_input.txt"
	file := "./day06/input.txt"
	// part 1
	sample_ans := day06.SimulateFish(sample_file, 80)
	ans := day06.SimulateFish(file, 80)
	fmt.Println("Day 06 - Part 1:")
	fmt.Printf("Result from sample input: %v\n", sample_ans)
	fmt.Printf("Result from input: %v\n", ans)

	// part 2
	p2_sample_ans := day06.SimulateFish(sample_file, 256)
	p2_ans := day06.SimulateFish(file, 256)
	fmt.Println("Day 06 - Part 2:")
	fmt.Printf("Result from sample input: %v\n", p2_sample_ans)
	fmt.Printf("Result from input: %v\n", p2_ans)
}

func main() {
	main_day06()
}
