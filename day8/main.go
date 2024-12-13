package main

import (
	"bufio"
	"log"
	"os"
	"slices"
)

type location struct {
	x int
	y int
}

func readGridAsMap(fileName string) map[int][]string {
	grid := make(map[int][]string)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := make([]string, 0) // Create a new slice for the current line
		for _, char := range line {
			lineSlice = append(lineSlice, string(char)) // Append valid chars to the line slice
		}
		grid[i] = lineSlice
		i++
	}
	return grid
}

func addHashes(location1 location, location2 location, locations []location, curGrid map[int][]string) []location {
	distx := location1.x - location2.x
	disty := location1.y - location2.y

	newLoc1 := location{location2.x - distx, location2.y - disty}
	newLoc2 := location{location1.x + distx, location1.y + disty}

	/* fmt.Print(location1)
	fmt.Print(location2)
	fmt.Print(newLoc1)
	fmt.Print(newLoc2) */

	if !(newLoc1.x < 0 || newLoc1.y < 0 || newLoc1.x >= len(curGrid[0]) || newLoc1.y >= len(curGrid)) {
		if !slices.Contains(locations, newLoc1) {
			locations = append(locations, newLoc1)
		}
	}
	if !(newLoc2.x < 0 || newLoc2.y < 0 || newLoc2.x >= len(curGrid[0]) || newLoc2.y >= len(curGrid)) {
		if !slices.Contains(locations, newLoc2) {
			locations = append(locations, newLoc2)
		}
	}
	return locations
}

func findAllMatchingLocationsBelow(grid map[int][]string, startloc location, matchedChar string) []location {
	matchedlocations := make([]location, 0)
	for i := 0; i < len(grid); i++ {
		ycord := i
		line := grid[ycord]
		if ycord < startloc.y {
			continue
		}
		for xcord, char := range line {
			if char == matchedChar {
				if xcord == startloc.x && ycord == startloc.y {
					continue
				}
				matchedlocations = append(matchedlocations, location{xcord, ycord})
			}
		}
	}
	return matchedlocations
}

func part1(grid map[int][]string) int {
	locations := make([]location, 0)
	for i := 0; i < len(grid); i++ {
		ycord := i
		line := grid[ycord]
		for xcord, char := range line {
			if char != "." {
				/* fmt.Print(xcord)
				fmt.Print("\t")
				fmt.Print(ycord)
				fmt.Print("\n") */
				matchedlocations := findAllMatchingLocationsBelow(grid, location{xcord, ycord}, char)
				for _, loc := range matchedlocations {
					locations = addHashes(location{xcord, ycord}, loc, locations, grid)
				}
			}
		}

	}

	return len(locations)
}

func addHashes2(location1 location, location2 location, locations []location, curGrid map[int][]string) []location {
	distx := location1.x - location2.x
	disty := location1.y - location2.y
	distxC := location1.x - location2.x
	distyC := location1.y - location2.y
	antennaCount := 0
	for {
		newLoc1 := location{location2.x - distxC, location2.y - distyC}
		if newLoc1.x < 0 || newLoc1.y < 0 || newLoc1.x >= len(curGrid[0]) || newLoc1.y >= len(curGrid) {
			break
		}
		if !slices.Contains(locations, newLoc1) {
			locations = append(locations, newLoc1)
			antennaCount++
		}
		distxC += distx
		distyC += disty
	}

	distxC = location1.x - location2.x
	distyC = location1.y - location2.y
	for {
		newLoc2 := location{location1.x + distxC, location1.y + distyC}
		if newLoc2.x < 0 || newLoc2.y < 0 || newLoc2.x >= len(curGrid[0]) || newLoc2.y >= len(curGrid) {
			break
		}
		if !slices.Contains(locations, newLoc2) {
			locations = append(locations, newLoc2)
			antennaCount++
		}
		distxC += distx
		distyC += disty
	}

	if !slices.Contains(locations, location1) {
		locations = append(locations, location1)
	}
	if !slices.Contains(locations, location2) {
		locations = append(locations, location2)
	}

	return locations
}

func part2(grid map[int][]string) int {
	locations := make([]location, 0)
	for i := 0; i < len(grid); i++ {
		ycord := i
		line := grid[ycord]
		for xcord, char := range line {
			if char != "." {
				matchedlocations := findAllMatchingLocationsBelow(grid, location{xcord, ycord}, char)
				for _, loc := range matchedlocations {
					locations = addHashes2(location{xcord, ycord}, loc, locations, grid)
				}
			}
		}

	}
	/* fmt.Print(locations)

	for i := 0; i < len(grid); i++ {
		ycord := i
		line := grid[ycord]
		for xcord, _ := range line {
			if slices.Contains(locations, location{xcord, ycord}) {
				grid[ycord][xcord] = "#"
			}
		}

	}

	fmt.Print(grid) */

	return len(locations)
}

func main() {
	grid := readGridAsMap("day8/input.txt")
	println(part1(grid))

	println(part2(grid))
}
