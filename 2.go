package main

import (
	"strconv"
	"strings"
)

type DayTwo struct{}

func (d DayTwo) SolveOne(input string) string {
	return d.solve(input, d.isValidPartOne)
}

func (d DayTwo) SolveTwo(input string) string {
	return d.solve(input, d.isValidPartTwo)
}

func (d DayTwo) solve(input string, isValid func(input uint64) bool) string {
	var invalidIDs uint64

	for _, line := range strings.Split(input, ",") {
		if line == "" {
			continue
		}

		ids := strings.Split(line, "-")

		first, err := strconv.ParseUint(ids[0], 10, 64)
		if err != nil {
			panic("invalid first id : " + err.Error())
		}

		second, err := strconv.ParseUint(ids[1], 10, 64)
		if err != nil {
			panic("invalid second id : " + err.Error())
		}

		for i := first; i <= second; i++ {
			if !isValid(i) {
				invalidIDs += i
			}
		}
	}

	return strconv.FormatUint(invalidIDs, 10)
}

func (d DayTwo) isValidPartOne(input uint64) bool {
	num := strconv.FormatUint(input, 10)

	if len(num)%2 == 1 {
		return true
	}

	if num[0:len(num)/2] == num[len(num)/2:] {
		return false
	}

	return true
}

func (d DayTwo) isValidPartTwo(input uint64) bool {
	num := strconv.FormatUint(input, 10)

	for l := 1; l < len(num); l++ {
		c := strings.Count(num, num[:l])

		if float64(c) == float64(len(num))/float64(l) {
			return false
		}
	}

	return true
}
