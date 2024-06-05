package art

import (
	outil "web/Outil"
)

// Print_Ascii prints the ASCII art representation of the input text.
func Print_Ascii(newOutils *outil.Outils) { // Initialize the result string
	isprint := false // Initialize the print flag
	// Loop through each part of the split word
	for index, part := range newOutils.SplitedWord {
		if part != "" {
			printascii(part, newOutils)
			isprint = true
		} else if index <= len(newOutils.SplitedWord)-1 && part == "" {
			// If the part is empty and it's not the last part, add a newline character to the result string
			if index != len(newOutils.SplitedWord)-1 || isprint {
				newOutils.Result += "\n"
			}
		}
	}
}

func printascii(part string, newOutils *outil.Outils) {
	for i := 2; i <= 9; i++ {
		// For each character in the part, calculate the line number
		for j := 0; j < len(part); j++ {
			charLine := (int(rune(part[j])-32) * 9) + i
			// If the line number exists in the map, add it to the result string
			if sliceOfline, ok := newOutils.LettersMap[charLine]; ok {
				newOutils.Result += sliceOfline
			}
		}
		// Add a newline character to the result string
		newOutils.Result += "\n"
	}
}
