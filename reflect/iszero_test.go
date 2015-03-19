// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import (
	"reflect"
	"testing"
)

type teststuct struct {
	val    reflect.Value
	iszero bool
}

var tests []teststuct = []teststuct{
	{reflect.Zero(reflect.TypeOf("foo")), true},
	{reflect.ValueOf("str"), false},
	{reflect.ValueOf(""), true},
	{reflect.Value{}, false},
}

func TestIsZero(t *testing.T) {
	for i, test := range tests {
		if IsZero(test.val) != test.iszero {
			t.Fatal("not equal", i)
		}
	}
}
