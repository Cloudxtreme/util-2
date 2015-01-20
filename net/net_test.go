// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/fcavani/e"
	"net/url"
	"testing"
)

type testHostPortStruct struct {
	hostport string
	host     string
	port     string
	fail     bool
}

var testhp []testHostPortStruct = []testHostPortStruct{
	{"[2001:db8:1f70::999:de8:7648:6e8]:100", "2001:db8:1f70::999:de8:7648:6e8", "100", false},
	{"127.0.0.1:169", "127.0.0.1", "169", false},
	{"www.isp.net:8080", "www.isp.net", "8080", false},
	{"www.isp.net", "", "", true},
	{"[2001:db8:1f70::999:de8:7648:6e8]", "", "", true},
	{"www.isp.net:", "", "", true},
	{"[2001:db8:1f70::999:de8:7648:6e8]:", "", "", true},
}

func TestSplitHostPort(t *testing.T) {
	for i, thp := range testhp {
		host, port, err := SplitHostPort(thp.hostport)
		if err != nil && !thp.fail {
			t.Fatal(i, e.Trace(e.Forward(err)))
		} else if err == nil && thp.fail {
			t.Fatal(i, "doesn't failed", host, port)
		}
		if host != thp.host {
			t.Fatal("wrong host", i, host)
		}
		if port != thp.port {
			t.Fatal("wrong port", i, port)
		}
	}
}

func TestResolveUrl(t *testing.T) {
	url, err := url.Parse("http://localhost:8080/foo.html?q=search#fragment")
	if err != nil {
		t.Fatal("parse failed", err)
	}
	u, err := ResolveUrl(url)
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	if u.Host != "127.0.0.1:8080" && u.Host != "[::1]:8080" {
		t.Fatal("can't resolve", u)
	}
	t.Log(u)
}
