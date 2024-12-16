package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Define a Node structure
type Node struct {
	data int
	next *Node
}

// Define a LinkedList structure
type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBack(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func initStones(fileName string) LinkedList {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var list LinkedList
	if scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		for _, num := range nums {
			convertedNum, _ := strconv.Atoi(num)
			list.insertAtBack(convertedNum)
		}
	}
	return list
}

func checkIfEvenDigits(number int) bool {
	numDigits := len(strconv.Itoa(number))
	return numDigits%2 == 0
}

func (l *LinkedList) blink() {
	current := l.head
	for current != nil {
		if current.data == 0 {
			current.data = 1
			current = current.next
		} else if checkIfEvenDigits(current.data) {
			nextNode := current.next
			numString := strconv.Itoa(current.data)
			firstHalf := numString[:len(numString)/2]
			secondHalf := numString[len(numString)/2:]
			firstNodeData, _ := strconv.Atoi(firstHalf)
			secondNodeData, _ := strconv.Atoi(secondHalf)
			current.data = firstNodeData
			secondNode := &Node{data: secondNodeData, next: nextNode}
			current.next = secondNode
			current = nextNode
		} else {
			current.data = current.data * 2024
			current = current.next
		}
	}
}

func part1(fileName string) int {
	totalStones := 0
	list := initStones(fileName)
	for i := 0; i < 25; i++ {
		list.blink()
	}
	current := list.head
	for current != nil {
		totalStones++
		current = current.next
	}

	return totalStones

}

type config struct {
	v int
	n int
}

func getStonesAfterBlink(stone int, iterationNum int, cache map[config]int) int {
	if iterationNum == 0 {
		return 1
	}
	if r, ok := cache[config{stone, iterationNum}]; ok {
		return r
	}

	if stone == 0 {
		res := getStonesAfterBlink(1, iterationNum-1, cache)
		cache[config{stone, iterationNum}] = res
		return res
	}

	if s := strconv.Itoa(stone); len(s)%2 == 0 {
		a, _ := strconv.Atoi(s[:len(s)/2])
		b, _ := strconv.Atoi(s[len(s)/2:])
		res := getStonesAfterBlink(a, iterationNum-1, cache) + getStonesAfterBlink(b, iterationNum-1, cache)
		cache[config{stone, iterationNum}] = res
		return res
	}

	res := getStonesAfterBlink(stone*2024, iterationNum-1, cache)
	cache[config{stone, iterationNum}] = res
	return res
}

func part2(fileName string) int {
	list := initStones(fileName)
	current := list.head
	totalStones := 0
	var cache = make(map[config]int)
	for current != nil {
		totalStones += getStonesAfterBlink(current.data, 75, cache)
		current = current.next
	}
	return totalStones
}

func main() {
	fmt.Println(part1("day11/input.txt"))
	fmt.Println(part2("day11/input.txt"))

}
