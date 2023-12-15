package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	initTemplate(time.Now().Day())
}

func initTemplate(day int) {
	template :=
		`package main

func main() {
	
}

func part1(input []string) (sum int) {
	return 0
}

func part2(input []string) (sum int) {
	return 0
}
`
	err := os.Mkdir(fmt.Sprintf("%d", day), os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(fmt.Sprintf("%d/%d.go", day, day), []byte(template), os.ModePerm)
	if err != nil {
		panic(err)
	}
	_, err = os.Create(fmt.Sprintf("%d/input.txt", day))
	if err != nil {
		panic(err)
	}

}
