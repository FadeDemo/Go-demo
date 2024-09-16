// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 308.
//!+

// Echo prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdout // modified during testing

func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	_, err := fmt.Fprint(out, strings.Join(args, sep))
	if err != nil {
		return err
	}
	if newline {
		_, err := fmt.Fprintln(out)
		if err != nil {
			return err
		}
	}
	return nil
}

//!-
