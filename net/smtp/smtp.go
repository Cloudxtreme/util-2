// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2012-03-08
// Last modification:	2012-

// Support for self-signed certificate in SendMail function
package smtp

import (
	"crypto/tls"
	"fmt"
	"github.com/fcavani/e"
	"log"
	"net"
	"net/smtp"
	"reflect"
	"time"
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

type Command struct {
	Timeout time.Duration
	Conn    net.Conn
	retvals chan []reflect.Value
}

type Return func(args ...interface{})

func SetError(retvals []reflect.Value, err error) {
	if err == nil {
		return
	}
	var i int
	var retval reflect.Value
	for i, retval = range retvals {
		if reflect.TypeOf(err).AssignableTo(retval.Type()) {
			if retval.CanSet() && retval.IsNil() {
				retval.Set(reflect.ValueOf(err))
			}
		}
	}
	if i == len(retvals) {
		panic("can't assign an error to none of the returned values")
	}
}

func (c *Command) Exec(f interface{}, args ...interface{}) Return {
	c.retvals = make(chan []reflect.Value)
	val := reflect.ValueOf(f)
	if val.Kind() != reflect.Func {
		panic("f is not a function")
	}
	t := val.Type()
	if t.NumIn() != len(args) {
		panic("invalid number of arguments")
	}
	a := make([]reflect.Value, len(args))
	for i, arg := range args {
		a[i] = reflect.ValueOf(arg)
		if !a[i].Type().AssignableTo(t.In(i)) {
			panic(fmt.Sprintf("invalid argument type, argument %v must be assignable to %v", a[i].Type(), t.In(i)))
		}
	}
	go func() {
		err := c.Conn.SetDeadline(time.Now().Add(c.Timeout))
		if err != nil {
			log.Println("SetDeadline failed with error:", e.Trace(e.Forward(err)))
		}
		retvals := val.Call(a)
		SetError(retvals, err)
		c.retvals <- retvals
		close(c.retvals)
	}()

	return func(args ...interface{}) {
		retvals := <-c.retvals
		if len(retvals) != len(args) {
			panic("the number of arguments in Returns must be equal to the number of return values in the function")
		}
		for i, retval := range retvals {
			val := reflect.ValueOf(args[i])
			if val.Kind() != reflect.Ptr {
				panic("Returns arguments must be pointers")
			}
			if retval.Kind() != val.Elem().Kind() {
				panic("diferent kind")
			}
			val.Elem().Set(retval)
		}
	}
}

// TestSMTP tests if can connect with the server and send some commands.
func TestSMTP(addr string, a smtp.Auth, from, hello string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return e.Forward(err)
	}

	command := &Command{
		Timeout: timeout,
		Conn:    conn,
	}

	var c *smtp.Client
	r := command.Exec(smtp.NewClient, conn, addr)
	r(&c, &err)
	if err != nil {
		return e.Forward(err)
	}

	if hello != "" {
		r = command.Exec(c.Hello, hello)
		r(&err)
		if err != nil {
			return e.Forward(err)
		}
	}

	if a != nil {
		if ok, _ := c.Extension("STARTTLS"); ok {
			r = command.Exec(c.StartTLS, &tls.Config{InsecureSkipVerify: true})
			r(&err)
			if err != nil {
				return e.Forward(err)
			}
		}
		found, _ := c.Extension("AUTH")
		if a != nil && found {
			r = command.Exec(c.Auth, a)
			r(&err)
			if err != nil {
				return e.Forward(err)
			}
		}
	}

	r = command.Exec(c.Mail, from)
	r(&err)
	if err != nil {
		return e.Forward(err)
	}

	r = command.Exec(c.Reset)
	r(&err)
	if err != nil {
		return e.New(err)
	}

	r = command.Exec(c.Quit)
	r(&err)
	if err != nil {
		return e.New(err)
	}

	return nil
}
