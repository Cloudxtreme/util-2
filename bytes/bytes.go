// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2014-12-04
// Last modification:	2014-x

// Functions to manipulate slices of bytes.
package bytes

import (
	"bytes"
)

func HeadTail(in []byte, sep string) (head, tail []byte) {
	sepi := bytes.LastIndex(in, []byte(sep))
	if sepi == -1 || sepi >= len(in)-1 {
		head = in
		tail = []byte{}
		return
	}
	tail = in[sepi+1:]
	head = in[:sepi]
	return
}