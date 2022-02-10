// Copyright 2016-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// dirname prints out the directory name of one or more args.
// If no arg is given it returns an error and prints a message which,
// per the man page, is incorrect, but per the standard, is correct.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func runDirname(w io.Writer, args []string) error {
	if len(args) < 1 {
		fmt.Fprintf(w, "%s\n", "dirname: missing operand")
		return nil
	}
	for _, n := range args {
		if _, err := fmt.Fprintf(w, "%s\n", filepath.Dir(n)); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := runDirname(os.Stdout, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
