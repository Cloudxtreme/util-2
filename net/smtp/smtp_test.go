// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2015-01-26
// Last modification:	2015-

// Support for self-signed certificate in SendMail function
package smtp

import (
	"testing"
	"time"

	"github.com/fcavani/e"
)

func TestTestSMTP(t *testing.T) {
	err := TestSMTP("smtp.atlantida.etc.br:587", nil, "none@isp.com", "localhost", 10*time.Second, true)
	if err != nil {
		t.Fatal(e.Trace(e.Forward(err)))
	}
}
