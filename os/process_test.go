// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"os"
	"testing"
	"github.com/fcavani/e"
)

func TestProcesses(t *testing.T) {
	procs, err := Processes()
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	p, found := procs[os.Getpid()]
	if !found {
		t.Fatal("current process not found")
	}
	t.Log(p)
}