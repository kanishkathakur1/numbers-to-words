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

	fmt.Printf("Your arg is: %v, of type: %T", num, num)
}
