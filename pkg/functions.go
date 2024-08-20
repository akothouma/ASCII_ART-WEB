package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetLine(filename string, num int) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1 // Line numbers start from 1
	for scanner.Scan() {
		if lineNum == num {
			return scanner.Text()
		}
		lineNum++
	}

	return ""
}

func GenerateASCIIArt(inputs, banner string) string {
	var result strings.Builder
	inputs = strings.ReplaceAll(inputs, "\n", "\\n")
	slice := strings.Split(inputs, "\\n")
	// Split directly on newline

	// Iterate over each line of input
	for _, input := range slice {
		// Iterate over each line (8 lines per character)
		for line := 1; line < 9; line++ {
			// Iterate over each character in input
			for _, char := range input {
				pos := 1 + int(char-' ')*9 + line
				lineContent := GetLine(banner, pos)
				result.WriteString(lineContent)
			}
			result.WriteString("\n") // New line after each set of characters for this line
		}
	}

	return result.String()
}
