// Copyright 2015-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !plan9 && !windows

package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestRunDF(t *testing.T) {
	for _, tt := range []struct {
		name    string
		args    []string
		fargs   flags
		wantErr error
	}{
		{
			name: "Usage",
			args: []string{"", ""},
		},
		{
			name: "NoArgs-NoFlags",
		},
		{
			name: "NoArgs-M-Flag",
			fargs: flags{
				m: true,
			},
		},
		{
			name: "NoArgs-K-Flag",
			fargs: flags{
				k: true,
			},
		},
		{
			name: "NoArgs-KM-Flag",
			fargs: flags{
				k: true,
				m: true,
			},
			wantErr: errKMExclusiv,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			if err := df(tt.args, tt.fargs, &buf); !errors.Is(err, tt.wantErr) {
				t.Errorf("df(%v, %v, buf)=%q, want %q", tt.args, tt.fargs, err, tt.wantErr)
			}
		})
	}
}
