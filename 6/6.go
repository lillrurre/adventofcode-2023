package main

import (
	"fmt"
	"github.com/lillrurre/adventofcode-2023/util"
	"math"
	"regexp"
)

func main() {
	solve := func(durations []int, distances []int) int {
		res := 1
		for i, dur := range durations {
			for j := 0; j < dur/2; j++ {
				if (dur-j)*j > distances[i] {
					res *= dur - j*2 + 1
					break
				}
			}
		}
		return res
	}

	input := util.FileAsStringArr(6, "\n")
	durations, distances := parse(input)
	util.Run(1, func() any {
		return solve(durations, distances)
	})

	distance := ""
	for _, d := range distances {
		distance = fmt.Sprintf("%s%d", distance, d)
	}
	duration := ""
	for _, d := range durations {
		duration = fmt.Sprintf("%s%d", duration, d)
	}

	util.Run(2, func() any {
		return solve([]int{util.Atoi(duration)}, []int{util.Atoi(distance)})
	})

	util.Run(3, func() any {
		return idealSolution(util.Atoi(duration), util.Atoi(distance))
	})
}

func parse(input []string) ([]int, []int) {

	re := regexp.MustCompile(`\d+`)
	duration := make([]int, 0)
	for _, match := range re.FindAllString(input[0], -1) {
		duration = append(duration, util.Atoi(match))
	}

	distance := make([]int, 0)
	for _, match := range re.FindAllString(input[1], -1) {
		distance = append(distance, util.Atoi(match))
	}

	return duration, distance
}

// After reading the AOC reddit thread for day 6:

func idealSolution(duration, distance int) int {
	b := float64(duration)
	c := float64(distance)

	// √b*b-4*a*c
	d := math.Sqrt(math.Pow(b, 2) - 4*c)

	lowest := math.Floor((b - d) / 2)
	highest := math.Floor((b + d) / 2)

	return int(highest - lowest)
}
