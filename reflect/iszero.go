// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import (
	"reflect"
)

func IsZero(val reflect.Value) bool {
	// from: http://stackoverflow.com/questions/13901819/quick-way-to-detect-empty-values-via-reflection-in-go
	if !val.IsValid() {
		return false
	}
	return val.Interface() == reflect.Zero(val.Type()).Interface()
}
