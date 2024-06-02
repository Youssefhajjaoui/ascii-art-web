package art

import (
	"strings"

	// Importing the outil package from ascii/pkg/Outil
	outil "ascii/pkg/Outil"
)

// Define constants for the flag prefixes
const (
	prefix1 = "--output="
	prefix2 = "--align="
)

// Check_Flags checks the flags provided by the user.
func Check_Flags(Outils *outil.Outils) {
	// If the first argument starts with the output prefix
	if strings.HasPrefix(Outils.Args[0], prefix1) {
		// If the argument doesn't contain a "/" and ends with ".txt" and its length is greater than 13
		if !strings.Contains(Outils.Args[0], "/") && strings.HasSuffix(Outils.Args[0], ".txt") && len(Outils.Args[0]) > 13 {
			// Set the flag to the output prefix and the file name to the part of the argument after the prefix
			Outils.Flag = prefix1
			Outils.FileName = Outils.Args[0][len(prefix1):]
		} else {
			// If the argument doesn't meet the conditions, display an error message
			Error_Msg()
		}
	} else if strings.HasPrefix(Outils.Args[0], prefix2) { // If the first argument starts with the align prefix
		// If the part of the argument after the prefix is "right", "left", "center", or "justify"
		if strings.EqualFold(Outils.Args[0][len(prefix2):], "right") || strings.EqualFold(Outils.Args[0][len(prefix2):], "left") || strings.EqualFold(Outils.Args[0][len(prefix2):], "center") || strings.EqualFold(Outils.Args[0][len(prefix2):], "justify") {
			// Set the flag to the align prefix and the position to the part of the argument after the prefix
			Outils.Flag = prefix2
			Outils.Position = Outils.Args[0][len(prefix2):]
		} else {
			// If the part of the argument after the prefix is not one of the valid positions, display an error message
			Error_Msg()
		}
	} else if len(Outils.Args) == 3 { // If there are three arguments
		// Display an error message indicating that the first argument is not a flag
		Error_Msg()
	}
}
