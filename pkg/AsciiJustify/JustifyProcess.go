package Ascii_Justify

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	// Importing the outil package from ascii/pkg/Outil
	outil "ascii/pkg/Outil"
)

// Getcols gets the number of columns in the terminal.
func Getcols(columnWidth *int) error {
	// Execute the command "tput cols" to get the number of columns
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}
	outputStr := strings.TrimSpace(string(out))
	cols, err := strconv.Atoi(outputStr)
	if err != nil {
		// If there's an error converting the output to an integer, return it
		return fmt.Errorf("failed to convert columns to integer: %w", err)
	}
	// Assign the number of columns to columnWidth
	*columnWidth = cols
	return nil
}

// justifyPart justifies a part of the text.
func justifyPart(newOutils *outil.Outils, part string) string {
	var result string
	var printSpace string
	// For each character in the part
	for _, char := range part {
		// If the character is not a space
		if char != ' ' {
			// Calculate the line of the character in the letters map
			charLine := ((int(char) - 32) * 9) + 2
			// If the line exists in the letters map, add it to the result
			if sliceOfLine, ok := newOutils.LettersMap[charLine]; ok {
				result += sliceOfLine
			}
		}
	}
	padding := newOutils.Cols - len(result)
	newSplit := strings.Split(part, " ")
	if len(newSplit) != 1 && padding > 0 {
		newSpaces := padding / (len(newSplit) - 1)
		printSpace = strings.Repeat(" ", newSpaces)
		return printSpace
	} else {
		return ""
	}
}
