package mapPackage

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAsciiMapping(t *testing.T) {
	filetext := "standard.txt"

	out, _ := AsciiMapping(filetext)
	expected := []string{
		"           ",
		`    /\     `,
		`   /  \    `,
		`  / /\ \   `,
		` / ____ \  `,
		`/_/    \_\ `,
		"           ",
		"           ",
	}
	if !reflect.DeepEqual(out['A'], expected) {
		fmt.Println("Test unsuccessful")
	} else {
		fmt.Println("Test passed successfully")
	}
}
