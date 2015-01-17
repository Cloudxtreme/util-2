// Copyright 2013 Felipe Alves Cavani. All rights reserved.
// Start date:		2015-01-07
// Last modification:	2015-x

package text

import (
	"testing"
)

type test struct {
	data   string
	length int
	result string
}

var tests []test = []test{
	{"qqqq", 1, "qqqq"},
	{"qqqq", 2, "qqqq"},
	{"qqqqqq", 2, "qq..."},
	{"qqqq", 10, "qqqq"},
	{"qq qq", 4, "qq qq"},
	{"ab cdefg", 4, "ab..."},
}

func TestReticence(t *testing.T) {
	for i, test := range tests {
		str := Reticence(test.data, test.length)
		t.Log(i, test.data, test.length, test.result, str)
		if str != test.result {
			t.Fatal(i, test.data, test.length, test.result, str)
		}
	}
}
