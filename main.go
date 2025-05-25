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
	Input   io.Reader
	Output  io.Writer
	Config  WcConfig
	Counter WordCounter
}

type WcConfig struct {
	Bytes            bool
	Lines            bool
	Words            bool
	Chars            bool
	FileNameProvided bool
	FileName         string
}

type WordCounter struct {
	linesTotal int
	wordsTotal int
	charsTotal int
	bytesTotal int
}

func main() {
	cmd, err := loadCommand()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading command:", err)
		os.Exit(1)
	}

	cmd.Run()
}

func loadCommand() (Command, error) {
	// var err error
	cmd := Command{
		Output: os.Stdout,
	}

	flag.BoolVar(&cmd.Config.Bytes, "c", false, "count bytes")
	flag.BoolVar(&cmd.Config.Lines, "l", false, "count lin1ss")
	flag.BoolVar(&cmd.Config.Words, "w", false, "count words")
	flag.BoolVar(&cmd.Config.Chars, "m", false, "count chars")

	flag.Parse()
	args := flag.Args()

	setDefaultFlags(&cmd.Config)

	switch {
	case len(args) == 0:
		cmd.Config.FileNameProvided = false
		cmd.Input = os.Stdin
	case len(args) == 1:
		file, err := os.Open(args[0])
		if err != nil {
			return cmd, fmt.Errorf("couldn't open file %v, error: %v", args[0], err)
		}
		// defer file.Close()
		cmd.Input = file
		cmd.Config.FileNameProvided = true
		cmd.Config.FileName = file.Name()
	default:
		return cmd, fmt.Errorf("wrong amount of args")
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
		if cmd.Config.Lines {
			cmd.Counter.linesTotal++
		}
		if cmd.Config.Words {
			cmd.Counter.wordsTotal += len(strings.Fields(line))
		}
		if cmd.Config.Bytes {
			cmd.Counter.bytesTotal += len(line)
		}
		if cmd.Config.Chars {
			cmd.Counter.charsTotal += utf8.RuneCountInString(line)
		}
	}
	printResult(cmd.Counter, cmd.Config, cmd.Output)
}

func printResult(count WordCounter, flag WcConfig, w io.Writer) {
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
func setDefaultFlags(flag *WcConfig) {
	if !flag.Bytes && !flag.Lines && !flag.Words && !flag.Chars {
		flag.Lines, flag.Words, flag.Bytes = true, true, true
	}
}
