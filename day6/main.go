package main

import (
	"aoc24/util"
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

func readAs2DArray(fileName string) ([][]string, int, int, error) {
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
			lineSlice = append(lineSlice, string(char)) // Append valid chars to the line slice
		}
		matrix = append(matrix, lineSlice) // Append the line slice to the main slice
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err) // Return any scanner error
	}
	x := 0
	y := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "^" {
				x = j
				y = i
			}
		}
	}

	return matrix, x, y, nil
}

func findNextObstacle(startx int, starty int, endx int, endy int, matrix [][]string) (int, int, [][]string) {
	obsX := endx
	obsY := endy
	if startx != endx {
		if endx < startx {
			for i := startx; i >= endx; i-- {
				if matrix[starty][i] == "#" {
					obsX = i
					obsY = starty
					break
				}
				matrix[starty][i] = "X"
			}
		} else {
			for i := startx; i <= endx; i++ {
				if matrix[starty][i] == "#" {
					obsX = i
					obsY = starty
					break
				}
				matrix[starty][i] = "X"
			}
		}
	} else if starty != endy {
		if endy < starty {
			for i := starty; i >= endy; i-- {
				if matrix[i][startx] == "#" {
					obsX = startx
					obsY = i
					break
				}
				matrix[i][startx] = "X"
			}
		} else {
			for i := starty; i <= endy; i++ {
				if matrix[i][startx] == "#" {
					obsX = startx
					obsY = i
					break
				}
				matrix[i][startx] = "X"
			}
		}
	}
	return obsX, obsY, matrix
}

func moveGuard(direction string, startx int, starty int, matrix [][]string) [][]string {
	if startx < 0 || starty < 0 || startx >= len(matrix[0]) || starty >= len(matrix) {
		return matrix
	}
	switch direction {
	case "n":
		obsX, obsY, matrixNew := findNextObstacle(startx, starty, startx, 0, matrix)
		if obsY == 0 {
			matrix[starty][startx] = "X"
			return matrix
		}
		moveGuard("e", obsX, obsY+1, matrixNew)
	case "s":
		obsX, obsY, matrixNew := findNextObstacle(startx, starty, startx, len(matrix)-1, matrix)
		if obsY == len(matrix)-1 {
			matrix[starty][startx] = "X"
			return matrix
		}
		moveGuard("w", obsX, obsY-1, matrixNew)
	case "e":
		obsX, obsY, matrixNew := findNextObstacle(startx, starty, len(matrix[0])-1, starty, matrix)
		if obsX == len(matrix[0])-1 {
			matrix[starty][startx] = "X"
			return matrix
		}
		moveGuard("e", obsX-1, obsY, matrixNew)
	case "w":
		obsX, obsY, matrixNew := findNextObstacle(startx, starty, 0, starty, matrix)
		if obsX == 0 {
			matrix[starty][startx] = "X"
			return matrix
		}
		moveGuard("e", obsX+1, obsY, matrixNew)

	}
	return matrix
}

/* func findStartingPostion(lines []string) (util.Vector, util.Vector) {
	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				return util.Vector{X: x, Y: y}, util.Vector{X: 0, Y: -1}
			}
		}
	}

	panic("Could not find \"^\" (starting position) in map.")
} */

/* func isInside(lines []string, position util.Vector) bool {
	if position.Y < 0 || position.Y > len(lines)-1 {
		return false
	}

	if position.X < 0 || position.X > len(lines[position.Y])-1 {
		return false
	}

	return true
} */

/* func nextPosition(lines []string, position util.Vector, direction util.Vector, directions_checked int) (util.Vector, util.Vector) {
	next := position.Add(direction)

	if isInside(lines, next) && rune(lines[next.Y][next.X]) == '#' {
		if directions_checked == 3 {
			panic(fmt.Sprintln("I'm trapped!", position))
		}

		return nextPosition(lines, position, direction.RotateOrigin90().Opposite(), directions_checked+1)
	}

	return next, direction
} */

/* func walkUntilLeaves(lines []string, position util.Vector, direction util.Vector) util.Set[util.Vector] {
	uniqueLocations := util.SetOf[util.Vector]()

	for isInside(lines, position) {
		uniqueLocations.Add(position)

		position, direction = nextPosition(lines, position, direction, 0)
	}

	return uniqueLocations
} */

type moment struct {
	position  util.Vector
	direction util.Vector
}

