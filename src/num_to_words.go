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
	2: "twenty", 3: "thirty", 4: "fourty",
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

func parseNumToWords(num int) (string, error) {

	if num > 1000 || num < 0 {
		return "", fmt.Errorf("input out of range: Please provide an integer from 0 to 99")
	}

	if num == 0 {
		return "zero", nil
	}

	if num == 1000 {
		return "one thousand", nil
	}

	var wordBuilder strings.Builder

	hundreds := num / 100

	if hundreds > 0 {
		wordBuilder.WriteString(onesPlace[hundreds])
		wordBuilder.WriteString(" hundred")
	}

	remainder := num % 100

	if remainder > 0 {

		if wordBuilder.Len() > 0 {
			wordBuilder.WriteString(" and ")
		}

		if remainder >= 10 && remainder <= 19 {
			wordBuilder.WriteString(teensPlace[remainder])
		} else {

			tens := remainder / 10

			if tens > 0 {
				wordBuilder.WriteString(tensPlace[tens])
			}

			ones := remainder % 10

			if wordBuilder.Len() > 0 && tens > 0 && ones > 0 {
				wordBuilder.WriteString("-")
			}

			if ones > 0 {
				wordBuilder.WriteString(onesPlace[ones])
			}
		}

	}

	return wordBuilder.String(), nil

}
