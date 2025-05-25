package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestWcUnit(t *testing.T) {
	t.Run("normal wordcount cmd", func(t *testing.T) {
		var buf bytes.Buffer
		cmd := Command{
			Input: strings.NewReader(`Hello testing
yes yes cool
ok goodbye now
`),
			Output: &buf,
			Flag: WcFlag{
				Words: true,
				Lines: true,
				Bytes: true,
			},
		}

		cmd.Run()

		got := buf.String()
		want := "       3       8      42"

		assertEqual(t, string(got), string(want))
	})

	t.Run("charcount from file with emojis, with filename", func(t *testing.T) {
		var buf bytes.Buffer
		cmd := Command{
			Input: strings.NewReader(`Hello testing😊
yes yes cool
ok goodbye 🌟now
`),
			Output: &buf,
			Flag: WcFlag{
				Chars:            true,
				FileNameProvided: true,
				FileName:         "faketest.txt",
			},
		}

		cmd.Run()

		got := buf.String()
		want := "      44 faketest.txt"

		assertEqual(t, string(got), string(want))
	})
}

// Old integration tests
var testFiles = getTestFiles("./testdata/")

func getTestFiles(testFolder string) []string {
	var res []string

	files, err := os.ReadDir(testFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		res = append(res, testFolder+file.Name())
	}
	return res
}

func TestWcWithoutFlag(t *testing.T) {
	for _, testFile := range testFiles {
		cmd := exec.Command("./ccwc", testFile)
		got, err := cmd.Output()
		if err != nil {
			t.Fatalf("Command %s failed with error: %v", cmd.String(), err)
		}

		unixCmd := exec.Command("wc", testFile)
		want, err := unixCmd.Output()
		if err != nil {
			t.Fatalf("Command %s failed with error: %v", unixCmd.String(), err)
		}

		assertEqual(t, string(got), string(want))
	}
}

func TestWcWithLinesFlag(t *testing.T) {
	for _, testFile := range testFiles {
		cmd := exec.Command("./ccwc", "-l", testFile)
		got, err := cmd.Output()
		if err != nil {
			t.Fatalf("Command %s failed with error: %v", cmd.String(), err)
		}

		unixCmd := exec.Command("wc", "-l", testFile)
		want, err := unixCmd.Output()
		if err != nil {
			t.Fatalf("Command %s failed with error: %v", unixCmd.String(), err)
		}

		assertEqual(t, string(got), string(want))
	}
}

func TestWcWithBytesFlag(t *testing.T) {
	for _, testFile := range testFiles {
		cmd := exec.Command("./ccwc", "-c", testFile)
		got, err := cmd.Output()
		if err != nil {
			t.Fatalf("Command %s failed with error: %v", cmd.String(), err)
		}

		unixCmd := exec.Command("wc", "-c", testFile)
		want, err := unixCmd.Output()
		if err != nil {
			t.Fatalf("Command %s failed with error: %v", unixCmd.String(), err)
		}

		assertEqual(t, string(got), string(want))
	}
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
