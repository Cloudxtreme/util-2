// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package url

import (
	"github.com/fcavani/e"
	"math"
	"net/url"
	"regexp"
	"strings"
)

//"mysql://root:pass@unix(/var/run/mysql.socket)/db"
func Socket(host string) (string, string, error) {
	r, err := regexp.Compile(`([A-Za-z0-9]*)\((.*)\)`)
	if err != nil {
		return "", "", e.New(err)
	}
	sub := r.FindStringSubmatch(host)
	if len(sub) != 3 {
		return "", "", e.New("path not found")
	}
	return sub[1], sub[2], nil
}

func ParseWithSocket(url_ string) (*url.URL, error) {
	u := new(url.URL)
	s := strings.SplitN(url_, "://", 2)
	if len(s) != 2 {
		return nil, e.New("invalid url")
	}
	u.Scheme = s[0]
	rest := ""
	s = strings.SplitN(s[1], "@", 2)
	if len(s) == 1 {
		rest = s[0]
	} else if len(s) == 2 {
		userpass := strings.SplitN(s[0], ":", 2)
		if len(userpass) == 1 {
			u.User = url.User(userpass[0])
		} else if len(userpass) == 2 {
			pass, err := url.QueryUnescape(userpass[1])
			if err != nil {
				return nil, e.New(err)
			}
			u.User = url.UserPassword(userpass[0], pass)
		} else {
			return nil, e.New("invalid user password")
		}
		rest = s[1]
	} else {
		return nil, e.New("invalid user string")
	}

	r, err := regexp.Compile(`.*\(.*\)`)
	if err != nil {
		return nil, e.New(err)
	}
	unix := r.FindAllString(rest, 1)
	if len(unix) == 1 {
		u.Host = unix[0]
		rest = strings.TrimSpace(r.ReplaceAllLiteralString(rest, ""))
		q := strings.Index(rest, "?")
		f := strings.Index(rest, "#")
		pend := f
		if q > f {
			pend = q
		}
		i := strings.Index(rest, "/")
		if i != -1 && pend != -1 {
			u.Path = rest[i:pend]
			rest = rest[pend:]
		} else if i == -1 && pend != -1 {
			rest = rest[pend:]
		} else if i != -1 && pend == -1 {
			u.Path = rest[i:]
			return u, nil
		} else if i == -1 && pend == -1 {
			return u, nil
		}
	} else if len(unix) == 0 {
		q := strings.Index(rest, "?")
		f := strings.Index(rest, "#")
		pend := f
		ff := f
		if f == -1 {
			ff = math.MaxInt64
		}
		if q < ff && q >= 0 {
			pend = q
		}
		i := strings.Index(rest, "/")
		if i != -1 && pend != -1 {
			u.Host = rest[:i]
			u.Path = rest[i:pend]
			rest = rest[pend:]
		} else if i == -1 && pend != -1 {
			u.Host = rest[:pend]
			rest = rest[pend:]
		} else if i != -1 && pend == -1 {
			u.Host = rest[:i]
			u.Path = rest[i:]
			return u, nil
		} else if i == -1 && pend == -1 {
			u.Host = rest
			return u, nil
		}
	} else {
		return nil, e.New("socket address is invalid")
	}

	q := strings.Index(rest, "?")
	f := strings.Index(rest, "#")

	if q+1 >= len(rest) {
		return nil, e.New("error parsing query")
	}
	if f+1 >= len(rest) {
		return nil, e.New("error parsing fragment")
	}

	if q != -1 && f != -1 && q <= f {
		u.RawQuery = rest[q+1 : f]
		u.Fragment = rest[f+1:]
	} else if q != -1 && f == -1 {
		u.RawQuery = rest[q+1:]
	} else if q == -1 && f != -1 {
		u.Fragment = rest[f+1:]
	} else if q == -1 && f == -1 {
		return u, nil
	} else {
		return nil, e.New("error parsing query and fragment %v, %v", q, f)
	}
	return u, nil
}
