package printingasciipackage

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintingAscii(t *testing.T) {
	args := "hello" // string to be patterned and printed
	file := "standard.txt"
	out := PrintingAscii(args, file)
	expected := []string{
		` _              _   _          `,
		`| |            | | | |         `,
		`| |__     ___  | | | |   ___   `,
		`|  _ \   / _ \ | | | |  / _ \  `,
		`| | | | |  __/ | | | | | (_) | `,
		`|_| |_|  \___| |_| |_|  \___/  `,
		"                               ",
		"                               ",
	}
	// PrintingAscii function returns a string of len 256
	// we have 8 lines ,so length of each line is (256/8)=32
	// we access from index 0 to 255
	if !reflect.DeepEqual(out[:31], expected[0]) {
		fmt.Println("Test  line 1 failed")
		t.Errorf("got\n %v, want %v", out[:31], expected[0])
		return
	} else {
		fmt.Println(" line 1 Test passed successfully")
	}
	if !reflect.DeepEqual(out[32:63], expected[1]) {
		fmt.Println("Test line 2 failed")
		t.Errorf("got\n %v, want %v", out[32:63], expected[1])
		return
	} else {
		fmt.Println(" line 2 Test passed successfully")
	}
	if !reflect.DeepEqual(out[64:95], expected[2]) {
		fmt.Println("Test line 3 failed")
		t.Errorf("got\n %v, want %v", out[64:95], expected[2])
		return
	} else {
		fmt.Println(" line 3 Test passed successfully")
	}
	if !reflect.DeepEqual(out[96:127], expected[3]) {
		fmt.Println("Test line 4 failed")
		t.Errorf("got\n %v, want %v", out[96:127], expected[3])
		return
	} else {
		fmt.Println(" line 4 Test passed successfully")
	}
	if !reflect.DeepEqual(out[128:159], expected[4]) {
		fmt.Println("Test line 5 failed")
		t.Errorf("got\n %v, want %v", out[128:159], expected[4])
		return
	} else {
		fmt.Println(" line 5 Test passed successfully")
	}
	if !reflect.DeepEqual(out[160:191], expected[5]) {
		fmt.Println("Test line 6 failed")
		t.Errorf("got\n %v, want %v", out[160:191], expected[5])
		return
	} else {
		fmt.Println(" line 6 Test passed successfully")
	}
	if !reflect.DeepEqual(out[192:223], expected[6]) {
		fmt.Println("Test line 7 failed")
		t.Errorf("got\n %v, want %v", out[192:223], expected[6])
		return
	} else {
		fmt.Println(" line 7 Test passed successfully")
	}
	if !reflect.DeepEqual(out[224:255], expected[7]) {
		fmt.Println("Test  line 8 failed")
		t.Errorf("got\n %v, want %v", out[224:255], expected[7])
		return
	} else {
		fmt.Println(" line 8 Test passed successfully")
	}
}
