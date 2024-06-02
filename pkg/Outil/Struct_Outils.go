package Outil

// Outils is a struct that contains various fields used throughout the ASCII program.
type Outils struct {
	Banner       string         // Banner is the banner text to be displayed.
	Args         []string       // Args are the command-line arguments.
	Flag         string         // Flag is the command-line flag.
	Position     string         // Position is the position of the text (e.g., "center", "right").
	FileName     string         // FileName is the name of the file to write the result to.
	WordToChange string         // WordToChange is the word to be changed in the ASCII art.
	Cols         int            // Cols is the number of columns in the terminal.
	LettersMap   map[int]string // LettersMap is a map of ASCII representations of characters.
	SplitedWord  []string       // SplitedWord is the word to be displayed, split into parts.
	Result       string         // Result is the final ASCII art string.
}
