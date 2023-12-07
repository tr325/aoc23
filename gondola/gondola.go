package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"github.com/samber/lo"
)

type Part struct {
	number int
	row    int
	start  int
	end    int
}

type SymbolLocation struct {
	row int
	col int
}

type SymbolData struct {
	adjacentParts []Part
	symbol        string
	location      SymbolLocation
}

func parse(fileScanner *bufio.Scanner) ([]Part, []SymbolData) {
	numberFinder := regexp.MustCompile(`[0-9]*`)
	symbolFinder := regexp.MustCompile(`[^A-Za-z0-9\.]`)

	var row = 0
	parts := []Part{}
	symbolData := []SymbolData{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := numberFinder.FindAllString(line, -1)
		numberLocs := numberFinder.FindAllStringIndex(line, -1)
		symbols := symbolFinder.FindAllString(line, -1)
		symbolLocations := symbolFinder.FindAllStringIndex(line, -1)
		for i, num := range numbers {
			number, _ := strconv.Atoi(num)
			parts = append(parts, Part{number, row, numberLocs[i][0], numberLocs[i][1] - 1})
		}
		for j, loc := range symbolLocations {
			location := SymbolLocation{row, loc[0]}
			data := SymbolData{[]Part{}, symbols[j], location}
			symbolData = append(symbolData, data)
		}
		row = row + 1
	}
	return parts, symbolData
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
		adjacentRows := (abs(symbol.row-part.row) <= 1)
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

	parts, symbols := parse(fileScanner)
	symbolLocations := lo.Map(symbols, func(s SymbolData, _ int) SymbolLocation {
	    return s.location
	})

	var partNumbersSum = 0
	for _, part := range parts {
		if IsPartNumberRelevant(part, symbolLocations) {
			partNumbersSum = partNumbersSum + part.number
		}
	}

	// TODO: Iterate over all symbols, and
	//  a) count neighbouring numbers
	//	b) check if "*"
	//	c) add sum of gear ratio to total if a && b
	// NB: a and b sound like `GetSymbolData`, `IsGear`, and `GetGearRatio` methods

	fmt.Printf("Sum of all the active part numbers: %d\n", partNumbersSum)
}
