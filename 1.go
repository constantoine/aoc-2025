package main

import (
	"strconv"
	"strings"
)

type DayOne struct{}

func (d DayOne) SolveOne(input string) string {
	var pointsToZero int

	pos := 50

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		shift, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		// remove full turns
		shift %= 100

		switch line[0] {
		case 'L':
			pos -= shift
		case 'R':
			pos += shift
		default:
			panic("invalid operand: " + string(line[0]))
		}

		pos = pos % 100
		if pos < 0 {
			pos = 100 + pos
		}

		if pos == 0 {
			pointsToZero++
		}
	}

	return strconv.Itoa(pointsToZero)
}

func (d DayOne) SolveTwo(input string) string {
	var clicksOnZero int

	pos := 50

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		shift, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		// Count and remove full turns
		clicksOnZero += shift / 100
		shift %= 100

		switch line[0] {
		case 'L':
			if pos != 0 && pos-shift < 0 {
				clicksOnZero++
			}

			pos -= shift
		case 'R':
			pos += shift

			if pos > 100 {
				clicksOnZero++
			}
		default:
			panic("invalid operand: " + string(line[0]))
		}

		pos = pos % 100

		if pos < 0 {
			pos = 100 + pos
		}

		if pos == 0 {
			clicksOnZero++
		}
	}

	return strconv.Itoa(clicksOnZero)
}
