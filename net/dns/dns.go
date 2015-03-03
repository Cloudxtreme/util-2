// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Resolve host name stuff.
package dns

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/fcavani/e"
	utilNet "github.com/fcavani/util/net"
	utilUrl "github.com/fcavani/util/net/url"
	"github.com/miekg/dns"
)

const ErrHostNotResolved = "host name not resolved"

var lookuphost func(host string) (addrs []string, err error) = LookupHost

func SetLookupHostFunction(f func(host string) (addrs []string, err error)) {
	lookuphost = f
}

var Timeout = 30
var ConfigurationFile = "/etc/resolv.conf"

func LookupHost(host string) (addrs []string, err error) {
	if utilNet.IsValidIpv4(host) || utilNet.IsValidIpv6(host) {
		return []string{host}, nil
	}
	config, err := dns.ClientConfigFromFile(ConfigurationFile)
	if err != nil {
		return nil, e.Forward(err)
	}
	config.Timeout = Timeout

	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(host), dns.TypeA)
	r, _, err := c.Exchange(m, config.Servers[0]+":"+config.Port)
	if err != nil {
		return nil, e.Forward(err)
	}
	if r.Rcode != dns.RcodeSuccess {
		return nil, e.New("no success")
	}

	addrs = make([]string, 0, 10)
	for _, a := range r.Answer {
		if addr, ok := a.(*dns.A); ok {
			addrs = append(addrs, addr.A.String())
		}
	}

	m.SetQuestion(dns.Fqdn(host), dns.TypeAAAA)
	r, _, err = c.Exchange(m, config.Servers[0]+":"+config.Port)
	if err != nil {
		return nil, e.Forward(err)
	}
	if r.Rcode != dns.RcodeSuccess {
		return nil, e.New("no success")
	}

	for _, a := range r.Answer {
		if addr, ok := a.(*dns.AAAA); ok {
			addrs = append(addrs, addr.AAAA.String())
		}
	}
	return
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
