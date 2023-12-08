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
	
}`
	err := os.Mkdir(fmt.Sprintf("%d", day), os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(fmt.Sprintf("%d/%d.go", day, day), []byte(template), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
