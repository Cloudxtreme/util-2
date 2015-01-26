// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Network util functions.
package net

import (
	"github.com/fcavani/e"
	"regexp"
	"strings"
	"net"
)

const ErrCantGetIp = "can't get remote ip"
const ErrCantSplitHostPort = "can't split host port"
const ErrCantFindHost = "can't find the host"
const ErrCantFindPort = "can't find the port number"

func IsValidIpv4(ip string) bool {
	r, err := regexp.Compile(`^(([0-9]{1,3}\.){3}[0-9]{1,3})$`)
	if err != nil {
		panic("can't compile ipv6 regexp")
	}
	x := r.FindAllStringSubmatch(ip, -1)
	if len(x) == 0 {
		return false
	}
	if len(x[0]) != 3 {
		return false
	}
	if x[0][1] != ip {
		return false
	}
	ipParsed := net.ParseIP(ip)
	if ipParsed == nil {
		return false
	}
	return true
}
	
func IsValidIpv6(ip string) bool {
	r, err := regexp.Compile(`\[([a-z0-9:]*)\]`)
	if err != nil {
		panic("can't compile ipv6 regexp")
	}
	x := r.FindAllStringSubmatch(ip, -1)
	if len(x) == 0 {
		return false
	}
	if len(x[0]) != 2 {
		return false
	}
	ip = strings.TrimSuffix(strings.TrimPrefix(ip, "["), "]")
	if x[0][1] !=  ip {
		return false
	}
	ipParsed := net.ParseIP(ip)
	if ipParsed == nil {
		return false
	}
	return true
}

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
		if len(x) == 0 {
			if IsValidIpv6(hp) {
				host = strings.TrimSuffix(strings.TrimPrefix(hp, "["), "]")
				port = ""
			} else {
				return "", "", e.New(ErrCantGetIp)
			}
		} else {
			if len(x[0]) == 2 {
				host = x[0][1]
				port = ""
			} else if len(x[0]) == 3 {
				host = x[0][1]
				port = x[0][2]
			} else {
				return "", "", e.New(ErrCantGetIp)
			}
		}
	} else {
		//ip4 and host name
		ipport := strings.SplitN(hp, ":", 2)
		if len(ipport) == 1 {
			host = ipport[0]
			port = ""
		} else if len(ipport) == 2 {
			host = ipport[0]
			port = ipport[1]
		} else {
			return "", "", e.New(ErrCantSplitHostPort)
		}
	}
	if host == "" {
		return "", "", e.New(ErrCantFindHost)
	}
	if port == "" {
		return host, "", e.New(ErrCantFindPort)
	}
	return
}
