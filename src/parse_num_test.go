package main

import (
	"testing"
)

func TestParseNumToWords(t *testing.T) {
	// A list of Structs for all test cases
	testCases := []struct {
		name          string
		input         int
		expectedWords string
		expectedError bool
	}{
		{name: "Zero", input: 0, expectedWords: "zero", expectedError: false},
		{name: "Single Digit", input: 7, expectedWords: "seven", expectedError: false},
		{name: "Double Digit", input: 29, expectedWords: "twenty-nine", expectedError: false},
		{name: "Teens", input: 13, expectedWords: "thirteen", expectedError: false},
		{name: "Double Digit without a one's place", input: 40, expectedWords: "forty", expectedError: false},
		{name: "Hundreds", input: 200, expectedWords: "two hundred", expectedError: false},
		{name: "Hundreds with ones", input: 109, expectedWords: "one hundred and nine", expectedError: false},
		{name: "Hundred and teens", input: 815, expectedWords: "eight hundred and fifteen", expectedError: false},
		{name: "Negative Number", input: -213, expectedWords: "", expectedError: true},
	}

	// Looping over the test cases
	for _, test := range testCases {
		// Running test for each test case
		t.Run(test.name, func(t *testing.T) {
			actualWords, err := parseNumToWords(test.input)

			// Checking for unexpected errors
			if !test.expectedError && err != nil {
				t.Fatalf("for input: %d, received unexpected error: %v", test.input, err)
			}

			// Checking for errors not triggering
			if test.expectedError && err == nil {
				t.Fatalf("expected an error with input: %d, but received none", test.input)
			}

			// Checking correctness of the output
			if actualWords != test.expectedWords {
				t.Fatalf("For input %d, expected to get: %s, but received: %s", test.input, test.expectedWords, actualWords)
			}
		})
	}
}
