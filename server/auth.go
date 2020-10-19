// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"crypto/subtle"
	"fmt"

	"github.com/niklaus-code/goftp-vdir/config"
)

// Auth is an interface to auth your ftp user login.
type Auth interface {
	CheckPasswd(string, string) (int, error)
}

// var (
// 	_ Auth = &SimpleAuth{}
// )

// SimpleAuth implements Auth interface to provide a memory user login auth
// type SimpleAuth struct {
// 	Name     string
// 	Password string
// }
type Ftpusr struct {
	Rpassword  string
	Wpassword  string
	Privileges int
	Datapath   string
}

func check(name string, pass string) Ftpusr {
	c := config.Db()

	headpass := string(pass[0])

	var ftpusr Ftpusr

	switch headpass {
	case "a":
		err := c.QueryRow("select rpassword, wpassword, datapath from user_datasets where id = $1", name).Scan(&ftpusr.Rpassword, &ftpusr.Wpassword, &ftpusr.Datapath)

		if err != nil {
			fmt.Println("--------------------")
			fmt.Println(err)
			return ftpusr
		}
		switch {
		case pass == ftpusr.Rpassword:
			ftpusr.Privileges = 1

		case pass == ftpusr.Wpassword:
			ftpusr.Privileges = 2
		default:
			ftpusr.Privileges = 0
		}

	case "b":
		err := c.QueryRow("select ftppassword from user_favor_datasets where id = $1 and ftppassword = $2", name, pass).Scan(&ftpusr.Rpassword)

		if err != nil {
			return ftpusr
		}
		ftpusr.Privileges = 1
		ftpusr.Datapath = "/tmp"
		return ftpusr

	case "c":
		err := c.QueryRow("select ftppassword from gscloud_batch_info where batchid = $1 and ftppassword = $2", name, pass).Scan(&ftpusr.Rpassword)

		if err != nil {
			return ftpusr
		}
		ftpusr.Privileges = 1
		ftpusr.Datapath = "/tmp"
		return ftpusr

	default:
		return ftpusr
	}
	return ftpusr
}

// CheckPasswd will check user's password
// func (a *SimpleAuth) CheckPasswd(name, pass string) (int, error) {
func CheckPasswd(name, pass string) (Ftpusr, error) {
	return check(name, pass), nil
	// return constantTimeEquals(name, a.Name) && constantTimeEquals(pass, a.Password), nil
}

func constantTimeEquals(a, b string) bool {
	return len(a) == len(b) && subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
