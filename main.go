/*
Copyright © 2025 boxy-pug
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

type wcCommand struct {
	wcObjects        []wcObj
	bytes            bool
	lines            bool
	words            bool
	chars            bool
	multipleFiles    bool
	fileNameProvided bool
	linesTotal       int
	wordsTotal       int
	charsTotal       int
	bytesTotal       int
}

type wcObj struct {
	file  *os.File
	bytes int
	lines int
	words int
	chars int
}

func main() {
	wc := wcCommand{}

	flag.BoolVar(&wc.bytes, "c", false, "count bytes")
	flag.BoolVar(&wc.lines, "l", false, "count lines")
	flag.BoolVar(&wc.words, "w", false, "count words")
	flag.BoolVar(&wc.chars, "m", false, "count chars")

	flag.Parse()

	wc.openFile()
	wc.getWordCount()

	if !wc.bytes && !wc.chars && !wc.lines && !wc.words {
		wc.chars, wc.lines, wc.words = true, true, true
	}

	wc.printToConsole()
}

func (wc *wcCommand) getWordCount() {
	for i := range wc.wcObjects {
		wcObj := &wc.wcObjects[i]
		reader := bufio.NewReader(wcObj.file)
		inWord := false

		for {
			r, b, err := reader.ReadRune()
			if err == io.EOF {
				break
			}

			wcObj.chars++
			wcObj.bytes += b

			if r == '\n' {
				wcObj.lines++
			}

			if unicode.IsSpace(r) {
				inWord = false
			} else if !inWord {
				wcObj.words++
				inWord = true
			}
		}
	}
}

func (wc *wcCommand) printToConsole() {
	for _, wcobj := range wc.wcObjects {
		if wc.lines {
			fmt.Printf("%8d", wcobj.lines)
			wc.linesTotal += wcobj.lines
		}
		if wc.words {
			fmt.Printf("%8d", wcobj.words)
			wc.wordsTotal += wcobj.lines
		}
		if wc.bytes {
			fmt.Printf("%8d", wcobj.bytes)
			wc.bytesTotal += wcobj.bytes
		}
		if wc.chars {
			fmt.Printf("%8d", wcobj.chars)
			wc.charsTotal += wcobj.chars
		}
		if wc.fileNameProvided {
			fmt.Printf(" %s\n", wcobj.file.Name())
		}
	}
	if wc.multipleFiles {
		wc.printTotalLine()
	}
}

func (wc *wcCommand) printTotalLine() {
	if wc.lines {
		fmt.Printf("%8d", wc.linesTotal)
	}
	if wc.words {
		fmt.Printf("%8d", wc.wordsTotal)
	}
	if wc.bytes {
		fmt.Printf("%8d", wc.bytesTotal)
	}
	if wc.chars {
		fmt.Printf("%8d", wc.charsTotal)
	}
	fmt.Println(" total")

}

func (wc *wcCommand) openFile() {
	filePaths := flag.Args()

	switch len(filePaths) {
	case 0:
		wc.wcObjects = append(wc.wcObjects, wcObj{file: os.Stdin})
		wc.fileNameProvided = false
	default:
		for _, filePath := range filePaths {
			openFile, err := os.Open(filePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			wc.wcObjects = append(wc.wcObjects, wcObj{file: openFile})
			wc.fileNameProvided = true
		}
	}

	if len(wc.wcObjects) > 1 {
		wc.multipleFiles = true
	}
}
