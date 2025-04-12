/*
Copyright © 2025 boxy-pug
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var bytesCount bool
var lines bool
var words bool
var chars bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "word, line, character, and byte count",
	Long: `The wc utility displays the number of lines, words, and bytes contained in each input file, or
     standard input (if no file is specified) to the standard output.  A line is defined as a string
     of characters delimited by a ⟨newline⟩ character.  Characters beyond the final ⟨newline⟩
     character will not be included in the line count.

     A word is defined as a string of characters delimited by white space characters.  White space
     characters are the set of characters for which the iswspace(3) function returns true.  If more
     than one input file is specified, a line of cumulative counts for all the files is displayed on
     a separate line after the output for the last file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]

		if !bytesCount && !lines && !words && !chars {
			bytesCount, lines, words = true, true, true
		}

		byteList, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("couldn't open file as bytes: %v", err)
			os.Exit(1)
		}

		if bytesCount {
			fmt.Printf("%d ", len(byteList))
		}

		if lines {
			lineCount := getLineCount(byteList)
			fmt.Printf("%d ", lineCount)

		}

		if words {
			wordList := strings.Fields(string(byteList))
			fmt.Printf("%d ", len(wordList))

		}

		if chars {
			runeList := []rune(string(byteList))
			fmt.Printf("%d ", len(runeList))
		}

		fmt.Printf("%s\n", inputFile)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVarP(&bytesCount, "bytes", "c", false, "use flag to display byte count")
	rootCmd.PersistentFlags().BoolVarP(&lines, "lines", "l", false, "use flag to display line count")
	rootCmd.PersistentFlags().BoolVarP(&words, "words", "w", false, "use flag to display word count")
	rootCmd.PersistentFlags().BoolVarP(&chars, "chars", "m", false, "use flag to display char count")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func openFile(input string) (*os.File, error) {

	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getWordCount() {

}

func getLineCount(bt []byte) int {
	i := 0
	lineCount := 0

	for i < len(bt) {
		ind := bytes.Index(bt[i:], []byte("\n"))
		if ind == -1 {
			break
		}
		lineCount += 1
		i += ind + 1
	}

	return lineCount
}
