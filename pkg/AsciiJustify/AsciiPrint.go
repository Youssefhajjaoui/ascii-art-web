package Ascii_Justify

import (
	"fmt"

	// Importing the outil package from ascii/pkg/Outil
	outil "ascii/pkg/Outil"
)

// Print_Ascii prints the ASCII art representation of the input text.
func Print_Ascii(newOutils *outil.Outils) {
	result := ""     // Initialize the result string
	isprint := false // Initialize the print flag
	// Loop through each part of the split word
	for index, part := range newOutils.SplitedWord {
		if part != "" {
			// If the part is not empty, loop through each line number
			for i := 2; i <= 9; i++ {
				// For each character in the part, calculate the line number
				for j := 0; j < len(part); j++ {
					charLine := (int(rune(part[j])-32) * 9) + i
					// If the line number exists in the map, add it to the result string
					if sliceOfline, ok := newOutils.LettersMap[charLine]; ok {
						result += sliceOfline
						isprint = true
					}
				}
				// Add a newline character to the result string
				result += "\n"
				isprint = true
			}
		} else if index <= len(newOutils.SplitedWord)-1 && part == "" {
			// If the part is empty and it's not the last part, add a newline character to the result string
			if index == len(newOutils.SplitedWord)-1 && !isprint {
			} else {
				result += "\n"
			}
		}
	}
	// If there's no output flag, print the result string
	if newOutils.Flag != "--output=" {
		fmt.Print(result)
	}
	// Assign the result string to the Result field
	newOutils.Result = result
}
