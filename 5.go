package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type DayFive struct{}

func (d DayFive) SolveOne(input string) string {
	ranges := make([]Range, 0)

	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			lines = lines[i+1:]
			break
		}

		ids := strings.Split(lines[i], "-")

		first, err := strconv.ParseInt(ids[0], 10, 64)
		if err != nil {
			panic("invalid first id : " + err.Error())
		}

		second, err := strconv.ParseInt(ids[1], 10, 64)
		if err != nil {
			panic("invalid last id : " + err.Error())
		}

		ranges = append(ranges, Range{first, second})
	}

	var totalFresh int

	for _, line := range lines {
		if line == "" {
			continue
		}

		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic("invalid id : " + err.Error())
		}

		for i := range ranges {
			if ranges[i].Contains(id) {
				totalFresh++
				break
			}
		}
	}

	return strconv.Itoa(totalFresh)
}

func (d DayFive) SolveTwo(input string) string {
	ranges := make([]Range, 0)

	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}

		ids := strings.Split(lines[i], "-")

		first, err := strconv.ParseInt(ids[0], 10, 64)
		if err != nil {
			panic("invalid first id : " + err.Error())
		}

		second, err := strconv.ParseInt(ids[1], 10, 64)
		if err != nil {
			panic("invalid last id : " + err.Error())
		}

		ranges = append(ranges, Range{first, second})
	}

	var total int64

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	for _, r := range ranges {
		fmt.Println(r)
	}

	nonIntersect := make([]Range, 0, len(ranges))

	for i := 0; i < len(ranges); i++ {
		if i == len(ranges)-1 {
			nonIntersect = append(nonIntersect, ranges[i])

			break
		}

		if ranges[i].Intersect(ranges[i+1]) {
			ranges[i+1] = Range{
				Min: min(ranges[i].Min, ranges[i+1].Min),
				Max: max(ranges[i].Max, ranges[i+1].Max),
			}
		} else {
			nonIntersect = append(nonIntersect, ranges[i])
		}
	}

	for _, r := range nonIntersect {
		total += r.Size()
	}

	return strconv.FormatInt(total, 10)
}
