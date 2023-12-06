package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input/6")

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

	durations, distances := parse(string(b))
	fmt.Printf("[1] Result: %d\n", solve(durations, distances))

	distance := ""
	for _, d := range distances {
		distance = fmt.Sprintf("%s%d", distance, d)
	}
	duration := ""
	for _, d := range durations {
		duration = fmt.Sprintf("%s%d", duration, d)
	}

	fmt.Printf("[2] Result: %d\n", solve([]int{atoi(duration)}, []int{atoi(distance)}))

	// Test ideal solution with quadratic formula for part 2
	fmt.Printf("[3] Result: %d\n", idealSolution(atoi(duration), atoi(distance)))
}

func parse(input string) ([]int, []int) {

	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`\d+`)
	duration := make([]int, 0)
	for _, match := range re.FindAllString(lines[0], -1) {
		duration = append(duration, atoi(match))
	}

	distance := make([]int, 0)
	for _, match := range re.FindAllString(lines[1], -1) {
		distance = append(distance, atoi(match))
	}

	return duration, distance
}

func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
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
