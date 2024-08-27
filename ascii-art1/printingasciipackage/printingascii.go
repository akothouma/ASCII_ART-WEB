package printingasciipackage

import (
	"fmt"
	"strings"

	"ASCII-WEB/ascii-art1/mapPackage"
)

// Reads input text,gets the pattern convert it to ascii art

func PrintingAscii(text, patternFile string) (string, error) {
	// text = strings.ReplaceAll(text, "\n", "\\n")
	res := ""

	lines := strings.Split(text, "\r\n")
	asciiMap, err := mapPackage.AsciiMapping(patternFile)
	if err != nil {
		return "", fmt.Errorf("error mapping %v", patternFile)
	}

	count := 0
	for _, word := range lines { // case of multiple newlines
		if word == "" {
			count++
			if count < len(lines) {
				res += "\n"
			}
		} else {
			for n := 0; n < 8; n++ {
				for _, ch := range word {
					res += asciiMap[ch][n]
				}
				res += "\n"
			}
		}
	}
	return res, nil
}
