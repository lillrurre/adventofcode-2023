package main

import (
	"fmt"
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
		case a.xv > 0 && a.x > x, a.xv < 0 && a.x < x:
			return false
		case b.xv > 0 && b.x > x, b.xv < 0 && b.x < x:
			return false
		case a.yv > 0 && a.y > y, a.yv < 0 && a.y < y:
			return false
		case b.yv > 0 && b.y > y, b.yv < 0 && b.y < y:
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
	hailstones := parse(input)
	rock := throwRock(hailstones, 250.0)
	fmt.Println(rock)
	return int(rock.x + rock.y + rock.z)
}

func throwRock(hailstones []hailstone, rng float64) hailstone {
	h1, h2 := hailstones[0], hailstones[1]
	for xv := -rng; xv <= rng; xv++ { // will always work in the last iteration, x == 250 for the real data...so this loop could be removed
		for yv := -rng; yv <= rng; yv++ {
			for zv := -rng; zv <= rng; zv++ {

				a := h1.xv - xv
				b := h1.yv - yv
				c := h2.xv - xv
				d := h2.yv - yv

				t := (d*(h2.x-h1.x) - c*(h2.y-h1.y)) / ((a * d) - (b * c))

				x := h1.x + h1.xv*t - xv*t
				y := h1.y + h1.yv*t - yv*t
				z := h1.z + h1.zv*t - zv*t

				found := true
				for _, h := range hailstones {
					u := (x - h.x) / (h.xv - xv)
					if (x+u*xv != h.x+u*h.xv) || (y+u*yv != h.y+u*h.yv) || (z+u*zv != h.z+u*h.zv) {
						found = false
						break
					}
				}
				if found {
					return hailstone{x: x, xv: xv, y: y, yv: yv, z: z, zv: zv}
				}
			}
		}
	}
	panic("lol")
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

func (h hailstone) String() string {
	return fmt.Sprintf("Location=(%2.0f %2.0f %2.0f) Velocity=(%2.0f %2.0f %2.0f)", h.x, h.y, h.z, h.xv, h.yv, h.zv)
}
