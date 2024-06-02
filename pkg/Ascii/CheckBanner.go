package art

import "strings"

// Check_Banner checks if the provided banner name is valid.
func Check_Banner(Arg string, Banner *string) {
	// If the argument doesn't end with ".txt", append ".txt" to it.
	if !strings.HasSuffix(Arg, ".txt") {
		Arg += ".txt"
	}
	// If the argument is not one of the valid banner names, display an error message.
	if Arg != "standard.txt" && Arg != "shadow.txt" && Arg != "thinkertoy.txt" {
		Error_Msg()
	} else {
		// If the argument is valid, assign it to the Banner pointer.
		*Banner = Arg
	}
}
