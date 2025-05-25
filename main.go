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
	// Input            []io.Reader
	Output           io.Writer
	FileConfig       []WcConfig
	TotalCounter     WordCounter
	BytesFlag        bool
	LinesFlag        bool
	WordsFlag        bool
	CharsFlag        bool
	FileNameProvided bool
}

type WcConfig struct {
	FileName string
	Counter  WordCounter
	Input    io.Reader
}

type WordCounter struct {
	Lines int
	Words int
	Chars int
	Bytes int
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

	flag.BoolVar(&cmd.BytesFlag, "c", false, "count bytes")
	flag.BoolVar(&cmd.LinesFlag, "l", false, "count lin1ss")
	flag.BoolVar(&cmd.WordsFlag, "w", false, "count words")
	flag.BoolVar(&cmd.CharsFlag, "m", false, "count chars")

	flag.Parse()
	args := flag.Args()

	setDefaultFlags(&cmd)

	switch {
	case len(args) == 0:
		cmd.FileNameProvided = false
		cmd.FileConfig = append(cmd.FileConfig, WcConfig{
			Input: os.Stdin,
		})
	case len(args) > 0:
		for _, a := range args {
			file, err := os.Open(a)
			if err != nil {
				return cmd, fmt.Errorf("couldn't open file %v, error: %v", a, err)
			}
			cmd.FileConfig = append(cmd.FileConfig, WcConfig{
				FileName: file.Name(),
				Input:    file,
			})
		}
		// defer file.Close()
		cmd.FileNameProvided = true
	default:
		return cmd, fmt.Errorf("wrong amount of args")
	}
	return cmd, nil
}

func (cmd *Command) Run() {
	for _, input := range cmd.FileConfig {
		reader := bufio.NewReader(input.Input)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			if cmd.LinesFlag {
				input.Counter.Lines++
			}
			if cmd.WordsFlag {
				input.Counter.Words += len(strings.Fields(line))
			}
			if cmd.BytesFlag {
				input.Counter.Bytes += len(line)
			}
			if cmd.CharsFlag {
				input.Counter.Chars += utf8.RuneCountInString(line)
			}
		}
		printResult(input.Counter, *cmd, input.FileName)

		if len(cmd.FileConfig) > 1 {
			cmd.addCountToTotal(input.Counter)
		}
	}
	if len(cmd.FileConfig) > 1 {
		printResult(cmd.TotalCounter, *cmd, "total")
	}
}

func printResult(counter WordCounter, cmd Command, fileName string) {
	if cmd.LinesFlag {
		fmt.Fprintf(cmd.Output, "%8d", counter.Lines)
	}
	if cmd.WordsFlag {
		fmt.Fprintf(cmd.Output, "%8d", counter.Words)
	}
	if cmd.BytesFlag {
		fmt.Fprintf(cmd.Output, "%8d", counter.Bytes)
	}
	if cmd.CharsFlag {
		fmt.Fprintf(cmd.Output, "%8d", counter.Chars)
	}
	if cmd.FileNameProvided {
		fmt.Fprintf(cmd.Output, " %s", fileName)
	}
	fmt.Fprintln(cmd.Output)
}

// TODO: multiple file input support

// If no flags provided enable standard wc options lines, words and bytes
func setDefaultFlags(cmd *Command) {
	if !cmd.BytesFlag && !cmd.LinesFlag && !cmd.WordsFlag && !cmd.CharsFlag {
		cmd.LinesFlag, cmd.WordsFlag, cmd.BytesFlag = true, true, true
	}
}

func (cmd *Command) addCountToTotal(input WordCounter) {
	if cmd.LinesFlag {
		cmd.TotalCounter.Lines += input.Lines
	}
	if cmd.WordsFlag {
		cmd.TotalCounter.Words += input.Words
	}
	if cmd.BytesFlag {
		cmd.TotalCounter.Bytes += input.Bytes
	}
	if cmd.CharsFlag {
		cmd.TotalCounter.Chars += input.Chars
	}
}
