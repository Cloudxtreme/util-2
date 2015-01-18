// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2014-12-17
// Last modification:	2014-x

// +build darwin

// Os package have some helpers os functions.
package os

import (
	"github.com/fcavani/e"
	"os"
	"syscall"
	"time"
)

// StatTimes return the file's access, modification and creation times.
func StatTimes(name string) (atime, mtime, ctime time.Time, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		err = e.New(err)
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec))
	ctime = time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec))
	return
}
