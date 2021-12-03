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

// seperate into Oxygen and CO2
func CalculateOxygenCO2(file string) int64 {
	diagnostic_report, _ := utils.ReadDataFromFile(file)
	oxy_diagnostic_report := diagnostic_report
	co_diagnostic_report := diagnostic_report
	var element_count int

	// Oxygen
	for n := 0; n < len(diagnostic_report[0]; n++) {
		for _, result := range diagnostic_report {
			if result[n] == 49 { // 1
				element_count++
			}
		}
		if element_count >= (len(oxy_diagnostic_report) / 2) {
			// keep the "1s" remove the 0s 
			for i, result := range oxy_diagnostic_report {
				if result[n] == 49 {
				oxy_diagnostic_report = append(oxy_diagnostic_report[:i], oxy_diagnostic_report[i+1:]...)
				}
			}
		} else {
			// keep the 0s, remove the 1s
			for i, result := range oxy_diagnostic_report {
				if result[n] == 48 {
				oxy_diagnostic_report = append(oxy_diagnostic_report[:i], oxy_diagnostic_report[i+1:]...)
				}
			}
		}
	// CO2
	for n := 0; n < len(diagnostic_report[0]; n++) {
		for _, result := range co_diagnostic_report {
			if result[n] == 49 { // 1
				element_count++
			}
		}

		if element_count >= (len(co_diagnostic_report) / 2) {
			// keep the "1s" remove the 0s 
			for i, result := range co_diagnostic_report {
				if result[n] == 49 {
				co_diagnostic_report = append(co_diagnostic_report[:i], co_diagnostic_report[i+1:]...)
				}
			}
		} else {
			// keep the 0s, remove the 1s
			for i, result := range co_diagnostic_report {
				if result[n] == 48 {
				co_diagnostic_report = append(co_diagnostic_report[:i], co_diagnostic_report[i+1:]...)
				}
			}
		}
	}

	oxy_diagnostic_report 
}