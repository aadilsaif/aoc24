package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func calcDistance(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {
	file, e := os.Open("day1/input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanWords)

	var input []string
	for s.Scan() {
		input = append(input, s.Text())
	}

	var listA, listB []int

	for i, e := range input {
		v, _ := strconv.Atoi(e)
		if i%2 == 0 {
			listA = append(listA, v)
			continue
		}
		listB = append(listB, v)
	}

	listA_sorted := make([]int, len(listA))
	listB_sorted := make([]int, len(listB))

	copy(listA_sorted, listA)
	copy(listB_sorted, listB)

	sort.Slice(listA_sorted, func(i, j int) bool {
		return listA_sorted[i] < listA_sorted[j]
	})

	sort.Slice(listB_sorted, func(i, j int) bool {
		return listB_sorted[i] < listB_sorted[j]
	})

	sum := 0
	for i := range listA_sorted {
		sum += calcDistance(listA_sorted[i], listB_sorted[i])
	}

	fmt.Println("Part 1 solution:", sum)

	//// Solution for Part 2

	n := 0
	similarity := 0
	for _, i := range listA {
		n = 0
		for _, j := range listB {
			if i == j {
				n++
			}
		}
		similarity += i * n
	}

	fmt.Println("Part 2 solution:", similarity)
}
