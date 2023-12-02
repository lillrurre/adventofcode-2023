package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	red, green, blue int
}

func main() {
	f, _ := os.Open("input/2")
	s := bufio.NewScanner(f)
	re := regexp.MustCompile(`(\d+) (\w+)`)

	gameNum, part1, part2 := 1, 0, 0
	for s.Scan() {
		g := new(game)
		for _, match := range re.FindAllStringSubmatch(strings.TrimSpace(s.Text()), -1) {
			n, _ := strconv.Atoi(match[1])
			color := match[2]
			fmt.Println(color)
			switch {
			case color == "red" && g.red < n:
				g.red = n
			case color == "green" && g.green < n:
				g.green = n
			case color == "blue" && g.blue < n:
				g.blue = n
			}
		}
		if g.red <= 12 && g.green <= 13 && g.blue <= 14 {
			part1 += gameNum
		}
		gameNum++

		part2 += g.red * g.green * g.blue
	}
	fmt.Printf("[1] Result: %d\n", part1)
	fmt.Printf("[2] Result: %d\n", part2)
}
