package main

import (
	"fmt"
	"github.com/mathew/advent-of-code-2018"
	"strings"
)

const filePath = "5/data.txt"
const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var count = 0

func oppositePolarities(l1 string, l2 string) bool {
	if l1 == l2 {
		return false
	}

	if strings.ToLower(l1) == strings.ToLower(l2) {
		count++
		return true
	}

	return false
}

func destroyPolaritiesInPolymer(rawPolymer string, startPos int) (string, int) {

	successTo := 0
	for i := startPos; i+1 < len(rawPolymer); i++ {
		successTo = i
		fmt.Println(startPos)
		if oppositePolarities(string(rawPolymer[i]), string(rawPolymer[i+1])) {
			successTo--
			if i == 0 {
				successTo = 0
			}
			return fmt.Sprint(rawPolymer[:i], rawPolymer[i+2:]), successTo
		}
	}

	return rawPolymer, successTo
}

func partOne(rawPolymer string) int {
	polymer := rawPolymer
	previousPolymer := rawPolymer
	startPos := 0

	for true {
		polymer, startPos = destroyPolaritiesInPolymer(previousPolymer, startPos)

		if previousPolymer == polymer {
			break
		}

		previousPolymer = polymer
	}

	return len(polymer)
}

func partTwo(rawPolymer string) (int, string) {
	strippedPolymer := rawPolymer

	smallestPolymerLength := -1
	smallestCharacter := ""
	length := 0
	stringLetter := ""

	for _, letter := range letters {
		stringLetter = string(letter)
		fmt.Println(stringLetter)
		strippedPolymer = strings.Replace(rawPolymer, stringLetter, "", -1)
		strippedPolymer = strings.Replace(strippedPolymer, strings.ToLower(stringLetter), "", -1)
		length = partOne(strippedPolymer)

		if smallestPolymerLength == -1 || smallestPolymerLength > length {
			smallestPolymerLength = length
			smallestCharacter = stringLetter
		}
	}

	return smallestPolymerLength, smallestCharacter

}

func main() {
	rawPolymer := strings.TrimSpace(files.LoadFile(filePath))

	fmt.Println("The Part One answer is: ", partOne(rawPolymer))
	smallestLength, _ := partTwo(rawPolymer)
	fmt.Println("The Part Two answer is: ", smallestLength)
}
