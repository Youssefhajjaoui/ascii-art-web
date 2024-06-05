package art

import (
	"bufio"
	"os"

	"web/Outil"
	outil "web/Outil"
)

// Make_Map creates a map of ASCII art characters.
func Make_Map(myoutils *outil.Outils) {
	// Create line numbers for each character in the word to change.
	lineNumbers := CreateLineNumbers(myoutils.WordToChange, myoutils)
	// Open the banner file.
	standardFile, err := os.Open("static/" + myoutils.Banner + ".txt")
	if err != nil {
		// If there's an error opening the file, display an error message.
		myoutils.Erreur = "erreur of open"
	}
	defer standardFile.Close()
	// Create a new scanner for the file.
	scanner := bufio.NewScanner(standardFile)
	lineCount := 1
	// Scan each line in the file.
	for scanner.Scan() {
		line := scanner.Text()
		// If the line number exists in the map, replace it with the line text.
		if _, ok := lineNumbers[lineCount]; ok {
			lineNumbers[lineCount] = line
		}
		lineCount++
	}
	// Assign the map of line numbers to the LettersMap field.
	myoutils.LettersMap = lineNumbers
}

// CreateLineNumbers creates a map of line numbers for each character in the text.
func CreateLineNumbers(text string, myoutils *Outil.Outils) map[int]string {
	lineNumbers := make(map[int]string)
	// Loop through each line number for each character.
	for i := 2; i <= 9; i++ {
		for j := 0; j < len(text); j++ {
			// Calculate the line number for the character.
			charLine := (int(rune(text[j])-32) * 9) + i
			// If the line number is greater than 855, display an error message.
			if charLine > 855 {
				myoutils.Erreur = "erreur for iinput"
			}
			// If the line number doesn't exist in the map and is greater than 0, add it to the map.
			if _, ok := lineNumbers[charLine]; !ok && charLine > 0 {
				lineNumbers[charLine] = ""
			}
		}
	}
	return lineNumbers
}
