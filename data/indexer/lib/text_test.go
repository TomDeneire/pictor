package lib

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplitLine0(t *testing.T) {
	result, err := SplitIntoIndexWords("abc def")
	if err != nil {
		t.Errorf(fmt.Sprintf("Split error[%s]\n", err))
	}
	expected := []string{"abc", "def"}
	if strings.Join(result, "") != strings.Join(expected, "") {
		t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
	}
}

func TestSplitLine1(t *testing.T) {
	result, err := SplitIntoIndexWords("abc, def")
	if err != nil {
		t.Errorf(fmt.Sprintf("Split error[%s]\n", err))
	}
	expected := []string{"abc", "def"}
	if strings.Join(result, "") != strings.Join(expected, "") {
		t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
	}
}

func TestSplitLine2(t *testing.T) {
	result, err := SplitIntoIndexWords("abc; def, '' !§^#")
	if err != nil {
		t.Errorf(fmt.Sprintf("Split error[%s]\n", err))
	}
	expected := []string{"abc", "def"}
	if strings.Join(result, "") != strings.Join(expected, "") {
		t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
	}
}

func TestSplitLine3(t *testing.T) {
	result, err := SplitIntoIndexWords("abc, def, ''")
	if err != nil {
		t.Errorf(fmt.Sprintf("Split error[%s]\n", err))
	}
	expected := []string{"abc", "def"}
	if strings.Join(result, "") != strings.Join(expected, "") {
		t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
	}
}

func TestSplitLine4(t *testing.T) {
	result, err := SplitIntoIndexWords("abc, def, aan, de, voor, the, of")
	if err != nil {
		t.Errorf(fmt.Sprintf("Split error[%s]\n", err))
	}
	expected := []string{"abc", "def"}
	if strings.Join(result, "") != strings.Join(expected, "") {
		t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
	}
}

func TestInterfaceToString1(t *testing.T) {
	result, err := InterfaceToString("abc")
	if err != nil {
		t.Errorf(fmt.Sprintf("Split error[%s]\n", err))
	}
	expected := "abc"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
	}
}

func TestInterfaceToString2(t *testing.T) {
	result, err := InterfaceToString([]string{"abc", "def"})
	if err != nil {
		t.Errorf(fmt.Sprintf("Split error[%s]\n", err))
	}
	expected := "abc"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
	}
}

func TestInterfaceToString3(t *testing.T) {
	testMap := make(map[string]interface{})
	test := make([]interface{}, 2)
	test[0] = "abc"
	test[1] = "def"
	testMap["en"] = test
	testMap["nl"] = test
	result, err := InterfaceToString(testMap)
	if err != nil {
		t.Errorf(fmt.Sprintf("Split error[%s]\n", err))
	}
	expected := "abc"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
	}
}

func TestInterfaceToString4(t *testing.T) {
	testMap := make(map[string]string)
	testMap["en"] = "abc"
	result, err := InterfaceToString(testMap)
	if err != nil {
		t.Errorf(fmt.Sprintf("Split error[%s]\n", err))
	}
	expected := "abc"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n%s\nExpected: \n%s\n", result, expected))
	}
}
