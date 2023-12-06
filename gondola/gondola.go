package main

import (
	"bufio"
	"fmt"
	"os"
)

type Part struct {
	number int
	row int
	start int
	end int
}

type SymbolLocation struct {
	row int
	col int
}

func abs(x int) int {
    if x < 0 {
            return -x
    }
    return x
}

func IsPartNumberRelevant(part Part, symbols []SymbolLocation) bool {
	for _, symbol := range symbols {
		withinWidth := ((symbol.col + 1) >= part.start) && ((symbol.col - 1) <= part.end)
		adjacentRows := (abs(symbol.row - part.row) <= 1)
		if withinWidth && adjacentRows {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------

func main() {
	readFile, _ := os.Open("engine.txt")

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var partNumbersSum = 0
	for fileScanner.Scan() {
		
	}

	fmt.Printf("Sum of all the active part numbers: %d\n", partNumbersSum)
}
