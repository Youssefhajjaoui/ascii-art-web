package art

import (
	"fmt"
	"os"
)

// Error_Msg prints an error message and exits the program.
func Error_Msg() {
	// Print the Specification string, followed by the usage instructions.
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	// Exit the program with status code 0.
	os.Exit(0)
}
