package lib

import (
	"fmt"
	"testing"
)

func TestToHex0(t *testing.T) {
	tests := []int{0, 1, 123456789}
	results := []string{"0", "1", "75BCD15"}
	for i, test := range tests {
		result, err := ToHex(test)
		if err != nil {
			t.Errorf(fmt.Sprintf("Hex error[%s]\n", err))
		}
		expected := results[i]
		if result != expected {
			t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
		}
	}
}

func TestToAscii850(t *testing.T) {
	tests := []int{0, 1, 871, 931, 1000}
	results := []string{"0", "1", "AL", "A{", "B%"}
	for i, test := range tests {
		result, err := ToAscii85(test)
		if err != nil {
			t.Errorf(fmt.Sprintf("Hex error[%s]\n", err))
		}
		expected := results[i]
		if result != expected {
			t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
		}
	}
}
