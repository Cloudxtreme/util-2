// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:        2014-10-10
// Last modification: 2014-

package url

import (
	"github.com/fcavani/e"
	"testing"
)

func TestSocket(t *testing.T) {
	method, path, err := Socket("unix(/path)")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	if method != "unix" {
		t.Fatal("wrong path", method)
	}
	if path != "/path" {
		t.Fatal("wrong path", path)
	}
}

func TestParse(t *testing.T) {
	url, err := ParseWithSocket("http://localhost")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	t.Logf("%#v\n", url)
	if url.Scheme != "http" {
		t.Fatal("wrong scheme")
	}
	if url.Host != "localhost" {
		t.Fatal("wrong host")
	}

	url, err = ParseWithSocket("http://www.isp.net:80/path#frag")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	t.Logf("%#v\n", url)
	if url.Scheme != "http" {
		t.Fatal("wrong scheme")
	}
	if url.Host != "www.isp.net:80" {
		t.Fatal("wrong host")
	}
	if url.Path != "/path" {
		t.Fatal("wrong path")
	}
	if url.Fragment != "frag" {
		t.Fatal("wrong fragment")
	}

	url, err = ParseWithSocket("http://www.isp.net:80/path?q=1")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	t.Logf("%#v\n", url)
	if url.Scheme != "http" {
		t.Fatal("wrong scheme")
	}
	if url.Host != "www.isp.net:80" {
		t.Fatal("wrong host")
	}
	if url.Path != "/path" {
		t.Fatal("wrong path")
	}
	if url.RawQuery != "q=1" {
		t.Fatal("wrong fragment")
	}

	url, err = ParseWithSocket("http://www.isp.net:80/path?q=1")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	t.Logf("%#v\n", url)
	if url.Scheme != "http" {
		t.Fatal("wrong scheme")
	}
	if url.Host != "www.isp.net:80" {
		t.Fatal("wrong host")
	}
	if url.Path != "/path" {
		t.Fatal("wrong path")
	}
	if url.RawQuery != "q=1" {
		t.Fatal("wrong fragment")
	}

	url, err = ParseWithSocket("http://www.isp.net:80/path?q=1#frag")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	t.Logf("%#v\n", url)
	if url.Scheme != "http" {
		t.Fatal("wrong scheme")
	}
	if url.Host != "www.isp.net:80" {
		t.Fatal("wrong host")
	}
	if url.Path != "/path" {
		t.Fatal("wrong path")
	}
	if url.RawQuery != "q=1" {
		t.Fatal("wrong fragment")
	}
	if url.Fragment != "frag" {
		t.Fatal("wrong fragment")
	}

	url, err = ParseWithSocket("http://unix(/var/run/app.socket)/path#frag")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	t.Logf("%#v\n", url)
	if url.Scheme != "http" {
		t.Fatal("wrong scheme")
	}
	if url.Host != "unix(/var/run/app.socket)" {
		t.Fatal("wrong host")
	}
	if url.Path != "/path" {
		t.Fatal("wrong path")
	}
	if url.Fragment != "frag" {
		t.Fatal("wrong fragment")
	}

	url, err = ParseWithSocket("mysql://tcp(www.isp.net)/path#frag")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	t.Logf("%#v\n", url)
	if url.Scheme != "mysql" {
		t.Fatal("wrong scheme")
	}
	if url.Host != "tcp(www.isp.net)" {
		t.Fatal("wrong host")
	}
	if url.Path != "/path" {
		t.Fatal("wrong path")
	}
	if url.Fragment != "frag" {
		t.Fatal("wrong fragment")
	}

	url, err = ParseWithSocket("socket:///var/run/file.socket#frag")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	t.Logf("%#v\n", url)
	if url.Scheme != "socket" {
		t.Fatal("wrong scheme")
	}
	if url.Host != "/var/run/file.socket" {
		t.Fatal("wrong host")
	}
	if url.Path != "" {
		t.Fatal("wrong path")
	}
	if url.Fragment != "frag" {
		t.Fatal("wrong fragment")
	}

	url, err = ParseWithSocket("foo://C:/bar/file.socket#frag")
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
	t.Logf("%#v\n", url)
	if url.Scheme != "foo" {
		t.Fatal("wrong scheme")
	}
	if url.Host != "C:/bar/file.socket" {
		t.Fatal("wrong host")
	}
	if url.Path != "" {
		t.Fatal("wrong path")
	}
	if url.Fragment != "frag" {
		t.Fatal("wrong fragment")
	}
}
