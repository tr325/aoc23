package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)


type Mappe struct {
	directives []Directive
	source string
	desination string
}

type Directive struct {
	sourceRangeStart int
	destinationRangeStart int
	rangeLength int
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

func ParseSeedsInputLine(line string) []int {
	return []int{1}
}

func ParseMappeTitleLine(line string) (string, string) {
	return "a", "b"
}

func ParseMappeDirectiveLine(line string) (int, int, int) {
	return 1, 2, 3
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
