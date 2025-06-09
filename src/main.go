package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]

	if len(arg) > 1 {
		log.Fatal("This CLI only takes one argument")
	} else if len(arg) < 1 {
		log.Fatal("Please enter a number in your argument to this CLI")
	}

	num, err := strconv.Atoi(arg[0])

	if err != nil {
		log.Fatal("Please provide a valid integer\n", err)
	}

	words, err := parseNumToWords(num)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(words)
}
