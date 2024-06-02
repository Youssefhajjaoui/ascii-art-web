package Ascii_Justify

import (
	"fmt"
	"os"
	"strings"

	// Importing the outil package from ascii/pkg/Outil
	outil "ascii/pkg/Outil"
)

// Print_Right prints the text right-justified.
func Print_Right(newOutils *outil.Outils) {
	justifiedstring := ""
	isprint := false
	// For each part in the split word
	for index, part := range newOutils.SplitedWord {
		// If the part is not an empty string
		if part != "" {
			for i := 2; i <= 9; i++ {
				result := ""
				for _, char := range part {
					charLine := (int(char)-32)*9 + i
					if sliceOfline, ok := newOutils.LettersMap[charLine]; ok {
						result += sliceOfline
					}
				}
				padding := newOutils.Cols - len(result)
				if padding > 0 {
					paddedResult := fmt.Sprintf("%s%s", strings.Repeat(" ", padding), result)
					justifiedstring += paddedResult + "\n"
					isprint = true
				} else {
					Print_Ascii(newOutils)
					os.Exit(0)
				}
			}
		} else if index <= len(newOutils.SplitedWord)-1 && part == "" {
			if index == len(newOutils.SplitedWord)-1 && !isprint {
			} else {
				justifiedstring += "\n"
			}
		}
	}
	// Assign the justified string to the result
	newOutils.Result = justifiedstring
}
