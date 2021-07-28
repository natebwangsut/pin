package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

// Verbose mode print extra log for more information
var Verbose bool

var pinTable = map[rune]rune{
	'a': '2',
	'b': '2',
	'c': '2',
	'd': '3',
	'e': '3',
	'f': '3',
	'g': '4',
	'h': '4',
	'i': '4',
	'j': '5',
	'k': '5',
	'l': '5',
	'm': '6',
	'n': '6',
	'o': '6',
	'p': '7',
	'q': '7',
	'r': '7',
	's': '7',
	't': '8',
	'u': '8',
	'v': '8',
	'w': '9',
	'x': '9',
	'y': '9',
	'z': '9',
}

func pin(r rune) rune {
	return pinTable[unicode.ToLower(r)]
}

func pinIgnoreChars(r rune) rune {
	if p, ok := pinTable[unicode.ToLower(r)]; ok {
		return p
	}
	return r
}

func main() {
	pinCmd := &cobra.Command{
		Use:   "pin <string>",
		Short: "Change the received string to PIN format (e.g. hi -> 44)",
		Long: `Convert strings to PIN format. For example:

pin "hello"             -> "43556"
pin "hello-world"       -> "4355696753"
pin "HELLO-WORLD"       -> "4355696753"
pin -i "hello-world"    -> "43556-96753"`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			if Verbose {
				fmt.Println("source: ", source)
			}

			isIgnoreChars, _ := cmd.Flags().GetBool("ignore")
			if isIgnoreChars {
				fmt.Println(strings.Map(pinIgnoreChars, source))
				return
			}

			fmt.Println(strings.Map(pin, source))
		},
	}

	pinCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	pinCmd.Flags().BoolP("ignore", "i", false, "ignore non convertable characters")
	_ = pinCmd.Execute()
}
