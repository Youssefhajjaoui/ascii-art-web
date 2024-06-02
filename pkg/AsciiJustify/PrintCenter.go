package Ascii_Justify

import (
	"fmt"
	"os"
	"strings"

	// Importing the outil package from ascii/pkg/Outil

	// Importing the outil package from ascii/pkg/Outil
	outil "ascii/pkg/Outil"
)

// Print_Center prints the text centered.
func Print_Center(newOutils *outil.Outils) {
	justifyedstring := ""
	isprint := false
	// For each part in the split word
	// fmt.Println(len(newOutils.SplitedWord))
	for index, part := range newOutils.SplitedWord {
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
					leftspaces := padding / 2
					rightspaces := padding - leftspaces
					paddedResult := fmt.Sprintf("%s%s%s", strings.Repeat(" ", leftspaces), result, strings.Repeat(" ", rightspaces))
					justifyedstring += paddedResult + "\n"
					isprint = true
				} else {
					Print_Ascii(newOutils)
					os.Exit(0)
				}

			}
			newOutils.Result += justifyedstring
			justifyedstring = ""
		} else if index <= len(newOutils.SplitedWord)-1 && part == "" {
			if index == len(newOutils.SplitedWord)-1 && !isprint {
			} else {
				newOutils.Result += "\n"
			}
		}
	}
}
