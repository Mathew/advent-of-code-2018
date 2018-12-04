package main

import (
	"fmt"
	"github.com/mathew/advent-of-code-2018"
	"strings"
)

const filePath = "2/data.txt"

func parseBoxIds(s string) []string {
	return strings.Split(s, "\n")
}

func containsExactNoOfLetters(s string, noOfLettersRequired int) bool {
	letters := map[int32]int{}

	for _, l := range s {
		letters[l] += 1
	}

	for _, number := range letters {
		if number == noOfLettersRequired {
			return true
		}
	}

	return false
}

func calculateChecksum(boxIds []string) int {
	// MEMORY EFFICIENT *HIGH FIVE*
	counts := make(map[int]int, 2)
	counts = map[int]int{
		2: 0, 3: 0,
	}

	for _, b := range boxIds {
		if containsExactNoOfLetters(b, 2) {
			counts[2]++
		}

		if containsExactNoOfLetters(b, 3) {
			counts[3]++
		}
	}

	return counts[2] * counts[3]
}

func doStringsMatch(s1 string, s2 string, letterComparator func(l1 int32, l2 int32) bool) bool {
	length := len(s1)

	for i := 0; i < length; i++ {
		if !letterComparator(rune(s1[i]), rune(s2[i])) {
			return false
		}
	}

	return true
}

// Closure for handling the off by one difference
func offByOneRuneComparator() func(int32, int32) bool {
	offByOneCount := 0

	return func(l1 int32, l2 int32) bool {
		if l1 == l2 {
			return true
		}

		if l1 != l2 {
			offByOneCount++
			if offByOneCount <= 1 {
				return true
			}

		}

		return false
	}
}

func calculateMatchingLetters(s1 string, s2 string) []string {
	length := len(s1)
	var matchingLetters []string

	for i := 0; i < length; i++ {
		if rune(s1[i]) == rune(s2[i]) {
			matchingLetters = append(matchingLetters, string(s1[i]))
		}
	}

	return matchingLetters
}

func calculateCommonBoxLetters(boxIds []string) []string {
	var box string
	var otherBox string

	for i := 0; i < len(boxIds); i++ {
		box = boxIds[i]

		for j := 0; j < len(boxIds); j++ {
			// don't compare the same box
			if i == j {
				continue
			}
			otherBox = boxIds[j]

			if doStringsMatch(box, otherBox, offByOneRuneComparator()) {
				return calculateMatchingLetters(box, otherBox)
			}
		}
	}

	fmt.Println("No match")
	return nil
}

func main() {
	rawBoxIds := files.LoadFile(filePath)
	boxIds := parseBoxIds(rawBoxIds)

	fmt.Println(calculateChecksum(boxIds))
	fmt.Println(strings.Join(calculateCommonBoxLetters(boxIds), ""))
}
