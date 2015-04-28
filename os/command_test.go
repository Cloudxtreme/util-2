// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"reflect"
	"testing"

	"github.com/fcavani/e"
)

type testcommandstruct struct {
	raw    string
	parsed *Command
}

var tests []testcommandstruct = []testcommandstruct{
	{"init [3]", &Command{Init, "init [3]", "", "init", ""}},
	{"[kworker/0:0H]", &Command{Kernel, "[kworker/0:0H]", "", "[kworker/0:0H]", ""}},
	{"sshd: felipe@pts/1", &Command{Ssh, "sshd: felipe@pts/1", "", "ssh", ""}},
	{"/usr/sbin/cron", &Command{UserProcess, "/usr/sbin/cron", "/usr/sbin", "cron", ""}},
	{"/usr/libexec/sandboxd -n PluginProcess -n", &Command{UserProcess, "/usr/libexec/sandboxd -n PluginProcess -n", "/usr/libexec", "sandboxd", "-n PluginProcess -n"}},
}

func TestNewCommand(t *testing.T) {
	for i, test := range tests {
		cmd, err := NewCommand(test.raw)
		if err != nil {
			t.Fatal(e.Trace(e.Forward(err)))
		}
		if !reflect.DeepEqual(cmd, test.parsed) {
			t.Fatalf("not equal %v %#v", i, cmd)
		}
		t.Log(cmd)
	}
}
