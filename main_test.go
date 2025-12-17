package main_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github/constantoine/aoc-2025"
)

func testPuzzle(t *testing.T, day, part int, solverFactory func() main.PartialPuzzle) {
	t.Helper()

	dirname := fmt.Sprintf("testdata/%d/%d", day, part)

	dir, err := os.ReadDir(dirname)
	if err != nil {
		t.Fatalf("failed to read dir %q: %s", dirname, err)
	}

	switch part {
	case 1:
	case 2:
		solver := solverFactory()

		if _, ok := solver.(main.CompletePuzzle); !ok {
			t.Fatalf("puzzle %d only implements part one", day)
		}
	default:
		t.Fatalf("invalid part: %d", part)
	}

	for i, entry := range dir {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			solver := solverFactory()

			fileName := fmt.Sprintf("%s/%s", dirname, entry.Name())

			input, err := os.ReadFile(fileName)
			if err != nil {
				t.Fatalf("failed to open solution file %q: %s", fileName, err)
			}

			switch part {
			case 1:
				testOutput := solver.SolveOne(string(input))
				if testOutput != entry.Name() {
					t.Logf("expected solution output %q, got %q", entry.Name(), testOutput)

					t.Fail()
				}
			case 2:
				testOutput := solver.(main.CompletePuzzle).SolveTwo(string(input))
				if testOutput != entry.Name() {
					t.Logf("expected solution output %q, got %q", entry.Name(), testOutput)

					t.Fail()
				}
			}
		})
	}
}
