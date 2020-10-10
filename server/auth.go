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

func check(name string, pass string) int {
	c := config.Db()

	headpass := string(pass[0])

	switch headpass {
	case "a":
		var rpassword string
		var wpassword string
		err := c.QueryRow("select rpassword, wpassword from user_datasets where id = $1", name).Scan(&rpassword, &wpassword)

		if err != nil {
			fmt.Println("--------------------")
			fmt.Println(err)
			return 0
		}

		if pass == rpassword {
			return 1
		}

		if pass == wpassword {
			return 2
		}

	case "b":
		var ftppassword string
		err := c.QueryRow("select ftppassword from user_favor_datasets where id = $1", name).Scan(&ftppassword)

		if err != nil {
			fmt.Println(err)
			return 0
		}

		if pass == ftppassword {
			return 1
		}

	case "c":
		var ftppassword string
		err := c.QueryRow("select ftppassword from gscloud_batch_info where batchid = $1", name).Scan(&ftppassword)

		if err != nil {
			return 0
		}

		if pass == ftppassword {
			return 1
		}
	default:
		return 0
	}
	fmt.Println(0000000000)
	return 0
}

// CheckPasswd will check user's password
// func (a *SimpleAuth) CheckPasswd(name, pass string) (int, error) {
func CheckPasswd(name, pass string) (int, error) {
	return check(name, pass), nil
	// return constantTimeEquals(name, a.Name) && constantTimeEquals(pass, a.Password), nil
}

func constantTimeEquals(a, b string) bool {
	return len(a) == len(b) && subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
