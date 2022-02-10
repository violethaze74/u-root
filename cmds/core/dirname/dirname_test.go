// Copyright 2016 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"testing"
)

func TestDirName(t *testing.T) {
	// Table-driven testing
	for _, tt := range []struct {
		name string
		args []string
		out  string
	}{
		{
			name: "Usage",
			args: []string{},
		},
		{
			name: "EmptyArgs",
			args: []string{""},
			out:  ".\n",
		},
		{
			name: "/this/that",
			args: []string{"/this/that"},
			out:  "/this\n",
		},
		{
			name: "/this/that_/other",
			args: []string{"/this/that", "/other"},
			out:  "/this\n/\n",
		},
		{
			name: "/this/that_/other thing/space",
			args: []string{"/this/that", "/other thing/space"},
			out:  "/this\n/other thing\n",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			if err := runDirname(&buf, tt.args); err != nil {
				t.Errorf("runDirname(%v, &buf)=%q, want nil", tt.args, err)
			}
		})
	}
}
