// Package art contains the main logic for processing ASCII art.
package art

import (
	"strings"

	outil "web/Outil"
)

// Process_Text is the main function for processing the input text into ASCII art.
func Process_Text(newOutils *outil.Outils) {
	// The word to change is split by newline characters.
	newOutils.SplitedWord = strings.Split(newOutils.WordToChange, "\r\n")
	// A map of ASCII art characters is created.
	Make_Map(newOutils)

	Print_Ascii(newOutils)
}
