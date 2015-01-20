// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Network util functions.
package net

import (
	"github.com/fcavani/e"
	"regexp"
	"strings"
)

const ErrCantGetIp = "can't get remote ip"
const ErrCantSplitHostPort = "can't split host port"
const ErrCantFindHost = "can't find the host"
const ErrCantFindPort = "can't find the port number"

// SplitHostPort splits a string with a ipv6, ipv4 or hostname with a port number.
func SplitHostPort(hp string) (host, port string, err error) {
	if len(hp) == 0 {
		return "", "", e.New("invalid host length")
	}
	if hp[0] == '[' {
		// ipv6 - [2001:db8:1f70::999:de8:7648:6e8]:100
		r, err := regexp.Compile(`\[([a-z0-9:]*)\]\:([0-9]*)`)
		if err != nil {
			return "", "", e.Push(err, "can't compile ipv6 regexp")
		}
		x := r.FindAllStringSubmatch(hp, -1)
		if len(x) <= 0 {
			return "", "", e.New(ErrCantGetIp)
		}
		if len(x[0]) != 3 {
			return "", "", e.New(ErrCantGetIp)
		}
		host = x[0][1]
		port = x[0][2]
	} else {
		//ip4 and host name
		ipport := strings.SplitN(hp, ":", 2)
		if len(ipport) != 2 {
			return "", "", e.New(ErrCantSplitHostPort)
		}
		host = ipport[0]
		port = ipport[1]
	}
	if host == "" {
		return "", "", e.New(ErrCantFindHost)
	}
	if port == "" {
		return "", "", e.New(ErrCantFindPort)
	}
	return
}