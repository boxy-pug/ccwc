package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestWc(t *testing.T) {
	var buf bytes.Buffer

	t.Run("normal wordcount cmd", func(t *testing.T) {
		cmd := Command{
			Input: strings.NewReader(`Hello testing
yes yes cool
ok goodbye now`),

			Output: &buf,
			Flag:   WcFlag{},
		}

		cmd.Run()

		got := buf.String()
		want := "       3       8      42"

		if got != want {
			t.Errorf("got: %q want: %q\n", got, want)
		}
	})
}

/*
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

		gotString := string(got)
		wantString := string(want)

		if gotString != wantString {
			t.Errorf("EXPECTED: \n%q\nGOT: \n%q\n", wantString, gotString)
		}
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

		gotString := string(got)
		wantString := string(want)

		if gotString != wantString {
			t.Errorf("EXPECTED: \n%q\nGOT: \n%q\n", wantString, gotString)
		}
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

		gotString := string(got)
		wantString := string(want)

		if gotString != wantString {
			t.Errorf("EXPECTED: \n%q\nGOT: \n%q\n", wantString, gotString)
		}
	}
}
*/
