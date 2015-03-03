// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dns

import (
	"net/url"
	"testing"

	"github.com/fcavani/e"
)

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
