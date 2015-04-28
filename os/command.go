// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"path/filepath"
	"strings"

	"github.com/fcavani/e"
)

// Common types of programs running in a linux machine.
type CommandType uint8

const (
	Undefined CommandType = iota
	UserProcess
	Kernel
	Init
	Ssh
)

func (c CommandType) String() string {
	switch c {
	case Undefined:
		return "undefined"
	case UserProcess:
		return "user process"
	case Kernel:
		return "kernel"
	case Init:
		return "init"
	case Ssh:
		return "ssh"
	default:
		return "invalid"
	}
}

// Command struct represents a command find in ps aux.
type Command struct {
	Type CommandType
	// Raw string find in ps aux
	Raw string
	// Program path
	Path string
	// Program file name
	File string
	// Program arguments
	Args string
}

func (c Command) String() string {
	return c.Raw
}

// NewCommand parses the contents of the column command in the ps aux output.
func NewCommand(raw string) (*Command, error) {
	if len(raw) == 0 {
		return nil, e.New("empty command")
	}
	switch {
	case raw[0] == '[':
		return &Command{
			Type: Kernel,
			Raw:  raw,
			File: raw,
		}, nil
	case strings.HasPrefix(raw, "init"):
		return &Command{
			Type: Init,
			Raw:  raw,
			File: "init",
		}, nil
	case strings.HasPrefix(raw, "ssh"):
		return &Command{
			Type: Ssh,
			Raw:  raw,
			File: "ssh",
		}, nil
	default:
		cmd := &Command{
			Type: UserProcess,
			Raw:  raw,
		}
		i := strings.Index(raw, " ")
		if i < 0 {
			cmd.File = filepath.Base(raw)
			cmd.Path = filepath.Dir(raw)
			return cmd, nil
		}
		path := raw[:i]
		cmd.File = filepath.Base(path)
		cmd.Path = filepath.Dir(path)
		if i+1 >= len(raw) {
			return cmd, nil
		}
		cmd.Args = raw[i+1:]
		return cmd, nil
	}
}
