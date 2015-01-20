// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2014-12-04
// Last modification:	2014-x

package text

import (
	"unicode"
)

func FirstCaps(str string) string {
	ret := ""
	if len(str) >= 1 {
		ret = string(unicode.ToUpper(rune(str[0])))
	}
	if len(str) >= 2 {
		ret += str[1:]
	}
	return ret
}

// Reticence trucate the string in the space or on pontuation mark and put
// reticences in the resulting string.
func Reticence(str string, length int) string {
	if length > len(str) {
		return str
	}
	var i int
F:
	for i = len(str) - 1; i >= 0; i-- {
		switch str[i] {
		case ' ', ',', '?', ';', ':', '\'', '"', '!':
			if i <= length {
				break F
			}
		case '.':
			if i-2 >= 0 {
				s := str[i-2 : i]
				if s == ".." {
					i = i - 2
					if i <= length {
						break F
					}
				}
			}
			break
		}

	}
	if i >= 2 {
		if i+3 >= len(str) {
			return str
		}
		return str[:i] + "..."
	}
	if length >= 2 && length < len(str) {
		if length+3 >= len(str) {
			return str
		}
		return str[:length] + "..."
	}
	return str
}