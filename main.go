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
	cmd, err := loadCommand()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading command:", err)
	}

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

	setDefaultFlags(&cmd.Flag)

	switch {
	case len(args) == 0:
		cmd.Flag.FileNameProvided = false
		cmd.Input = os.Stdin
	case len(args) == 1:
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// defer file.Close()
		cmd.Input = file
		cmd.Flag.FileNameProvided = true
		cmd.Flag.FileName = file.Name()
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
		if cmd.Flag.Lines {
			cmd.Count.linesTotal++
		}
		if cmd.Flag.Words {
			cmd.Count.wordsTotal += len(strings.Fields(line))
		}
		if cmd.Flag.Bytes {
			cmd.Count.bytesTotal += len(line)
		}
		if cmd.Flag.Chars {
			cmd.Count.charsTotal += utf8.RuneCountInString(line)
		}
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
	fmt.Fprintln(w)
}

// TODO: multiple file input support

// If no flags provided enable standard wc options lines, words and bytes
func setDefaultFlags(flag *WcFlag) {
	if !flag.Bytes && !flag.Lines && !flag.Words && !flag.Chars {
		flag.Lines, flag.Words, flag.Bytes = true, true, true
	}
}
