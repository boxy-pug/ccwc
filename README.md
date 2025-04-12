# ccwc

`ccwc` is a command-line utility that replicates the functionality of the Unix `wc` command, providing counts of lines, words, characters, and bytes in a given file or standard input.

## Features

-  **Line Count**: Count the number of lines in the input.
-  **Word Count**: Count the number of words in the input.
-  **Character Count**: Count the number of characters in the input.
-  **Byte Count**: Count the number of bytes in the input.

## Usage

```bash
ccwc [flags] [file]
```

-  If no file is specified, `ccwc` reads from standard input.
-  If multiple flags are provided, `ccwc` displays counts for the specified metrics.

### Flags

-  `-l`, `--lines`: Display the line count.
-  `-w`, `--words`: Display the word count.
-  `-m`, `--chars`: Display the character count.
-  `-c`, `--bytes`: Display the byte count.

## Examples

-  Count lines, words, and bytes in a file:
  ```bash
  ccwc -lwc testdata/test.txt
  ```

-  Count characters from standard input:
  ```bash
  echo "Hello, world!" | ccwc -m
  ```

## Author

Created by boxy-pug as part of the [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc) assignment.
