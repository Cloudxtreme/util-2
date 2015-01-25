// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2012-03-08
// Last modification:	2012-

// Support for self-signed certificate in SendMail function
package smtp

import (
	"crypto/tls"
	"github.com/fcavani/e"
	"net/smtp"
)

// Generate a comma separated list of e-mails from a array of e-mails
func EmailsToString(mails []string) (s string) {
	for i, mail := range mails {
		if i > 0 {
			s += ", "
		}
		s += mail
	}
	return
}

// SendMail send a message to specific destination (to) using smtp server in addrs
// and a auth.
func SendMail(addr string, a smtp.Auth, from string, to []string, hello string, msg []byte) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return e.New(err)
	}
	if hello != "" {
		err = c.Hello(hello)
		if err != nil {
			return e.New(err)
		}
	}
	if a != nil {
		if ok, _ := c.Extension("STARTTLS"); ok {
			if err = c.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
				return e.New(err)
			}
		}
		found, _ := c.Extension("AUTH")
		if a != nil && found {
			if err = c.Auth(a); err != nil {
				return e.New(err)
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return e.New(err)
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return e.New(err)
		}
	}
	w, err := c.Data()
	if err != nil {
		return e.New(err)
	}
	_, err = w.Write(msg)
	if err != nil {
		return e.New(err)
	}
	err = w.Close()
	if err != nil {
		return e.New(err)
	}
	err = c.Quit()
	if err != nil {
		return e.New(err)
	}
	return nil
}

// TestSMTP tests if can connect with the server and send some commands.
func TestSMTP(addr string, a smtp.Auth, from, hello string) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return e.New(err)
	}
	if hello != "" {
		err = c.Hello(hello)
		if err != nil {
			return e.New(err)
		}
	}
	if a != nil {
		if ok, _ := c.Extension("STARTTLS"); ok {
			if err = c.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
				return e.New(err)
			}
		}
		found, _ := c.Extension("AUTH")
		if a != nil && found {
			if err = c.Auth(a); err != nil {
				return e.New(err)
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return e.New(err)
	}
	if err = c.Reset(); err != nil {
		return e.New(err)
	}
	err = c.Quit()
	if err != nil {
		return e.New(err)
	}
	return nil
}
