// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux, bsd, darwin

package os

import (
	"github.com/fcavani/e"
	"os/exec"
	"bytes"
	"strings"
	"strconv"
)

type Process struct {
	User string
	Pid  int
	Cpu float32
	Memory float32
	Vsz	uint
	Rss uint
}

// Processes return information about the running processes (user, pid, cpu percentage usage,
// memory percentage usage, vsz and rss). All information is parsed from ps aux.
func Processes() (map[int]*Process, error) {
	// From: http://stackoverflow.com/questions/11356330/getting-cpu-usage-with-golang
    cmd := exec.Command("ps", "aux")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return nil, e.New(err)
    }
    processes := make(map[int]*Process)
    for {
        line, err := out.ReadString('\n')
        if err!=nil {
            break;
        }
        tokens := strings.Split(line, " ")
        ft := make([]string, 0)
        for _, t := range(tokens) {
            if t != "" && t != "\t" {
                ft = append(ft, t)
            }
        }
		user := ft[0]
		if user == "USER" {
			continue
		}
        pid, err := strconv.ParseInt(ft[1], 10, 32)
        if err != nil {
            return nil, e.Push(err, "failed to retrieve processes data")
        }
        cpu, err := strconv.ParseFloat(strings.Replace(ft[2], ",", ".", 1), 32)
        if err != nil {
            return nil, e.Push(err, "failed to retrieve processes data")
        }
        mem, err := strconv.ParseFloat(strings.Replace(ft[3], ",", ".", 1), 32)
        if err != nil {
            return nil, e.Push(err, "failed to retrieve processes data")
        }
		vsz, err := strconv.ParseUint(ft[4], 10, 32)
        if err != nil {
            return nil, e.Push(err, "failed to retrieve processes data")
        }
		rss, err := strconv.ParseUint(ft[5], 10, 32)
        if err != nil {
            return nil, e.Push(err, "failed to retrieve processes data")
        }
		
        processes[int(pid)] = &Process{
			User: user,
			Pid: int(pid),
			Cpu: float32(cpu)/100.0,
			Memory: float32(mem)/100.0,
			Vsz: uint(vsz),
			Rss: uint(rss),
        }
    }
	return processes, nil
}