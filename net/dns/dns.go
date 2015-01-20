// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Resolve host name stuff.
package dns

import (
	"github.com/fcavani/e"
	utilUrl "github.com/fcavani/util/net/url"
	utilNet "github.com/fcavani/util/net"
	"net"
	"net/url"
	"strings"
)

const ErrHostNotResolved = "host name not resolved"

var lookuphost func(host string)(addrs []string, err error) = net.LookupHost

func SetLookupHostFunction(f func(host string)(addrs []string, err error)) {
	lookuphost = f
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
	host, port, err := utilNet.SplitHostPort(url.Host)
	if err != nil && !e.Equal(err, utilNet.ErrCantFindPort) {
		return nil, e.Forward(err)
	}
	addrs, err := lookuphost(host)
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

