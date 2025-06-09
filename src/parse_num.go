package main

import (
	"fmt"
	"strings"
)

// Creating a map to bind the one's place integers to their string value
var onesPlace = map[int]string{
	1: "one", 2: "two", 3: "three",
	4: "four", 5: "five", 6: "six",
	7: "seven", 8: "eight", 9: "nine",
}

// Creating a map to bind the ten's place integers to their string value
var tensPlace = map[int]string{
	2: "twenty", 3: "thirty", 4: "forty",
	5: "fifty", 6: "sixty", 7: "seventy",
	8: "eighty", 9: "ninety",
}

// Creating a map to bind the 'teen' integers to their string value
var teensPlace = map[int]string{
	10: "ten", 11: "eleven", 12: "twelve",
	13: "thirteen", 14: "fourteen", 15: "fifteen",
	16: "sixteen", 17: "seventeen", 18: "eighteen",
	19: "nineteen",
}

// Function for converting chunks of the number
func convertChunkToWords(num int) string {

	// Quickly returns ones place chuncks
	if num < 10 {
		return onesPlace[num]
	}

	// Quickly returns the teens place chunks
	if num < 20 {
		return teensPlace[num]
	}

	// strings.Builder struct used for efficient string concatenation
	var wordBuilder strings.Builder

	// Look for the number in the hundred's place and remainders
	hundreds := num / 100
	remainder := num % 100

	// If a number in the hundred's place is present
	if hundreds > 0 {
		wordBuilder.WriteString(onesPlace[hundreds])
		wordBuilder.WriteString(" hundred")
		if remainder > 0 {
			// Handles adding "and" after "hundred", if remainder is not 0
			wordBuilder.WriteString(" and ")
		}
	}

	// Handling the special case of 'teens' early
	if remainder >= 10 && remainder <= 19 {
		wordBuilder.WriteString(teensPlace[remainder])
		// Returning early in case of teens
		return wordBuilder.String()
	}

	// Checking tens and ones place
	tens := remainder / 10
	ones := remainder % 10

	// If a number is present in the ten's place
	if tens > 0 {
		wordBuilder.WriteString(tensPlace[tens])
		if ones > 0 {
			// Adding '-' after ten's place, if one's place number is not 0 eg: "twenty-nine" for 29
			wordBuilder.WriteString("-")
			wordBuilder.WriteString(onesPlace[ones])
			return wordBuilder.String() // returning early to avoid the final one's place check
		}
	}

	// Final check for one's place
	if ones > 0 {
		wordBuilder.WriteString(onesPlace[ones])
	}

	// Return the string result for chunk, to the main parsing function
	return wordBuilder.String()
}

func parseNumToWords(num int) (string, error) {

	// Returning an error if the number is out of range
	if num > 100000 || num < 0 {
		return "", fmt.Errorf("input out of range: Please provide an integer from 0 to 99")
	}

	// Handling '0' early
	if num == 0 {
		return "zero", nil
	}

	// Handling '100,000' early
	if num == 100000 {
		return "one hundred thousand", nil
	}

	// Checking the thousand's place and the remainder
	thousands := num / 1000
	remainder := num % 1000

	// strings.Builder struct used for efficient string concatenation
	var wordBuilder strings.Builder

	// Checking if a number is present in the thousand's place
	if thousands > 0 {
		wordBuilder.WriteString(convertChunkToWords(thousands))
		wordBuilder.WriteString(" thousand")
	}

	// If remainder present
	if remainder > 0 {
		if thousands > 0 {
			if remainder < 100 {
				wordBuilder.WriteString(" and ")
			} else {
				wordBuilder.WriteString(", ")
			}
		}
		wordBuilder.WriteString(convertChunkToWords(remainder))
	}

	return wordBuilder.String(), nil

}
