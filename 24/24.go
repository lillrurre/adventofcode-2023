package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"strings"
)

type hailstone struct {
	x, y, z    float64
	xv, yv, zv float64
}

func main() {
	util.Run(1, func() (sum int) { return part1(util.FileAsStringArr(24, "\n")) })
	util.Run(2, func() (sum int) { return part2(util.FileAsStringArr(24, "\n")) })
}

func part1(input []string) (sum int) {
	hailstones := parse(input)
	return intersectionsInArea(hailstones, 200000000000000.0, 400000000000000.0)
}

// intersectionsInArea finds the number of intersections of hailstone trails inside a given area
func intersectionsInArea(hailstones []hailstone, a, b float64) (sum int) {
	// intersection creates a lines of h1, h2 and returns the intersection - 0 is returned if they are parallel
	intersection := func(h1, h2 hailstone) (float64, float64) {
		slope1 := h1.yv / h1.xv
		slope2 := h2.yv / h2.xv

		// parallel -> no intersect
		if slope1 == slope2 {
			return 0, 0
		}

		intercept1 := h1.y - slope1*h1.x
		intercept2 := h2.y - slope2*h2.x

		intersectionX := (intercept2 - intercept1) / (slope1 - slope2)
		intersectionY := slope1*intersectionX + intercept1

		return intersectionX, intersectionY
	}

	// check if the intersection is inside the area
	inBounds := func(x, y float64) bool { return min(x, y) >= a && max(x, y) <= b }

	// Check if the intersections are valid - in the future, not the past
	valid := func(x, y float64, a, b hailstone) bool {
		switch {
		case a.xv > 0 && a.x > x:
			return false
		case a.xv < 0 && a.x < x:
			return false
		case b.xv > 0 && b.x > x:
			return false
		case b.xv < 0 && b.x < x:
			return false
		case a.yv > 0 && a.y > y:
			return false
		case a.yv < 0 && a.y < y:
			return false
		case b.yv > 0 && b.y > y:
			return false
		case b.yv < 0 && b.y < y:
			return false
		}
		return true
	}

	// loop through each hailstone and check the intersection with the rest
	for i, h1 := range hailstones {
		for j := i + 1; j < len(hailstones); j++ {
			h2 := hailstones[j]
			x, y := intersection(h1, h2)
			if inBounds(x, y) && valid(x, y, h1, h2) {
				sum++
			}
		}
	}

	return sum
}

func part2(input []string) (sum int) {
	_ = parse(input)
	return sum
}

func parse(input []string) (hailstones []hailstone) {
	for _, line := range input {
		parts := strings.FieldsFunc(line, func(r rune) bool { return r == '@' || r == ',' || r == ' ' })
		hailstones = append(hailstones, hailstone{
			x: float64(util.Atoi(parts[0])), y: float64(util.Atoi(parts[1])), z: float64(util.Atoi(parts[2])),
			xv: float64(util.Atoi(parts[3])), yv: float64(util.Atoi(parts[4])), zv: float64(util.Atoi(parts[5])),
		})
	}
	return hailstones
}
