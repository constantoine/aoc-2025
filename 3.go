package main

import (
	"math"
	"strconv"
	"strings"
)

type DayThree struct{}

func (d DayThree) SolveOne(input string) string {
	return d.solve(input, 2)
}

func (d DayThree) SolveTwo(input string) string {
	return d.solve(input, 12)
}

func (d DayThree) solve(input string, batteryCount int) string {
	var jolts uint64

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		for i := batteryCount - 1; i >= 0; i-- {
			highest, pos := uint64(0), 0
			highestIndex := len(line) - 1 - i
			for ii := 0; ii <= highestIndex; ii++ {
				digit := uint64(line[ii] - '0')
				if digit > highest {
					highest = digit
					pos = ii
				}
			}

			line = line[pos+1:]
			jolts += highest * uint64(math.Pow10(i))
		}
	}

	return strconv.FormatUint(jolts, 10)
}
