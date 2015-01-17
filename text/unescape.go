// Copyright 2013 Felipe Alves Cavani. All rights reserved.
// Start date:		2013-07-04
// Last modification:	2013-x

package text

import (
	"net/url"
	"strings"
)

func EscapeCommaSeparated(in ...string) string {
	var out string
	for i, str := range in {
		escaped := strings.Replace(url.QueryEscape(str), "%2F", "%252F", -1)
		escaped = strings.Replace(escaped, "\"", "%22", -1)
		escaped = strings.Replace(escaped, " ", "%20", -1)
		out += escaped
		if i < len(in)-1 {
			out += ","
		}
	}
	return out
}
