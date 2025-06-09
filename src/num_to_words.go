package main

import (
	"fmt"
	"strings"
)

// Creating a map to bind the one's place integers to their string value
var onesPlace = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

// Creating a map to bind the ten's place integers to their string value
var tensPlace = map[int]string{
	2: "twenty",
	3: "thirty",
	4: "fourty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

func parseNumToWords(num int) (string, error) {
	var wordBuilder strings.Builder

	if num > 99 {
		return "", fmt.Errorf("input out of range: Please provide an integer from 0 to 99")
	}

	if tens := num / 10; tens > 1 {
		wordBuilder.WriteString(tensPlace[tens])
		wordBuilder.WriteString(" ")
	}

	remainder := num % 10

	if remainder > 0 {
		wordBuilder.WriteString(onesPlace[remainder])
	}

	return wordBuilder.String(), nil

}
