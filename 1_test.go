package main_test

import (
	"testing"

	"github/constantoine/aoc-2025"
)

func TestOnePartOne(t *testing.T) {
	testPuzzle(t, 1, 1, func() main.PartialPuzzle {
		return main.DayOne{}
	})
}

func TestOnePartTwo(t *testing.T) {
	testPuzzle(t, 1, 2, func() main.PartialPuzzle {
		return main.DayOne{}
	})
}