func findStartingPostion(lines []string) (util.Vector, util.Vector) {
	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				return util.Vector{X: x, Y: y}, util.Vector{X: 0, Y: -1}
			}
		}
	}

	panic("Could not find \"^\" (starting position) in map.")
}

func isInside(lines []string, position util.Vector) bool {
	if position.Y < 0 || position.Y > len(lines)-1 {
		return false
	}

	if position.X < 0 || position.X > len(lines[position.Y])-1 {
		return false
	}

	return true
}

func nextPosition(lines []string, position util.Vector, direction util.Vector) (util.Vector, util.Vector) {
	next := position.Add(direction)

	if isInside(lines, next) && rune(lines[next.Y][next.X]) == '#' {
		return position, direction.RotateOrigin90().Opposite()
	}

	return next, direction
}

func nextPositionWithObstacle(lines []string, position util.Vector, direction util.Vector, obstacle util.Vector) (util.Vector, util.Vector) {
	next := position.Add(direction)

	if isInside(lines, next) && (next == obstacle || rune(lines[next.Y][next.X]) == '#') {
		return position, direction.RotateOrigin90().Opposite()
	}

	return next, direction
}

func isLoop(lines []string, path util.Set[moment], position util.Vector, direction util.Vector, obstacle util.Vector) bool {
	seen := maps.Clone(path)

	for isInside(lines, position) {
		current := moment{position, direction}
		if seen.Contains(current) {
			return true
		}

		seen.Add(current)
		position, direction = nextPositionWithObstacle(lines, position, direction, obstacle)
	}

	return false
}

func walkUntilLeaves(lines []string, position util.Vector, direction util.Vector) util.Set[util.Vector] {
	path := util.SetOf[moment]()
	stepped_on := util.SetOf[util.Vector]()
	loopObstacles := util.SetOf[util.Vector]()

	for isInside(lines, position) {
		path.Add(moment{position, direction})
		stepped_on.Add(position)

		in_front := position.Add(direction)
		if !stepped_on.Contains(in_front) && isLoop(lines, path, position, direction.RotateOrigin90().Opposite(), in_front) {
			loopObstacles.Add(in_front)
		}

		position, direction = nextPosition(lines, position, direction)
	}

	return loopObstacles
}

func writeOutput(lines []string, obstacles util.Set[util.Vector]) {
	for y, line := range lines {
		for x, c := range line {
			if obstacles.Contains(util.Vector{X: x, Y: y}) {
				fmt.Print("O")
			} else {
				fmt.Print(string(c))
			}
		}

		fmt.Print("\n")
	}
}

/* func main() {
	matrix, startX, startY, _ := readAs2DArray("day6/test1.txt")
	matrixDone := moveGuard("n", startX, startY, matrix)
	distinctCount := 0
	for i := 0; i < len(matrixDone); i++ {
		for j := 0; j < len(matrixDone[0]); j++ {
			if matrixDone[i][j] == "X" {
				distinctCount++
			}
		}
	}
	file := util.ReadFile("day6/input.txt")
	lines := strings.Split(file, "\n")

	position, direction := findStartingPostion(lines)

	start := time.Now()
	obstacles := walkUntilLeaves(lines, position, direction)

	fmt.Println(len(obstacles))
	elapsed := time.Since(start)
	fmt.Println("Took:", elapsed)
} */

const (
	guardCh    = '^'
	obstacleCh = '#'
	freeCh     = '.'
)

var dirs = []fld.Pos{fld.Up, fld.Right, fld.Down, fld.Left}

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)
	guardPos := field.FindFirst(guardCh)
	dirIdx := 0
	visited := containers.NewSet(guardPos)
	for {
		npos := guardPos.Add(dirs[dirIdx])
		if !field.Inside(npos) {
			break
		}
		if field.Get(npos) == obstacleCh {
			dirIdx = (dirIdx + 1) % len(dirs) // Turn right.
		} else {
			guardPos = npos
			visited.Add(guardPos)
		}
	}
	return len(visited)
}

