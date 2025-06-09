package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Taking the number as an argument from the command line
	arg := os.Args[1:]

	// Handling too many or no arguments
	if len(arg) > 1 {
		log.Fatal("This CLI only takes one argument")
	} else if len(arg) < 1 {
		log.Fatal("Please enter a number in your argument to this CLI")
	}

	// Converting the argument from string to integer
	num, err := strconv.Atoi(arg[0])

	// Handling errors if it is not a string
	if err != nil {
		log.Fatal("Please provide a valid integer\n", err)
	}

	// Calling the parseNumToWords function on the number
	words, err := parseNumToWords(num)

	// Handling any unexpected errors
	if err != nil {
		log.Fatal(err)
	}

	// Printing the words to the terminal
	fmt.Println(words)
}
