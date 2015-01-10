package logger

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestLogWrite(t *testing.T) {

	text1 := "One"
	text2 := "Two"
	expected := text1 + " " + text2

	// Create temp file
	file, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		t.Error("Cannot create temp file")
	}
	defer os.Remove(file.Name())

	// Set log file location to temp file
	SetLogFileLocation(file.Name())

	// Don't show timestamp at all
	SetTimeFormat("")

	I(text1, text2)

	// Open temp file
	file, err = os.Open(file.Name())
	if err != nil {
		t.Error("Cannot open temp file", err)
	}
	defer file.Close()

	// Read temp file row
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		text := scanner.Text()
		if !strings.Contains(text, expected) {
			t.Error("Expected:", expected, ", got:", text)
		}
	}
}
