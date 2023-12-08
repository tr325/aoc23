package main

import (
	"bufio"
	"github.com/samber/lo"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mappe struct {
	directives []Directive
	source     string
	desination string
}

type Directive struct {
	sourceRangeStart      int
	destinationRangeStart int
	rangeLength           int
}

func parse(fileScanner *bufio.Scanner) ([]int, map[string]Mappe) {
	mapOfMappes := make(map[string]Mappe)
	var seeds []int

	var currentMappe Mappe
	mappeTitleLine := regexp.MustCompile(`^[a-z]+`)
	mappeDirectiveLine := regexp.MustCompile(`^[0-9]+`)

	for fileScanner.Scan() {
		line := strings.Trim(fileScanner.Text(), " ")
		if strings.HasPrefix(line, "seeds:") {
			seeds = ParseSeedsInputLine(line)
		} else if line == "" && currentMappe.source != "" {
			mapOfMappes[currentMappe.source] = currentMappe
		} else if mappeTitleLine.MatchString(line) {
			source, destination := ParseMappeTitleLine(line)
			currentMappe = Mappe{
				[]Directive{},
				source,
				destination,
			}
		} else if mappeDirectiveLine.MatchString(line) {
			source, destination, length := ParseMappeDirectiveLine(line)
			directive := Directive{
				source,
				destination,
				length,
			}
			currentMappe.directives = append(currentMappe.directives, directive)
		}
	}
	// No newline at end of file --> EOF ends the final Mappe
	mapOfMappes[currentMappe.source] = currentMappe

	return seeds, mapOfMappes
}

func splitOnWhitespace(list string) []string {
	whiteSpace := regexp.MustCompile(`\s+`)
	split := whiteSpace.Split(list, -1)
	// Remove any stray empty string elements
	return lo.Filter(split, func(s string, _ int) bool {
		return s != ""
	})
}

func parseListOfInts(list string) []int {
	filteredList := splitOnWhitespace(list)
	return lo.Map(filteredList, func(s string, _ int) int {
		number, _ := strconv.Atoi(s)
		return number
	})
}

func ParseSeedsInputLine(line string) []int {
	list := strings.Replace(line, "seeds:", "", -1)
	return parseListOfInts(list)
}

func ParseMappeTitleLine(line string) (string, string) {
	lineParts := splitOnWhitespace(line)
	mappeNames := strings.Split(lineParts[0], "-to-")
	return mappeNames[0], mappeNames[1]
}

func ParseMappeDirectiveLine(line string) (int, int, int) {
	nums := parseListOfInts(line)
	return nums[0], nums[1], nums[2]
}

// ------------------------------------------------------------
// Part 1

// ------------------------------------------------------------
// Part 2

// ------------------------------------------------------------

func main() {
	readFile, _ := os.Open("almanac.txt")

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// seeds, mapOfMappes := parse(fileScanner)

	// fmt.Printf("Lowest location number for initial seeds: %d\n", lowestLocation)
}
