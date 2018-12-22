package main

import (
	"fmt"
	"github.com/mathew/advent-of-code-2018"
	"math"
	"sort"
	"strconv"
	"strings"
)

const filePath = "6/data.txt"

type location struct {
	x int
	y int
}

type locationCount struct {
	loc         location
	coords      []location
	count       int
	touchesEdge bool
}

func (l location) calculateDistance(otherLocation location) int {
	xDist := math.Abs(float64(l.x)) - math.Abs(float64(otherLocation.x))
	yDist := math.Abs(float64(l.y)) - math.Abs(float64(otherLocation.y))

	return int(math.Abs(float64(xDist))) + int(math.Abs(float64(yDist)))
}

func makeLocations(rawData string) []location {
	data := strings.Split(rawData, "\n")
	var coords []location

	for _, d := range data {
		s := strings.Split(d, ", ")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])

		coords = append(coords, location{x: x, y: y,})
	}

	return coords
}

func getGridSize(locations []location) int {
	x, y := 0, 0

	for _, coord := range locations {
		x = int(math.Max(float64(coord.x), float64(x)))
		y = int(math.Max(float64(coord.y), float64(y)))
	}

	max := int(math.Max(float64(x), float64(y)))

	return max
}

func makeCoordLocations(gridSize int) []location {

	coordLocations := []location{}

	for i := 0; i <= gridSize; i++ {
		for j := 0; j <= gridSize; j++ {
			coordLocations = append(coordLocations, location{x: i, y: j})
		}
	}

	return coordLocations
}

func getMaxLocationCount(ownedMap map[location]int, excludedLocations []location) int {
	var llcc []locationCount
	for l, c := range ownedMap {
		llcc = append(llcc, locationCount{loc: l, count: c})
	}

	sort.Slice(llcc, func(i int, j int) bool {
		return llcc[i].count > llcc[j].count
	})

	for _, lc := range llcc {
		ignore := false
		for _, excluded := range excludedLocations {
			if excluded.y == lc.loc.y && excluded.x == lc.loc.x {
				ignore = true
			}
		}

		if !ignore {
			return lc.count
		}
	}

	return 0
}

func isEdgeLocation(loc location, gridSize int) bool {
	return loc.y == 0 || loc.y == gridSize || loc.x == 0 || loc.x == gridSize
}

func calculateMostOwnedCoords(locations []location, coordLocations []location, gridSize int) int {

	ownedMap := make(map[location]int)
	excludedLocations := []location{}

	for _, coordLocation := range coordLocations {
		var closestDist int
		tie := false
		var closestLocation location

		for _, l := range locations {
			dist := l.calculateDistance(coordLocation)

			if dist == 0 {
				continue
			}

			if dist == closestDist {
				tie = true
				continue
			}

			if dist < closestDist || closestDist == 0 {
				closestDist = dist
				tie = false
				closestLocation = l
			}

			if dist < 0 {
				fmt.Println("This needs fixed.", dist)
			}
		}

		if !tie {
			ownedMap[closestLocation]++

			if isEdgeLocation(coordLocation, gridSize) {
				excludedLocations = append(excludedLocations, closestLocation)
			}
		}
	}

	return getMaxLocationCount(ownedMap, excludedLocations) - 1
}

func getRegionSize(locations []location, coordLocations []location, distance int) int {
	regionSize := 0
	for _, coordLocation := range coordLocations {
		totalDistance := 0
		for _, l := range locations {
			totalDistance += l.calculateDistance(coordLocation)
		}

		if totalDistance < distance {
			regionSize++
		}
	}

	return regionSize
}

func main() {
	rawData := files.LoadFile(filePath)
	distance := 10000

	locations := makeLocations(rawData)
	coordLocations := makeCoordLocations(getGridSize(locations))

	fmt.Println("The answer to part one is: ", calculateMostOwnedCoords(locations, coordLocations, getGridSize(locations)))
	fmt.Println("The answer to part two is: ", getRegionSize(locations, coordLocations, distance))
}
