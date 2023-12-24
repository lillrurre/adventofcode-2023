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
	return solve(hailstones, 200000000000000.0, 400000000000000.0)
}

func solve(hailstones []hailstone, a, b float64) (sum int) {
	// check if the intersection is inside the area
	inBounds := func(x, y float64) bool { return a <= x && b >= x && a <= y && b >= y }

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

	for i, h := range hailstones {
		for j := i + 1; j < len(hailstones); j++ {
			h1 := hailstones[j]
			x, y := intersection(h, h1)
			if inBounds(x, y) && valid(x, y, h, h1) {
				sum++
			}
		}
	}

	return sum
}

func intersection(h1, h2 hailstone) (float64, float64) {
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

func part2(input []string) (sum int) {
	return 0
}

func parse(input []string) (hailstones []hailstone) {
	for _, line := range input {
		parts := strings.FieldsFunc(line, func(r rune) bool { return r == '@' || r == ',' || r == ' ' })
		hailstones = append(hailstones, hailstone{
			x:  float64(util.Atoi(parts[0])),
			y:  float64(util.Atoi(parts[1])),
			z:  float64(util.Atoi(parts[2])),
			xv: float64(util.Atoi(parts[3])),
			yv: float64(util.Atoi(parts[4])),
			zv: float64(util.Atoi(parts[5])),
		})
	}
	return hailstones
}
