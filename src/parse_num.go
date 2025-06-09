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
	if num == 0 {
		return ""
	}

	if num < 10 {
		return onesPlace[num]
	}

	if num < 20 {
		return teensPlace[num]
	}

	var wordBuilder strings.Builder

	hundreds := num / 100
	remainder := num % 100

	if hundreds > 0 {
		wordBuilder.WriteString(onesPlace[hundreds])
		wordBuilder.WriteString(" hundred")
		if remainder > 0 {
			wordBuilder.WriteString(" and ")
		}
	}

	if remainder >= 10 && remainder <= 19 {
		wordBuilder.WriteString(teensPlace[remainder])
		return wordBuilder.String()
	}

	tens := remainder / 10
	ones := remainder % 10

	if tens > 0 {
		wordBuilder.WriteString(tensPlace[tens])
		if ones > 0 {
			wordBuilder.WriteString("-")
			wordBuilder.WriteString(onesPlace[ones])
			return wordBuilder.String()
		}
	}

	if ones > 0 {
		wordBuilder.WriteString(onesPlace[ones])
	}

	return wordBuilder.String()
}

func parseNumToWords(num int) (string, error) {

	if num > 100000 || num < 0 {
		return "", fmt.Errorf("input out of range: Please provide an integer from 0 to 99")
	}

	if num == 0 {
		return "zero", nil
	}

	if num == 100000 {
		return "one hundred thousand", nil
	}

	thousands := num / 1000
	remainder := num % 1000

	var wordBuilder strings.Builder

	if thousands > 0 {
		wordBuilder.WriteString(convertChunkToWords(thousands))
		wordBuilder.WriteString(" thousand")
	}

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
