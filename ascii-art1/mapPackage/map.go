package mapPackage

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

// AsciiMapping given a banner file, reads all graphics representations of the ASCII characters and
// returns a map of the ASCII character to the graphics representations of the ASCII character
func AsciiMapping(patternFile string) (map[rune][]string, error) {
	var splitted []string
	standardHash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadowHash := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoyHash := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"

	testfile, err := os.ReadFile(patternFile)
	if err != nil {
		return nil, fmt.Errorf("%v doesn't exist", patternFile)
	}

	testFileHash := fmt.Sprintf("%x", sha256.Sum256(testfile))

	if patternFile == "thinkertoy.txt" {
		if testFileHash != thinkertoyHash {
			return nil, fmt.Errorf("%v has been modified", patternFile)
		}
		splitted = strings.Split(string(testfile), "\r\n") // Thinkertoy uses \r\n
	} else {
		if testFileHash != standardHash && testFileHash != shadowHash {
			return nil, fmt.Errorf("%v has been modified", patternFile)
		}
		splitted = strings.Split(string(testfile), "\n")
	}

	if len(splitted) != 856 {
		fmt.Printf("error : %v file modified, exiting...\n", patternFile)
		os.Exit(1)
	}

	asciiMapping := make(map[rune][]string)
	startAscii := ' '
	for i := 1; i < len(splitted); {
		arrayString := []string{}
		for j := 0; j < 8; j++ {
			if i >= len(splitted) { // Avoid out-of-range errors
				break
			}
			arrayString = append(arrayString, splitted[i])
			i++
		}
		i++
		asciiMapping[startAscii] = arrayString
		startAscii++
	}

	return asciiMapping, nil
}
