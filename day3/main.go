package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readOneLine(fileName string) (string, error) {
	var inputLine strings.Builder
	file, err := os.Open(fileName)
	if err != nil {
		return inputLine.String(), err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine.WriteString(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return inputLine.String(), err
	}
	return inputLine.String(), nil
}

func collectRegEx(regex string, text string) ([]string, error) {
	var collection []string
	re, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return collection, err
	}
	collection = re.FindAllString(text, -1)
	return collection, nil
}

func multiplier(matches []string) (int, error) {
	total := 0
	if len(matches) < 1 {
		return -1, fmt.Errorf("slice must contain at least one element")
	}
	pattern := `\d{1,3}`

	// Compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		return total, err
	}
	for _, match := range matches {
		integers := re.FindAllString(match, -1)
		if len(integers) != 2 {
			return -1, fmt.Errorf("matches are incorrect")
		}
		x, err1 := strconv.Atoi(integers[0])
		y, err2 := strconv.Atoi(integers[1])
		if err1 != nil {
			return -1, err1
		}
		if err2 != nil {
			return -1, err2
		}
		total += (x * y)
	}
	return total, nil
}

func part2(text string) ([]string, error) {
	var matches []string
	instructions := true
	for i := 0; i < len(text)-5; i++ {
		switch text[i : i+4] {
		case "don'":
			if (len(text) - i) > 6 {
				if text[i:i+7] == "don't()" {
					instructions = false
					i += 6
				}
			}
		case "do()":
			if (len(text) - i) > 3 {
				if text[i:i+4] == "do()" {
					instructions = true
					i += 3
				}
			}
		case "mul(":
			if instructions == true {
				endcheck := len(text)
				if (len(text) - i) > 11 {
					firstBracket := strings.IndexRune(text[i:], ')')
					if firstBracket != -1 {
						endcheck = i + firstBracket + 1
					}
				}
				re, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
				if err != nil {
					fmt.Println("Error compiling regex:", err)
					return matches, err
				}
				matchedString := re.FindString(text[i:endcheck])
				if matchedString != "" {
					matches = append(matches, matchedString)
					i = endcheck - 1
				}

			}
		}
	}

	return matches, nil

}

func main() {
	text, err := readOneLine("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	matches, err := collectRegEx(pattern, text)
	fmt.Print(matches)
	total, err := multiplier(matches)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nTotal Value is %d\n", total)
	matches2, err := part2(text)
	fmt.Print(matches2)
	total2, err := multiplier(matches2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nTotal Value is %d\n", total2)
}
