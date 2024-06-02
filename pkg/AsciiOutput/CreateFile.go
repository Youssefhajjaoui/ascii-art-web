package Ascii_Output

import (
	outil "ascii/pkg/Outil"
	"os"
)

// Create_File creates a new file with the given name and writes the result into it.
func Create_File(newOutils *outil.Outils) {
	// Write the result to a file with the given name.
	// The file is created with permissions 0644 (read and write for the owner, and read for others).
	os.WriteFile(newOutils.FileName, []byte(newOutils.Result), 0o644)
}
