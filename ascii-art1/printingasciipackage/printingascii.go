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

	for i := 0; i < len(text); {
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'a' {
			return "", fmt.Errorf("Character not supported")
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'b' {
			l := len(text) - 2
			if i == 0 {
				text = text[i+2:]
			} else if i == l {
				text = text[:l]
			} else {
				text = text[:i-1] + text[i+2:]
				i = 0
			}
			continue
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 't' {
			text = text[:i] + "   " + text[i+2:]
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'v' {

			return "", fmt.Errorf("Character not supported")
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'f' {
			return "", fmt.Errorf("Character not supported")
		}
		if i+1 < len(text) && text[i] == '\\' && text[i+1] == 'r' {
			return "", fmt.Errorf("Character not supported")
		}
		if i+1 < len(text) && text[i] > 127 {
			return "", fmt.Errorf("Character not supported")
		}
		i++
	}
	lines := strings.Split(text, "\r\n")
	asciiMap, err := mapPackage.AsciiMapping(patternFile)
	if err != nil {
		return "", fmt.Errorf("Error mapping %v", patternFile)
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
