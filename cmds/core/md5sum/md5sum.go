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

func md5Sum(w io.Writer, r io.Reader, args ...string) (err error) {
	if *help {
		util.Usage(usage)
		flag.Usage()
		return nil
	}

	fileName := ""
	if len(args) >= 1 {
		fileName = args[0]
	}
	if fileName == "" {
		input, err := io.ReadAll(r)
		if err != nil {
			fmt.Println("Error getting input.")
			return err
		}
		fmt.Fprintf(w, "%x\n", md5.Sum(input))
	} else {
		fileDesc, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer fileDesc.Close()
		h := md5.New()
		if _, err := io.Copy(h, fileDesc); err != nil {
			return err
		}
		fmt.Fprintf(w, "%x %s\n", h.Sum(nil), fileName)
	}
	return nil
}

func main() {
	flag.Parse()
	if err := md5Sum(os.Stdout, os.Stdin, flag.Args()...); err != nil {
		log.Fatal(err)
	}
}
