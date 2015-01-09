// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Start date:		2014-12-16
// Last modification:	2014-x

package os

import (
	"os"
	"projects/e"
	"syscall"
)

func StatUidGid(name string) (uid, gid int, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	stat := fi.Sys().(*syscall.Stat_t)
	return int(stat.Uid), int(stat.Gid), nil
}

func StatMode(name string) (os.FileMode, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return 0, e.New(err)
	}
	return fi.Mode(), nil
}

func Inode(name string) (uint64, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return 0, e.New(err)
	}
	stat := fi.Sys().(*syscall.Stat_t)
	return stat.Ino, nil
}
