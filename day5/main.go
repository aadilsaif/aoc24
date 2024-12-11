package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func readFile(fileName string) (map[int][]int, [][]int) {
	inputSlice := make([]string, 0)
	rules := make(map[int][]int)
	updates := make([][]int, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return rules, updates
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputSlice = append(inputSlice, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
		return rules, updates
	}
	linebreak := slices.Index(inputSlice, "")
	if linebreak == -1 {
		log.Fatal("Cant find line break")
	}
	ruleStrings := inputSlice[0:linebreak]
	updateStrings := inputSlice[linebreak+1:]

	for _, line := range ruleStrings {
		numbers := strings.Split(line, "|")
		if len(numbers) != 2 {
			log.Fatal("line is not a valid rule type")
		}
		prefix, _ := strconv.Atoi(numbers[0])
		suffix, _ := strconv.Atoi(numbers[1])
		_, ok := rules[prefix]
		if ok {
			rules[prefix] = append(rules[prefix], suffix)
		} else {
			rules[prefix] = []int{suffix}
		}
	}

	for _, line := range updateStrings {
		numbers := strings.Split(line, ",")
		if len(numbers) == 1 {
			log.Fatal("Update string is empty")
		}
		numSlice := make([]int, 0)
		for _, num := range numbers {
			val, _ := strconv.Atoi(num)
			numSlice = append(numSlice, val)
		}
		updates = append(updates, numSlice)
	}
	return rules, updates
}

func matchingNumbers(sliceA []int, sliceB []int) bool {
	out := []int{}
	bucket := map[int]bool{}
	for _, i := range sliceA {
		for _, j := range sliceB {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}
	if len(out) == 0 {
		return false
	}
	return true
}

func part1(rules map[int][]int, updates [][]int) {
	validCount := 0
	totalScore := 0
	for _, update := range updates {
		passing := true
		for val, number := range update {
			v, ok := rules[number]
			if ok {
				if matchingNumbers(v, update[0:val]) {
					passing = false
				}
			}
		}
		if passing {
			validCount++
			totalScore += update[len(update)/2]
		}
	}
	fmt.Printf("\nThere are %d valid updates", validCount)
	fmt.Printf("\nThe score is %d", totalScore)
}

func matchingNumbers2(sliceA []int, sliceB []int) (bool, int) {
	for val, i := range sliceA {
		for _, j := range sliceB {
			if i == j {
				return true, val
			}
		}
	}
	return false, -1
}

func validateFinding(rules map[int][]int, update []int) []int {
	/* 	if ok, _ := validateFinding(rules, update); ok {
		return true, update
	} */
	for val, number := range update {
		v, ok := rules[number]
		if ok {
			if match, index := matchingNumbers2(update[0:val], v); match {
				sliceEnd := make([]int, 0)
				if index == 0 {
					continue
				} else {
					sliceEnd = update[index-1:]
				}
				sliceStart := update[0:index]
				sliceEnd = slices.Delete(sliceEnd, val-index, val-index+1)
				slice := slices.Concat(sliceStart, []int{update[val]}, sliceEnd)
				validateFinding(rules, slice)
			}
		}
	}
	return update
}

func part2(rules map[int][]int, updates [][]int) {
	invalidCount := 0
	totalScore := 0
	for _, update := range updates {
		for val, number := range update {
			v, ok := rules[number]
			if ok {
				if matchingNumbers(v, update[0:val]) {
					fixedUpdate := validateFinding(rules, update)
					invalidCount++
					totalScore += update[len(fixedUpdate)/2]
					break
				}
			}
		}
	}
	fmt.Printf("\nThere are %d invalid updates", invalidCount)
	fmt.Printf("\nThe score is %d", totalScore)
}

/* func main() {
	rules, updates := readFile("day5/input.txt")
	part2(rules, updates)
} */

func main() {
	result1 := Task1("day5/input.txt")
	result2 := Task2("day5/input.txt")
	fmt.Printf("RESULT T1: %d, RESULT T2: %d\n", result1, result2)
}

func Task1(fileName string) int {
	rawLines := ReadFile(fileName)
	rules, updates := getPartLists(rawLines)
	filteredUpdates := filterValidLines(rules, updates)
	return calcResult1(filteredUpdates)
}

func Task2(fileName string) int {
	rawLines := ReadFile(fileName)
	rules, updates := getPartLists(rawLines)
	invalids := invalidUpdates(&rules, &updates)
	var ordered [][]int
	for _, invalidRule := range invalids {
		ordered = append(ordered, orderLine(&rules, invalidRule))
	}
	return calcResult1(ordered)
}

func getPartLists(input []string) ([][]int, [][]int) {
	var rules [][]int
	var updates [][]int
	r_r, _ := regexp.Compile(`[0-9]+\|[0-9]+`)
	r_u, _ := regexp.Compile(`[0-9]+(,[0-9]+)+`)
	for _, line := range input {
		if r_r.FindString(line) != "" {
			command := strings.Split(line, "|")
			c_1, _ := strconv.Atoi(command[0])
			c_2, _ := strconv.Atoi(command[1])
			rules = append(rules, []int{c_1, c_2})
		} else if r_u.FindString(line) != "" {
			elements := strings.Split(line, ",")
			var temp_update []int
			for _, num := range elements {
				num_i, _ := strconv.Atoi(num)
				temp_update = append(temp_update, num_i)
			}
			updates = append(updates, temp_update)
		}
	}
	return rules, updates
}

func filterValidLines(ruleset [][]int, updates [][]int) [][]int {
	var validLines [][]int
	for _, update := range updates {
		isInvalid := false
		var forbiddenNumbers []int
		for _, num := range update {
			if isInvalid {
				continue
			}
			if slices.Contains(forbiddenNumbers, num) {
				isInvalid = true
				continue
			} else {
				forbiddenNumbers = append(forbiddenNumbers, returnRules(ruleset, num)...)
			}
		}
		if !isInvalid {
			validLines = append(validLines, update)
		}
	}
	return validLines
}

func invalidUpdates(ruleset *[][]int, updates *[][]int) [][]int {
	validLines := filterValidLines(*ruleset, *updates)
	var filtered [][]int
	for _, update := range *updates {
		found := false
		for _, validLine := range validLines {
			if found {
				continue
			} else if equalSlices(update, validLine) {
				found = true
			}
		}
		if !found {
			filtered = append(filtered, update)
		}
	}
	return filtered
}

func returnRules(ruleset [][]int, num int) []int {
	var invalidNumbers []int
	for _, rule := range ruleset {
		if rule[1] == num {
			invalidNumbers = append(invalidNumbers, rule[0])
		}
	}
	return invalidNumbers
}

func orderLine(rules *[][]int, line []int) []int {
	var forbiddenNumbers []int
	for i, num := range line {
		if slices.Contains(forbiddenNumbers, num) {
			line[i-1], line[i] = line[i], line[i-1]
			return orderLine(rules, line)
		} else {
			forbiddenNumbers = append(forbiddenNumbers, returnRules(*rules, num)...)
		}
	}
	return line
}

func calcResult1(validPages [][]int) int {
	result := 0
	for _, element := range validPages {
		middleIndex := int(len(element) / 2)
		result += element[middleIndex]
	}
	return result
}

func equalSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func ReadFile(fileName string) []string {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	var returnArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		returnArray = append(returnArray, text)
	}
	return returnArray
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
