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
	"strings"
	"unicode/utf8"
)

type Command struct {
	Input  io.Reader
	Output io.Writer
	Flag   WcFlag
	Count  Count
}

type WcFlag struct {
	Bytes            bool
	Lines            bool
	Words            bool
	Chars            bool
	FileNameProvided bool
	FileName         string
}

type Count struct {
	linesTotal int
	wordsTotal int
	charsTotal int
	bytesTotal int
}

func main() {
	cmd, _ := loadCommand()

	cmd.Run()
}

func loadCommand() (Command, error) {
	// var err error
	cmd := Command{
		Output: os.Stdout,
	}

	flag.BoolVar(&cmd.Flag.Bytes, "c", false, "count bytes")
	flag.BoolVar(&cmd.Flag.Lines, "l", false, "count lines")
	flag.BoolVar(&cmd.Flag.Words, "w", false, "count words")
	flag.BoolVar(&cmd.Flag.Chars, "m", false, "count chars")

	flag.Parse()
	args := flag.Args()

	// If no flags provided enable standard wc options lines, words and bytes
	if !cmd.Flag.Bytes && !cmd.Flag.Lines && !cmd.Flag.Words && !cmd.Flag.Chars {
		cmd.Flag.Lines, cmd.Flag.Words, cmd.Flag.Bytes = true, true, true
	}

	switch {
	case len(args) == 0:
		cmd.Flag.FileNameProvided = false
		cmd.Input = os.Stdin
	case len(args) == 1:
		filePath, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cmd.Input = filePath
		cmd.Flag.FileNameProvided = true
		cmd.Flag.FileName = filePath.Name()
	}
	return cmd, nil
}

func (cmd *Command) Run() {
	reader := bufio.NewReader(cmd.Input)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		cmd.Count.linesTotal++

		cmd.Count.wordsTotal += len(strings.Fields(line))

		cmd.Count.bytesTotal += len(line)

		cmd.Count.charsTotal += utf8.RuneCountInString(line)
	}
	printResult(cmd.Count, cmd.Flag, cmd.Output)
}

func printResult(count Count, flag WcFlag, w io.Writer) {
	if flag.Lines {
		fmt.Fprintf(w, "%8d", count.linesTotal)
	}
	if flag.Words {
		fmt.Fprintf(w, "%8d", count.wordsTotal)
	}
	if flag.Bytes {
		fmt.Fprintf(w, "%8d", count.bytesTotal)
	}
	if flag.Chars {
		fmt.Fprintf(w, "%8d", count.charsTotal)
	}
	if flag.FileNameProvided {
		fmt.Fprintf(w, " %s", flag.FileName)
	}
	// fmt.Fprintln(w)
}

/*

type wcCommand struct {
	wcObjects        []wcObj
	bytes            bool
	lines            bool
	words            bool
	chars            bool
	multipleFiles    bool
	fileNameProvided bool
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

*/
