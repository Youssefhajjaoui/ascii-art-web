package Ascii_Justify

import (
	"strings"

	// Importing the outil package from ascii/pkg/Outil
	outil "ascii/pkg/Outil"
)

// Print_Justifie prints the text justified.
func Print_Justifie(newOutils *outil.Outils) {
	justifyedstring := ""
	isprint := false
	allwords := newOutils.SplitedWord
	for index, part := range allwords {
		if part != "" {
			newOutils.SplitedWord = []string{}
			newOutils.SplitedWord = append(newOutils.SplitedWord, part)
			part = strings.TrimSpace(part)
			justifiedPart := justifyPart(newOutils, part)
			if justifiedPart != "" {
				for i := 2; i <= 9; i++ {
					for _, char := range part {
						if char == ' ' {
							justifyedstring += justifiedPart
						} else {
							charLine := (int(char)-32)*9 + i
							if sliceOfline, ok := newOutils.LettersMap[charLine]; ok {
								justifyedstring += sliceOfline
								isprint = true
							}
						}
					}
					justifyedstring += "\n"
				}
				newOutils.Result += justifyedstring
				justifyedstring = ""
			} else {
				Print_Center(newOutils)
				isprint = true
			}
		} else if index <= len(allwords)-1 && part == "" {
			if index == len(allwords)-1 && !isprint {
			} else {
				newOutils.Result += "\n"
			}
		}
	}
}
