// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2014-12-04
// Last modification:	2014-x

// Functions to manipulate slices of bytes.
package bytes

import (
	"testing"
)

type testTailStruct struct {
	msg string
	head string
	mid string
	tail string
	fail bool
}

var testTail []testTailStruct = []testTailStruct {
	{"a|b|c", "a", "b","c", false},
	{"aaaa|bbbb|cccc", "aaaa", "bbbb", "cccc", false},
	{"aaaa|bbbb|", "aaaa", "bbbb", "", true},
	{"aaaa|bbbb", "aaaa", "bbbb", "", true},
	{"aaaa", "aaaa", "", "", true},
	{"a|a|aa|bbbb|cccc", "a|a|aa", "bbbb", "cccc", false},
}

func TestTail(t *testing.T) {
	for i, test := range testTail {
		headmid, tail_ := HeadTail([]byte(test.msg), "|")
		head, mid := HeadTail(headmid, "|")
		t.Logf("\"%v\", \"%v\", \"%v\"", string(head), string(mid), string(tail_))
		if (string(head) == "" || string(mid) == "" || string(tail_) == "") && !test.fail {
			t.Fatal("Failed:", i, string(head), string(mid), string(tail_))
		}
		if (string(head) != test.head || string(mid) != test.mid || string(tail_) != test.tail) && !test.fail {
			t.Fatal("Not spected:", i, string(head), string(mid), string(tail_))
		}
	}
}
