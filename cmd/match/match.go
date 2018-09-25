package main

import (
	"bufio"
	"fmt"
	"github.com/paulgriffiths/automata/regex"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "match: missing regular expression\n")
		os.Exit(1)
	}

	rex := regex.Compile(os.Args[1])
	if rex == nil {
		fmt.Fprintf(os.Stderr, "match: invalid regular expression\n")
		os.Exit(1)
	}

	infiles := []*os.File{}
	if len(os.Args) > 2 {
		for _, filename := range os.Args[2:] {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "match: couldn't open file %s: %v\n",
					filename, err)
				os.Exit(1)
			}
			infiles = append(infiles, f)
		}
	} else {
		infiles = append(infiles, os.Stdin)
	}

	for _, f := range infiles {
		input := bufio.NewScanner(f)
		for input.Scan() {
			if rex.Match(input.Text()) {
				fmt.Printf("%s\n", input.Text())
			}
		}
		if f != os.Stdin {
			if err := f.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "match: error closing file: %v\n",
					err)
			}
		}
	}
}
