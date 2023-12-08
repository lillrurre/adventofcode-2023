package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("input/8")

	s := bufio.NewScanner(f)

	runeToInt := func(r rune) (n int) {
		if r == 'L' {
			return 1
		}
		return 2
	}

	// Directions as int
	s.Scan()
	line := s.Text()
	directions := make([]int, len(line))
	for i, r := range line {
		directions[i] = runeToInt(r)
	}

	re := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)
	network := make(map[string][2]string)
	for s.Scan() {
		line = s.Text()
		if line == "" {
			continue
		}
		matches := re.FindStringSubmatch(line)
		network[matches[1]] = [2]string{matches[2], matches[3]}
	}

	fmt.Println(part1(directions, network, "AAA", "ZZZ"))
	fmt.Println(part2(directions, network))
}

func part1(directions []int, network map[string][2]string, start string, targets ...string) int {
	moves := 0
	current := start
	for !contains(targets, current) {
		moves++
		move := directions[(moves-1)%len(directions)]
		next := network[current][move-1]
		current = next
	}
	return moves
}

func contains(arr []string, item string) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}
	return false
}

func part2(directions []int, network map[string][2]string) int {
	starts := make([]string, 0)
	targets := make([]string, 0)
	for key := range network {
		if key[2] == 'A' {
			starts = append(starts, key)
		}
		if key[2] == 'Z' {
			targets = append(targets, key)
		}
	}

	periods := make([]int, len(starts))
	for i, start := range starts {
		periods[i] = part1(directions, network, start, targets...)
	}

	return lcm(periods)
}

func lcm(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = (result * numbers[i]) / gcd(result, numbers[i])
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}
