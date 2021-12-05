package day03

import (
	"adventofcode2021/utils"
	"strconv"
	"strings"
)

// find the most common binaries - gamma
// epsilon will be the opposite to gamma
// convert to decimal and multiple together

func CalculatePowerConsumption(file string) int64 {
	diagnostic_report, _ := utils.ReadDataFromFile(file)
	l := len(diagnostic_report[0])
	num_of_ones := make([]int, l) // counting the bits

	for _, result := range diagnostic_report { // results = 00100
		for i := 0; i < len(result); i++ {
			if result[i] == 49 { // 49 is "1"
				num_of_ones[i]++
			}
		}
	}
	total_diagnostic_report := len(diagnostic_report)
	gamma_array := make([]string, l)
	epsilon_array := make([]string, l)
	for i, num := range num_of_ones {
		if num > (total_diagnostic_report / 2) {
			gamma_array[i] = "1"
			epsilon_array[i] = "0"
		} else {
			gamma_array[i] = "0"
			epsilon_array[i] = "1"
		}
	}

	gamma_string := strings.Join(gamma_array[:], "")
	epsilon_string := strings.Join(epsilon_array[:], "")

	gamma_binary, _ := strconv.ParseInt(gamma_string, 2, 64)
	epsilon_binary, _ := strconv.ParseInt(epsilon_string, 2, 64)

	product := gamma_binary * epsilon_binary

	return product
}

// seperate into Oxygen
// create a slice for oxygen
// then do through slice and keep the values needed in another slice
// iterate through this slice again and repeat the above 2 steps
func CalculateOxygenCO2(file string) int64 {
	diagnostic_report, _ := utils.ReadDataFromFile(file)
	oxygen_slice := diagnostic_report
	CO2_slice := diagnostic_report

	// Oxygen
	element_count := 0
	for i, _ := range diagnostic_report[0] {
		for _, number := range oxygen_slice {
			if number[i] == 49 {
				element_count++ // counting the 1's
			}

			temp_oxygen_slice := oxygen_slice
			for j, number := range temp_oxygen_slice {
				if element_count >= (len(diagnostic_report) / 2) { // checks if 1s are majority or if there is a tie
					if number[i] == 48 { // removing 0
						oxygen_slice[j] = oxygen_slice[len(oxygen_slice)-1]
						oxygen_slice = oxygen_slice[:len(oxygen_slice)-1]
					}
				} else {
					if number[i] == 49 { // removing 1
						oxygen_slice[j] = oxygen_slice[len(oxygen_slice)-1]
						oxygen_slice = oxygen_slice[:len(oxygen_slice)-1]
					}
				}
			}
			temp_CO2_slice := CO2_slice
			for j, number := range temp_CO2_slice {
				if element_count <= (len(diagnostic_report) / 2) { // checks if 1s are minotiry or if there is a tie
					if number[i] == 48 { // removing 1
						CO2_slice[j] = CO2_slice[len(CO2_slice)-1]
						CO2_slice = CO2_slice[:len(CO2_slice)-1]
					}
				} else {
					if number[i] == 49 { // removing 0
						CO2_slice[j] = CO2_slice[len(CO2_slice)-1]
						CO2_slice = CO2_slice[:len(CO2_slice)-1]
					}
				}
			}
		}
	}
	oxygen_binary, _ := strconv.ParseInt(oxygen_slice[0], 2, 64)
	CO2_binary, _ := strconv.ParseInt(CO2_slice[0], 2, 64)

	product := oxygen_binary * CO2_binary

	return product
}
