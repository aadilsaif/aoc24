package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func inputConstructor(fileName string) []int {
	var blockedLine []int
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		id := 0
		for index, number := range line {
			num, _ := strconv.Atoi(string(number))
			if index%2 == 0 {
				for i := 0; i < num; i++ {
					blockedLine = append(blockedLine, id)
				}
				id++
			} else {
				for i := 0; i < num; i++ {
					blockedLine = append(blockedLine, -1)
				}
			}
		}
	}
	return blockedLine
}

func reOrder(blockedLine []int) []int {
	pointerB := len(blockedLine) - 1
	orderedLine := make([]int, len(blockedLine))
	_ = copy(orderedLine, blockedLine)
	for pointerA := 0; pointerA < len(blockedLine); pointerA++ {
		if pointerA >= pointerB {
			break
		}
		if orderedLine[pointerA] != -1 {
			continue
		}
		for {
			if pointerB <= pointerA {
				break
			}
			if orderedLine[pointerB] != -1 {
				orderedLine[pointerA] = orderedLine[pointerB]
				orderedLine[pointerB] = -1
				pointerB--
				break
			}
			pointerB--
		}
	}
	return orderedLine
}

func checkSum(orderedLine []int) int {
	sum := 0
	for index, value := range orderedLine {
		if value == -1 {
			continue
		}
		sum += (index * value)
	}
	return sum
}

func findLastChunk(orderedLine []int) (int, int) {
	start := -1
	end := -1
	for i := len(orderedLine) - 1; i > 0; i-- {
		if end == -1 && orderedLine[i] != -1 {
			end = i
			continue
		}
		if end == -1 {
			continue
		}
		if orderedLine[i] != orderedLine[end] {
			start = i + 1
			break
		}
	}
	return start, end

}

func reOrder2(blockedLine []int) []int {
	orderedLine := make([]int, len(blockedLine))
	_ = copy(orderedLine, blockedLine)
	lastChecked := len(orderedLine)
	for {
		firstSpace := slices.Index(orderedLine, -1)
		if lastChecked <= firstSpace {
			break
		}
		start, end := findLastChunk(orderedLine[0:lastChecked])
		if start == -1 || end == -1 || start < firstSpace {
			break
		}
		breaksize := end + 1 - start
		spaceStart := firstSpace
		for i := firstSpace; i < start; i++ {
			if orderedLine[i] != -1 {
				spaceStart = i + 1
				continue
			}
			if (i - spaceStart + 1) == breaksize {
				for j := 0; j < breaksize; j++ {
					orderedLine[spaceStart+j] = orderedLine[start+j]
					orderedLine[start+j] = -1
				}
				lastChecked = start
				break
			}
		}
		lastChecked = start
	}

	return orderedLine
}

func main() {
	line := inputConstructor("day9/input.txt")
	line = reOrder2(line)
	fmt.Println(checkSum(line))
}
