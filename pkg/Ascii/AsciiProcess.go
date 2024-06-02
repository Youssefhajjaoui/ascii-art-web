// Package art contains the main logic for processing ASCII art.
package art

import (
	"strings"

	justify "ascii/pkg/AsciiJustify"
	outil "ascii/pkg/Outil"
)

// Process_Text is the main function for processing the input text into ASCII art.
func Process_Text(newOutils *outil.Outils) {
	ArgsLen := len(newOutils.Args)

	// Switch case based on the number of arguments
	switch ArgsLen {
	case 1:
		// If there's only one argument, it's the word to change into ASCII art.
		// The banner style is set to standard.
		newOutils.WordToChange = newOutils.Args[0]
		newOutils.Banner = "standard.txt"
	case 2:
		// If there are two arguments, the first one is the word to change.
		// The second one is checked if it's a flag or a banner style.
		newOutils.WordToChange = newOutils.Args[0]
		Check_Flags(newOutils)
		if newOutils.Flag == "--output=" || newOutils.Flag == "--align=" {
			newOutils.Banner = "standard.txt"
			newOutils.WordToChange = newOutils.Args[1]
		} else {
			Check_Banner(newOutils.Args[1], &newOutils.Banner)
		}
	case 3:
		// If there are three arguments, the first one is checked if it's a flag,
		// the second one is the word to change, and the third one is the banner style.
		Check_Flags(newOutils)
		newOutils.WordToChange = newOutils.Args[1]
		Check_Banner(newOutils.Args[2], &newOutils.Banner)
	default:
		// If the number of arguments is not 1, 2, or 3, an error message is displayed.
		Error_Msg()
	}

	// The word to change is split by newline characters.
	newOutils.SplitedWord = strings.Split(newOutils.WordToChange, "\\n")
	// A map of ASCII art characters is created.
	Make_Map(newOutils)

	// If there's no output or align flag, the ASCII art is printed.
	// Otherwise, the flags are processed.
	if newOutils.Flag != "--output=" && newOutils.Flag != "--align=" {
		justify.Print_Ascii(newOutils)
	} else {
		if newOutils.FileName != "" && newOutils.Position == "" {
			Process_Flags(newOutils)
		} else if newOutils.Position != "" && newOutils.FileName == "" {
			justify.Getcols(&newOutils.Cols)
			Process_Flags(newOutils)
		}
	}
}
