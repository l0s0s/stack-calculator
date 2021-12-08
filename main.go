package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/l0s0s/stack-calculator/calc"
	"github.com/l0s0s/stack-calculator/stack"
)

func main() {
	result, err := calc.NewCalculator([]rune("+-*/"), stack.New()).Calculate(strings.Split(os.Args[1], " "))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
