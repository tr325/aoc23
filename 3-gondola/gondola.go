package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"regexp"
	"strconv"
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

	var row = 0
	parts := []Part{}
	symbolData := []SymbolData{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		ParseLine(line, &parts, &symbolData, row)
		row = row + 1
	}
	return parts, symbolData
}

func ParseLine(line string, parts *[]Part, symbolData *[]SymbolData, row int) {
	numberFinder := regexp.MustCompile(`[0-9]+`)
	symbolFinder := regexp.MustCompile(`[^A-Za-z0-9\.]`)

	numbers := numberFinder.FindAllString(line, -1)
	numberLocs := numberFinder.FindAllStringIndex(line, -1)
	symbols := symbolFinder.FindAllString(line, -1)
	symbolLocations := symbolFinder.FindAllStringIndex(line, -1)
	for i, num := range numbers {
		number, _ := strconv.Atoi(num)
		*parts = append(*parts, Part{number, row, numberLocs[i][0], numberLocs[i][1] - 1})
	}
	for j, loc := range symbolLocations {
		location := SymbolLocation{row, loc[0]}
		data := SymbolData{[]Part{}, symbols[j], location}
		*symbolData = append(*symbolData, data)
	}
}

// ------------------------------------------------------------
// Part 1

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isAdjacent(part Part, symbol SymbolLocation) bool {
	withinWidth := ((symbol.col + 1) >= part.start) && ((symbol.col - 1) <= part.end)
	adjacentRows := (abs(symbol.row-part.row) <= 1)
	return withinWidth && adjacentRows
}

func IsPartNumberRelevant(part Part, symbols []SymbolLocation) bool {
	for _, symbol := range symbols {
		if isAdjacent(part, symbol) {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------
// Part 2

func AddAdjacentParts(symbol *SymbolData, parts []Part) {
	for _, part := range parts {
		if isAdjacent(part, symbol.location) {
			symbol.adjacentParts = append(symbol.adjacentParts, part)
		}
	}
}

func IsGear(symbol SymbolData) bool {
	return symbol.symbol == "*" && len(symbol.adjacentParts) == 2
}

func GetGearRatio(symbol SymbolData) int {
	return symbol.adjacentParts[0].number * symbol.adjacentParts[1].number
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

	// Part 1
	var partNumbersSum = 0
	for _, part := range parts {
		if IsPartNumberRelevant(part, symbolLocations) {
			partNumbersSum = partNumbersSum + part.number
		}
	}

	// Part 2
	var gearRatiosSum = 0
	for _, symbol := range symbols {

		AddAdjacentParts(&symbol, parts)
		if IsGear(symbol) {
			gearRatiosSum = gearRatiosSum + GetGearRatio(symbol)
		}
	}

	fmt.Printf("Sum of all the active part numbers: %d\n", partNumbersSum)
	fmt.Printf("Sum of all the gear ratios: %d\n", gearRatiosSum)
}
