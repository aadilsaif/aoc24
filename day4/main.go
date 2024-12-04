package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readAs2DArray(fileName string) ([][]string, error) {
	var matrix [][]string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed at the end of the function

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := make([]string, 0) // Create a new slice for the current line
		for _, char := range line {
			lineSlice = append(lineSlice, string(char)) // Append valid numbers to the line slice
		}
		matrix = append(matrix, lineSlice) // Append the line slice to the main slice
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err) // Return any scanner error
	}

	return matrix, nil
}

func getValue(matrix [][]string, x int, y int) string {
	if x < 0 || y < 0 || y > len(matrix)-1 || x > (len(matrix[0]))-1 {
		return ""
	}
	return matrix[y][x]
}

func checkDirections(matrix [][]string, x int, y int) int {
	if getValue(matrix, x, y) == "" {
		return 0
	}
	score := 0
	if getValue(matrix, x, y-3) == "S" && getValue(matrix, x, y-2) == "A" && getValue(matrix, x, y-1) == "M" {
		score += 1
	}
	if getValue(matrix, x+3, y-3) == "S" && getValue(matrix, x+2, y-2) == "A" && getValue(matrix, x+1, y-1) == "M" {
		score += 1
	}
	if getValue(matrix, x+3, y) == "S" && getValue(matrix, x+2, y) == "A" && getValue(matrix, x+1, y) == "M" {
		score += 1
	}
	if getValue(matrix, x+3, y+3) == "S" && getValue(matrix, x+2, y+2) == "A" && getValue(matrix, x+1, y+1) == "M" {
		score += 1
	}
	if getValue(matrix, x, y+3) == "S" && getValue(matrix, x, y+2) == "A" && getValue(matrix, x, y+1) == "M" {
		score += 1
	}
	if getValue(matrix, x-3, y+3) == "S" && getValue(matrix, x-2, y+2) == "A" && getValue(matrix, x-1, y+1) == "M" {
		score += 1
	}
	if getValue(matrix, x-3, y) == "S" && getValue(matrix, x-2, y) == "A" && getValue(matrix, x-1, y) == "M" {
		score += 1
	}
	if getValue(matrix, x-3, y-3) == "S" && getValue(matrix, x-2, y-2) == "A" && getValue(matrix, x-1, y-1) == "M" {
		score += 1
	}
	return score
}

func part1(matrix [][]string) int {
	score := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if getValue(matrix, j, i) != "X" {
				continue
			}
			score += checkDirections(matrix, j, i)
		}
	}
	return score
}

func checkDirections2(matrix [][]string, x int, y int) int {
	if getValue(matrix, x, y) == "" {
		return 0
	}
	score := 0
	if getValue(matrix, x, y-1) == "M" && getValue(matrix, x, y+1) == "S" {
		if getValue(matrix, x+1, y) == "M" && getValue(matrix, x-1, y) == "S" {
			score = 1
		}
		if getValue(matrix, x+1, y) == "S" && getValue(matrix, x-1, y) == "M" {
			score = 1
		}
	}
	if getValue(matrix, x, y-1) == "S" && getValue(matrix, x, y+1) == "M" {
		if getValue(matrix, x+1, y) == "M" && getValue(matrix, x-1, y) == "S" {
			score = 1
		}
		if getValue(matrix, x+1, y) == "S" && getValue(matrix, x-1, y) == "M" {
			score = 1
		}
	}
	if getValue(matrix, x-1, y-1) == "M" && getValue(matrix, x+1, y+1) == "S" {
		if getValue(matrix, x+1, y-1) == "M" && getValue(matrix, x-1, y+1) == "S" {
			score = 1
		}
		if getValue(matrix, x+1, y-1) == "S" && getValue(matrix, x-1, y+1) == "M" {
			score = 1
		}
	}
	if getValue(matrix, x-1, y-1) == "S" && getValue(matrix, x+1, y+1) == "M" {
		if getValue(matrix, x+1, y-1) == "M" && getValue(matrix, x-1, y+1) == "S" {
			score = 1
		}
		if getValue(matrix, x+1, y-1) == "S" && getValue(matrix, x-1, y+1) == "M" {
			score = 1
		}
	}
	return score
}
func part2(matrix [][]string) int {
	score := 0
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			if getValue(matrix, j, i) != "A" {
				continue
			}
			score += checkDirections2(matrix, j, i)
		}
	}
	return score
}

func main() {
	matrix, err := readAs2DArray("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(part1(matrix))
	fmt.Print("\n")
	fmt.Print(part2(matrix))
}
