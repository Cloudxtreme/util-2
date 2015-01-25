// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Resolve host name stuff.
package dns

import (
	"github.com/fcavani/e"
	utilNet "github.com/fcavani/util/net"
	utilUrl "github.com/fcavani/util/net/url"
	"net"
	"net/url"
	"regexp"
	"strings"
)

const ErrHostNotResolved = "host name not resolved"

var lookuphost func(host string) (addrs []string, err error) = net.LookupHost

func SetLookupHostFunction(f func(host string) (addrs []string, err error)) {
	lookuphost = f
}

// Resolve simple resolver one host name to one ip
func Resolve(h string) (out string, err error) {
	host, port, err := utilNet.SplitHostPort(h)
	if err != nil && !e.Equal(err, utilNet.ErrCantFindPort) {
		return "", e.Forward(err)
	}
	
	addrs, err := lookuphost(host)
	if err != nil {
		return "", e.Forward(err)
	}
	if len(addrs) <= 0 {
		return "", e.New(ErrHostNotResolved)
	}

	if strings.Contains(addrs[0], ":") {
		out = "[" + addrs[0] + "]"
	} else {
		out = addrs[0]
	}
	if port != "" {
		out += ":" + port
	}
	return
}

// ResolveUrl replaces the host name with the ip address. Supports ipv4 and ipv6.
// If use in the place of host a path or a scheme for sockets, file or unix,
// ResolveUrl will only copy the url.
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
	r, err := regexp.Compile(`.*\(.*\)`)
	if err != nil {
		return nil, e.New(err)
	}
	mysqlNotation := r.FindAllString(url.Host, 1)
	if len(mysqlNotation) >= 1 {
		return utilUrl.Copy(url), nil
	}
	
	out := utilUrl.Copy(url)
	
	host, err := Resolve(url.Host)
	if err != nil {
		return nil, e.Forward(err)
	}
	out.Host = host
	return out, nil
}
