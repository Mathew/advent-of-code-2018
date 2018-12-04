package main

import (
	"fmt"
	"github.com/mathew/advent-of-code-2018"
	"strconv"
	"strings"
)

const filePath = "3/data.txt"

type claim struct {
	id     string
	posX   int
	posY   int
	width  int
	height int
}

type fabric map[int]map[int][]claim

func (f fabric) markClaim(c claim) {
	posX, posY := c.posX, c.posY

	for x := posX; x < posX+c.width; x++ {

		// check and initialise nested map
		if f[x] == nil {
			f[x] = make(map[int][]claim)
		}

		for y := posY; y < posY+c.height; y++ {
			f[x][y] = append(f[x][y], c)
		}
	}
}

func (f fabric) countDuplicateAreaClaims() int {
	fabricSquaresClaimedMoreThanOnce := 0

	for _, xv := range f {
		for _, yv := range xv {
			if len(yv) > 1 {
				fabricSquaresClaimedMoreThanOnce++
			}
		}
	}

	return fabricSquaresClaimedMoreThanOnce
}

func (f fabric) getClaimWithoutOverlap() string {
	noOverlapClaims := make(map[string]claim)
	overlappedClaims := make(map[string]claim)

	for _, xv := range f {
		for _, yv := range xv {
			if len(yv) == 1 {
				if _, ok := overlappedClaims[yv[0].id]; !ok {
					noOverlapClaims[yv[0].id] = yv[0]
				}
			} else if len(noOverlapClaims) > 0 {
				for _, overlappedClaim := range yv {
					if _, ok := noOverlapClaims[overlappedClaim.id]; ok {
						delete(noOverlapClaims, overlappedClaim.id)
					}

					overlappedClaims[overlappedClaim.id] = overlappedClaim
				}
			}
		}
	}

	for k, _ := range noOverlapClaims {
		return k
	}

	return "None"
}

// facepalm
func parseRawClaims(raw string) []claim {
	rawClaims := strings.Split(raw, "\n")

	var claims []claim

	for _, rc := range rawClaims {
		splitRc := strings.Split(rc, "@")
		id, rest := strings.TrimSpace(splitRc[0]), strings.TrimSpace(splitRc[1])

		splitRc = strings.Split(rest, ":")
		coords, area := strings.TrimSpace(splitRc[0]), strings.TrimSpace(splitRc[1])

		splitRc = strings.Split(coords, ",")
		posX, _ := strconv.Atoi(splitRc[0])
		posY, _ := strconv.Atoi(splitRc[1])

		splitRc = strings.Split(area, "x")
		width, _ := strconv.Atoi(splitRc[0])
		height, _ := strconv.Atoi(splitRc[1])

		claims = append(claims, claim{
			id:     id,
			posX:   posX,
			posY:   posY,
			width:  width,
			height: height,
		})
	}

	return claims
}

func main() {
	rawClaims := files.LoadFile(filePath)
	claims := parseRawClaims(rawClaims)
	fabric := fabric{}

	for _, c := range claims {
		fabric.markClaim(c)
	}

	fmt.Println(fabric.countDuplicateAreaClaims())
	fmt.Println(fabric.getClaimWithoutOverlap())
}
