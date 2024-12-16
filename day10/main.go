package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type location struct {
	x int
	y int
}

type Grid struct {
	data       map[int][]int
	trailheads []location
	endpoints  []location
}

func readFileAsGrid(fileName string) map[int][]int {
	grid := make(map[int][]int)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := make([]int, 0) // Create a new slice for the current line
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			lineSlice = append(lineSlice, num) // Append valid chars to the line slice
		}
		grid[i] = lineSlice
		i++
	}
	return grid
}

func createGrid(fileName string) *Grid {
	grid := readFileAsGrid(fileName)
	trailheadLocs := make([]location, 0)
	for i := 0; i < len(grid); i++ {
		ycord := i
		line := grid[ycord]
		for xcord, num := range line {
			if num == 0 {
				trailheadLocs = append(trailheadLocs, location{xcord, i})
			}
		}
	}
	return &Grid{
		data:       grid,
		trailheads: trailheadLocs,
		endpoints:  make([]location, 0),
	}
}

func (g *Grid) getLocationValue(loc location) int {
	return g.data[loc.y][loc.x]
}

func (g *Grid) getScore() int {
	score := 0
	for _, trailhead := range g.trailheads {
		found := g.walk(trailhead)
		fmt.Println(found)
		score += found
		slices.Delete(g.endpoints, 0, len(g.endpoints))
	}
	return score
}

func (g *Grid) getEndpoints() int {
	return len(g.endpoints)
}

func (g *Grid) walk(loc location) int {
	if !g.validSquare(loc) {
		return 0
	}
	height := g.getLocationValue(loc)
	if height == 9 {
		if !slices.Contains(g.endpoints, loc) {
			g.endpoints = append(g.endpoints, loc)
		}
		return 1
	}

	score := 0
	if loc.x > 0 && (g.getLocationValue(location{loc.x - 1, loc.y}) == (g.getLocationValue(loc) + 1)) {
		score += g.walk(location{loc.x - 1, loc.y})
	}
	if loc.y > 0 && (g.getLocationValue(location{loc.x, loc.y - 1}) == (g.getLocationValue(loc) + 1)) {
		score += g.walk(location{loc.x, loc.y - 1})
	}
	if loc.y < len(g.data)-1 && (g.getLocationValue(location{loc.x, loc.y + 1}) == (g.getLocationValue(loc) + 1)) {
		score += g.walk(location{loc.x, loc.y + 1})
	}
	if loc.x < len(g.data[0])-1 && (g.getLocationValue(location{loc.x + 1, loc.y}) == (g.getLocationValue(loc) + 1)) {
		score += g.walk(location{loc.x + 1, loc.y})
	}

	return score
}

func (g *Grid) validSquare(loc location) bool {
	if loc.x >= 0 && loc.y >= 0 && loc.x < len(g.data[0]) && loc.y < len(g.data) {
		return true
	}
	return false
}

func main() {
	grid := createGrid("day10/input.txt")
	fmt.Println(grid.trailheads)
	fmt.Println(grid.getScore())
	fmt.Println(grid.getEndpoints())
}
