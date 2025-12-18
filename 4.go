package main

import (
	"image"
	"strconv"
	"strings"
)

type DayFour struct{}

func (d DayFour) SolveOne(input string) string {
	grid := make(map[image.Point]struct{})

	for x, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		for y, char := range line {
			if char == '@' {
				grid[image.Point{X: x, Y: y}] = struct{}{}
			}
		}
	}

	var total int

	for pos := range grid {
		neighbours := 0

		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x == 0 && y == 0 {
					continue
				}

				target := pos.Add(image.Point{X: x, Y: y})

				if _, ok := grid[target]; ok {
					neighbours++
				}
			}
		}

		if neighbours < 4 {
			total++
		}
	}

	return strconv.Itoa(total)
}

func (d DayFour) SolveTwo(input string) string {
	grid := make(map[image.Point]struct{})

	for x, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		for y, char := range line {
			if char == '@' {
				grid[image.Point{X: x, Y: y}] = struct{}{}
			}
		}
	}

	var total int
	var previousTotal int

	for {
		for pos := range grid {
			neighbours := 0

			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}

					target := pos.Add(image.Point{X: x, Y: y})

					if _, ok := grid[target]; ok {
						neighbours++
					}
				}
			}

			if neighbours < 4 {
				delete(grid, pos)
				total++
			}
		}

		if total == previousTotal {
			break
		}

		previousTotal = total
	}

	return strconv.Itoa(total)
}
