package main_test

import (
	"testing"

	"github/constantoine/aoc-2025"
)

func TestThreePartOne(t *testing.T) {
	testPuzzle(t, 3, 1, func() main.PartialPuzzle {
		return main.DayThree{}
	})
}

func TestThreePartTwo(t *testing.T) {
	testPuzzle(t, 3, 2, func() main.PartialPuzzle {
		return main.DayThree{}
	})
}
