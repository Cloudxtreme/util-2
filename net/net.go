// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Network util functions.
package net

import (
	"github.com/fcavani/e"
	utilUrl "github.com/fcavani/util/net/url"
	"net"
	"net/url"
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

const ErrHostNotResolved = "host name not resolved"

// ResolveUrl replaces the host name with the ip address.
func ResolveUrl(url *url.URL) (*url.URL, error) {
	if url.Scheme == "file" || url.Scheme == "socket" || url.Scheme == "unix" {
		return utilUrl.Copy(url), nil
	}
	if len(url.Host) > 0 && url.Host[0] == '/' {
		return utilUrl.Copy(url), nil
	}
	if len(url.Host) >= 3 && url.Host[1] == ':' && url.Host[2] == '/' {
		return utilUrl.Copy(url), nil
	}
	host, port, err := SplitHostPort(url.Host)
	if err != nil && !e.Equal(err, ErrCantFindPort) {
		return nil, e.Forward(err)
	}
	addrs, err := net.LookupHost(host)
	if e.Contains(err, "invalid domain name") {
		return utilUrl.Copy(url), nil
	} else if err != nil {
		return nil, e.Forward(err)
	}
	if len(addrs) <= 0 {
		return nil, e.New(ErrHostNotResolved)
	}

	out := utilUrl.Copy(url)

	if strings.Contains(addrs[0], ":") {
		out.Host = "[" + addrs[0] + "]"
	} else {
		out.Host = addrs[0]
	}
	if port != "" {
		out.Host += ":" + port
	}
	return out, nil
}
