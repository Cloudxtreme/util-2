// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:        2014-12-22
// Last modification: 2014-

// +build darwin dragonfly freebsd linux netbsd openbsd

// Os package have some helpers os functions.
package os

import (
	"github.com/fcavani/e"
	"syscall"
)

// DiskSpace return the total and free disk space.
func DiskSpace(path string) (total, free int, err error) {
	s := syscall.Statfs_t{}
	err = syscall.Statfs(path, &s)
	if err != nil {
		return 0, 0, e.New(err)
	}
	total = int(s.Bsize) * int(s.Blocks)
	free = int(s.Bsize) * int(s.Bavail)
	return
}
