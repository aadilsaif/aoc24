package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func copySliceWOindex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func MultipleLinesOfNumbers(fileName string) [][]int {
	var inputSlice [][]int // Slice of slices to hold numbers from each line
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed at the end of the function

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line) // Split the line into fields (numbers)
		lineSlice := make([]int, 0)     // Create a new slice for the current line
		for _, numStr := range numbers {
			thisNumber, err := strconv.Atoi(numStr)
			if err != nil {
				continue // Skip invalid numbers
			}
			lineSlice = append(lineSlice, thisNumber) // Append valid numbers to the line slice
		}
		inputSlice = append(inputSlice, lineSlice) // Append the line slice to the main slice
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err) // Return any scanner error
	}

	return inputSlice
}

func CheckTrend(nums []int) (string, error) {
	if len(nums) < 2 {
		return "", fmt.Errorf("slice must contain at least two elements")
	}
	trend := "neither" // Default trend
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			if trend == "neither" {
				trend = "increasing"
			} else if trend == "decreasing" {
				return "neither", nil
			}
		} else if nums[i] > nums[i+1] {
			if trend == "neither" {
				trend = "decreasing"
			} else if trend == "increasing" {
				return "neither", nil
			}
		} else {
			return "neither", nil // Equal elements
		}
	}
	return trend, nil
}

func checkDistance(nums []int) (bool, error) {
	if len(nums) < 2 {
		return false, fmt.Errorf("slice must contain at least two elements")
	}
	for i := 0; i < len(nums)-1; i++ {
		if distance := func(x, y int) int {
			if x < y {
				return y - x
			}
			return x - y
		}(nums[i], nums[i+1]); distance > 0 && distance < 4 {
			continue
		}
		return false, nil
	}
	return true, nil
}

func part1(numberSlices [][]int) int {
	score := 0
	for _, line := range numberSlices {
		if !checkPass(line) {
			continue
		}
		score++
	}
	return score
}

func CheckTrend2(nums []int) (string, int, error) {
	if len(nums) < 2 {
		return "", -1, fmt.Errorf("slice must contain at least two elements")
	}
	badElement := -1
	trend := "neither" // Default trend
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			if trend == "neither" {
				trend = "increasing"
			} else if trend == "decreasing" {
				if badElement != -1 {
					return "neither", badElement, nil
				}
				secondCheck, err := CheckTrend(copySliceWOindex(nums, i+1))
				if err != nil {
					log.Fatal(err)
				}
				if secondCheck != "decreasing" {
					return "neither", badElement, nil
				}
				badElement = i + 1
				i++
			}
		} else if nums[i] > nums[i+1] {
			if trend == "neither" {
				trend = "decreasing"
			} else if trend == "increasing" {
				if badElement != -1 {
					return "neither", badElement, nil
				}
				secondCheck, err := CheckTrend(copySliceWOindex(nums, i+1))
				if err != nil {
					log.Fatal(err)
				}
				if secondCheck != "increasing" {
					return "neither", badElement, nil
				}
				badElement = i + 1
				i++
			}
		} else {
			if badElement != -1 {
				return "neither", badElement, nil
			}
			secondCheck, err := CheckTrend(copySliceWOindex(nums, i+1))
			if err != nil {
				log.Fatal(err)
			}
			if secondCheck != trend {
				return "neither", badElement, nil
			}
			badElement = i + 1
		}
	}
	return trend, badElement, nil
}

func checkDistance2(nums []int) (bool, int, error) {
	if len(nums) < 2 {
		return false, -1, fmt.Errorf("slice must contain at least two elements")
	}
	badElement := -1
	for i := 0; i < len(nums)-1; i++ {
		if distance := func(x, y int) int {
			if x < y {
				return y - x
			}
			return x - y
		}(nums[i], nums[i+1]); distance > 0 && distance < 4 {
			continue
		}
		if badElement == -1 {
			res, err := checkDistance(copySliceWOindex(nums, i+1))
			if err != nil {
				log.Fatal(err)
			}
			if res == true {
				return true, i + 1, nil
			}
			return false, badElement, nil
		}
		return false, badElement, nil
	}
	return true, badElement, nil
}

func checkPass(slice []int) bool {
	trend, err := CheckTrend(slice)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if trend == "neither" {
		return false
	}
	difference, err := checkDistance(slice)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if difference == false {
		return false
	}
	return true
}

func part2(numberSlices [][]int) int {
	score := 0
	for _, line := range numberSlices {
		if !checkPass(line) {
			for index, _ := range line {
				if checkPass(copySliceWOindex(line, index)) {
					score++
					break
				}
			}
			continue
		}
		score++
	}
	return score
}

func main() {
	fmt.Print(part1(MultipleLinesOfNumbers("day2/input.txt")))
	fmt.Print("\n")
	fmt.Print(part2(MultipleLinesOfNumbers("day2/input.txt")))
}
