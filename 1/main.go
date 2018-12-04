package main

import (
	"fmt"
	"github.com/mathew/advent-of-code-2018"
	"strconv"
	"strings"
)

const filePath = "1/data.txt"

func main() {
	rawFrequencyChanges := files.LoadFile(filePath)
	frequencyChanges := convertInputToInts(rawFrequencyChanges)

	fmt.Println("The answer to part one is: ", partOne(frequencyChanges))
	fmt.Println("The answer to part two is: ", partTwo(frequencyChanges))
}

func convertInputToInts(newlineSeparatedInts string) []int {
	splitChanges := strings.Split(newlineSeparatedInts, "\n")

	convertedInput := make([]int, 0, len(splitChanges))

	for _, num := range splitChanges {
		trimmedNum := strings.TrimSpace(num)

		if trimmedNum == "" {
			continue
		}
		convertedNum, err := strconv.Atoi(trimmedNum)

		if err != nil {
			fmt.Println(num)
			fmt.Println(err)
		} else {
			convertedInput = append(convertedInput, convertedNum)
		}
	}

	return convertedInput
}

func partOne(frequencyChanges []int) int {
	frequency := 0

	for _, num := range frequencyChanges {
		frequency += num
	}

	return frequency
}

func partTwo(frequencyChanges []int) int {
	frequency := 0

	set := map[int]int{0: 1}

	for {
		for _, num := range frequencyChanges {
			frequency += num
			fmt.Println(frequency)
			set[frequency] += 1

			if set[frequency] == 2 {
				return frequency
			}

		}
	}

}
