package art

import (
	"fmt"

	// Importing the necessary packages
	Justify "ascii/pkg/AsciiJustify"
	output "ascii/pkg/AsciiOutput"
	outil "ascii/pkg/Outil"
)

// Process_Flags processes the flags provided by the user.
func Process_Flags(newOutils *outil.Outils) {
	// Define a map for output functions
	outpute := map[string]func(newOutils *outil.Outils){
		"--output=": output.Create_File, // If the flag is "--output=", call the Create_File function
	}
	// Define a map for justify functions
	justify := map[string]func(newOutils *outil.Outils){
		"right":   Justify.Print_Right,    // If the position is "right", call the Print_Right function
		"left":    Justify.Print_Ascii,    // If the position is "left", call the Print_Ascii function
		"center":  Justify.Print_Center,   // If the position is "center", call the Print_Center function
		"justify": Justify.Print_Justifie, // If the position is "justify", call the Print_Justifie function
	}
	// If the position exists in the justify map, call the corresponding function
	if justifyFunc, ok := justify[newOutils.Position]; ok {
		justifyFunc(newOutils)
		// If the position is not "left", print the result
		if newOutils.Position != "left" {
			fmt.Print(newOutils.Result)
		}
	}
	// If the flag exists in the output map, call the corresponding function
	if outputfunc, ok := outpute[newOutils.Flag]; ok {
		Justify.Print_Ascii(newOutils)
		outputfunc(newOutils)
	}
}
