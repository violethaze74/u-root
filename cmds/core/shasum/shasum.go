// Copyright 2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/pflag"
)

var (
	algorithm = pflag.IntP("algorithm", "a", 1, "SHA algorithm, valid args are 1 and 256")
	help      = pflag.BoolP("help", "h", false, "Show this help and exit")
)

var usage = "Usage:\nshasum -a <algorithm> <File Name>"

func helpPrinter() {
	fmt.Println(usage)
	pflag.PrintDefaults()
}

func getInput(r io.Reader, fileName string) (input []byte, err error) {
	if fileName != "" {
		return os.ReadFile(fileName)
	}
	return io.ReadAll(r)
}

//
// shaPrinter prints sha1/sha256 of given data. The
// value of algorithm is expected to be 1 for SHA1
// and 256 for SHA256
//
func shaPrinter(w io.Writer, data []byte) error {
	if *algorithm == 256 {
		fmt.Fprintf(w, "%x", sha256.Sum256(data))
	} else if *algorithm == 1 {
		fmt.Fprintf(w, "%x", sha1.Sum(data))
	} else {
		return fmt.Errorf("invalid algorithm")
	}
	return nil
}

func shasum(w io.Writer, r io.Reader, args ...string) error {
	cliArgs := ""

	if *help {
		helpPrinter()
		return nil
	}

	if len(args) == 1 {
		cliArgs = args[0]
	}
	input, err := getInput(r, cliArgs)
	if err != nil {
		return fmt.Errorf("error getting input")
	}
	if err := shaPrinter(w, input); err != nil {
		return err
	}
	if cliArgs == "" {
		fmt.Fprintf(w, " -\n")
	} else {
		fmt.Fprintf(w, " %s\n", cliArgs)
	}
	return nil
}

func main() {
	pflag.Parse()
	if err := shasum(os.Stdout, os.Stdin, pflag.Args()...); err != nil {
		log.Fatal(err)
	}
}
