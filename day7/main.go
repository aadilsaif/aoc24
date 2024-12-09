package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calibration struct {
	Total int
	Parts []int
}

func ParseInput(file string) []Calibration {
	lines := strings.Split(file, "\n")
	parsed := make([]Calibration, 0, len(lines))

	for _, line := range lines {
		split := strings.Split(line, ": ")
		total := split[0]
		key, err := strconv.Atoi(total)
		if err != nil {
			panic(err)
		}

		parts := strings.Split(split[1], " ")
		calibration := Calibration{key, make([]int, 0, len(parts))}
		for _, part := range parts {
			unit, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			calibration.Parts = append(calibration.Parts, unit)
		}

		parsed = append(parsed, calibration)
	}

	return parsed
}

func Reverse(s []int) []int {
	size := len(s)
	opposite := make([]int, size)

	for i, e := range s {
		opposite[size-1-i] = e
	}

	return opposite
}

/* func canListResultInTotal(expected int, parts []int) bool {
	if expected < 0 {
		// We cant subtract or divide a negative number by a positive one and result in a positive number - we can short circuit our search here
		return false
	}

	if len(parts) == 0 {
		return expected == 0
	}

	current := parts[0]
	if next := expected / current; expected == next*current && canListResultInTotal(next, parts[1:]) {
		return true
	}

	return canListResultInTotal(expected-current, parts[1:])
}

func findPossiblyCorrectCalibrations(calibrations []Calibration) []int {
	correct_calibrations := make([]int, 0)

	for _, calibration := range calibrations {
		if canListResultInTotal(calibration.Total, Reverse(calibration.Parts)) {
			correct_calibrations = append(correct_calibrations, calibration.Total)
		}
	}

	return correct_calibrations
} */

func ReadFile(path string) string {
	b, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(b)
}
func SumList(s []int) uint64 {
	var sum uint64 = 0

	for _, e := range s {
		sum += uint64(e)
	}

	return sum
}

func canListResultInTotal(expected int, parts []int) bool {
	if expected < 0 {
		// We cant subtract or divide a negative number by a positive one and result in a positive number - we can short circuit our search here
		return false
	}

	if len(parts) == 0 {
		return expected == 0
	}

	current := parts[0]
	if next := expected / current; expected == next*current && canListResultInTotal(next, parts[1:]) {
		return true
	}

	if next, found := strings.CutSuffix(strconv.Itoa(expected), strconv.Itoa(current)); found && len(next) != 0 {
		n, err := strconv.Atoi(next)
		if err != nil {
			panic(err)
		}

		if canListResultInTotal(n, parts[1:]) {
			return true
		}
	}

	return canListResultInTotal(expected-current, parts[1:])
}

func findPossiblyCorrectCalibrations(calibrations []Calibration) []int {
	correct_calibrations := make([]int, 0)

	for _, calibration := range calibrations {
		if canListResultInTotal(calibration.Total, Reverse(calibration.Parts)) {
			correct_calibrations = append(correct_calibrations, calibration.Total)
		}
	}

	return correct_calibrations
}

func main() {
	file := ReadFile("day7/input.txt")
	calibrations := ParseInput(file)

	correct := findPossiblyCorrectCalibrations(calibrations)
	answer := SumList(correct)

	fmt.Println(answer)

	fmt.Println(answer)
}