func SolvePart2(lines []string) any {
	// Brute force. Put obstacles on the path and check if the guard will loop.
	// Optimizations:
	// 1. Put obstacles only on the path from part1, not on the whole field (speedup x4)
	// 2. Remember visited states only on turns, not on every step (speedup x3)
	// 3. Start loop detection from the new obstacle, not from the initial guard position (speedup x3)
	// 4. Jump to the next obstable using precomputed obstacles positions (speedup x3)
	// 5. Track only visited states on turns from UP directions (speedup x2)
	// Total: 0.02s for the input (from 5s without optimizations).
	field := fld.NewByteField(lines)
	guardPos := field.FindFirst(guardCh)

	scanRowObstacles := func(row int) []int {
		var cols []int
		for col := range field.Cols() {
			if field.Get(fld.NewPos(row, col)) == obstacleCh {
				cols = append(cols, col)
			}
		}
		return cols
	}

	scanColObstacles := func(col int) []int {
		var rows []int
		for row := range field.Rows() {
			if field.Get(fld.NewPos(row, col)) == obstacleCh {
				rows = append(rows, row)
			}
		}
		return rows
	}

	rowObstacles := make([][]int, field.Rows())
	for row := range field.Rows() {
		rowObstacles[row] = scanRowObstacles(row)
	}

	colObstacles := make([][]int, field.Cols())
	for col := range field.Cols() {
		colObstacles[col] = scanColObstacles(col)
	}

	ans := 0
	dirIdx := 0
	visited := containers.NewSet(guardPos)
	for {
		npos := guardPos.Add(dirs[dirIdx])
		if !field.Inside(npos) {
			break
		}
		if field.Get(npos) == obstacleCh {
			dirIdx = (dirIdx + 1) % len(dirs) // Turn right.
		} else {
			if !visited.Has(npos) {
				field.Set(npos, obstacleCh)
				rowObstacles[npos.Row] = scanRowObstacles(npos.Row)
				colObstacles[npos.Col] = scanColObstacles(npos.Col)
				if isLooped(rowObstacles, colObstacles, guardPos, dirIdx) {
					ans++
				}
				field.Set(npos, freeCh)
				rowObstacles[npos.Row] = scanRowObstacles(npos.Row)
				colObstacles[npos.Col] = scanColObstacles(npos.Col)
				visited.Add(npos)
			}
			guardPos = npos
		}
	}
	return ans
}

func isLooped(rowObstacles, colObstacles [][]int, guardPos fld.Pos, dirIdx int) bool {
	type State struct {
		pos    fld.Pos
		dirIdx int
	}
	state := State{pos: guardPos, dirIdx: dirIdx}
	// Only track visited states on turns from UP directions, to speed up overall execution.
	visitedUp := containers.NewSet[fld.Pos]()
	for state.dirIdx != 0 || !visitedUp.Has(state.pos) {
		switch state.dirIdx {
		case 0: // UP
			visitedUp.Add(state.pos)
			// Find obstable in this col above the guard (row value less than guard's row).
			state.pos.Row = lowerBound(colObstacles[state.pos.Col], state.pos.Row, -100) + 1
		case 1: // RIGHT
			// Find obstable in this row to the right of the guard (col value greater than guard's col).
			state.pos.Col = upperBound(rowObstacles[state.pos.Row], state.pos.Col, -100) - 1
		case 2: // DOWN
			// Find obstable in this col below the guard (row value greater than guard's row).
			state.pos.Row = upperBound(colObstacles[state.pos.Col], state.pos.Row, -100) - 1
		case 3: // LEFT
			// Find obstable in this row to the left of the guard (col value less than guard's col).
			state.pos.Col = lowerBound(rowObstacles[state.pos.Row], state.pos.Col, -100) + 1
		default:
			panic("invalid direction")
		}
		if state.pos.Row < 0 || state.pos.Col < 0 {
			return false
		}
		// Turn right
		state.dirIdx = (state.dirIdx + 1) % len(dirs)
	}
	return true
}

// lowerBound returns the maximum value in `arr` which is less than `value`.
// `arr` is not empty and sorted in ascending order.
func lowerBound(arr []int, value int, defaultValue int) int {
	// Binary search doesn't make sense here because the array is small.
	ans := defaultValue
	for _, x := range arr {
		if x >= value {
			break
		}
		ans = x
	}
	return ans
}

// upperBound returns the minimun value in `arr` which is greater than `value`.
// `arr` is not empty and sorted in ascending order.
func upperBound(arr []int, value int, defaultValue int) int {
	// Binary search doesn't make sense here because the array is small.
	for _, x := range arr {
		if x > value {
			return x
		}
	}
	return defaultValue
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
