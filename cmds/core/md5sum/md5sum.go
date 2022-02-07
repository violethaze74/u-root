// Copyright 2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// md5sum prints an md5 hash generated from file contents.
package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/u-root/u-root/pkg/uroot/util"
)

var (
	help = flag.Bool("h", false, "Show this help and exit")
)

var usage = "md5sum: md5sum <File Name>"

func calculateMd5Sum(w io.Writer, fileName string, data []byte) (string, error) {
	if len(data) > 0 {
		_, err := fmt.Fprintf(w, "%x", md5.Sum(data))
		if err != nil {
			return "", err
		}
		return "", nil
	}

	fileDesc, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer fileDesc.Close()

	md5Generator := md5.New()
	if _, err := io.Copy(md5Generator, fileDesc); err != nil {
		return "", err
	}

	_, err = fmt.Fprintf(w, "%x", md5Generator.Sum(nil))
	if err != nil {
		return "", err
	}
	return "", nil
}

func md5Sum(w io.Writer, r io.Reader, args ...string) error {
	var (
		input []byte
		err   error
	)

	if *help {
		util.Usage(usage)
		flag.Usage()
		return nil
	}

	cliArgs := ""
	if len(args) >= 1 {
		cliArgs = args[0]
	}
	if cliArgs == "" {
		input, err = io.ReadAll(r)
		if err != nil {
			fmt.Println("Error getting input.")
			return err
		}
	}
	_, err = calculateMd5Sum(w, cliArgs, input)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, " %s\n", cliArgs)
	return nil
}

func main() {
	flag.Parse()
	if err := md5Sum(os.Stdout, os.Stdin, flag.Args()...); err != nil {
		log.Fatal(err)
	}
}
