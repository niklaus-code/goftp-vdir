// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"crypto/subtle"
	"fmt"

	"github.com/niklaus-code/hello/db"
)

// Auth is an interface to auth your ftp user login.
type Auth interface {
	CheckPasswd(string, string) (int, error)
	//CheckPasswd(string, string) (bool, error)
}

var (
	_ Auth = &SimpleAuth{}
)

// SimpleAuth implements Auth interface to provide a memory user login auth
type SimpleAuth struct {
	Name     string
	Password string
}

func check(name string, pass string) int {
	c := db.Db()

	var rpassword string
	var wpassword string
	rows, err := c.Query("select rpassword, wpassword from user_datasets where id = $1", name)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	for rows.Next() {
		rows.Scan(&rpassword, &wpassword)
	}

    if pass == rpassword {
        return 1
        }

    if pass == wpassword {
        return 2
        }
    return 0
    /*
	var password string
	rows, err := c.Query("select password from auth_user where userid = $1", name)

	if err != nil {
		fmt.Println(err)
		return false
	}

	for rows.Next() {
		rows.Scan(&password)
	}

	pwd := encryption(password)
	if subtle.ConstantTimeCompare([]byte(pwd), []byte(password)) == 1 {
		return true
	}
	return false
    */
}

// CheckPasswd will check user's password
func (a *SimpleAuth) CheckPasswd(name, pass string) (int, error) {
	return check(name, pass), nil
	// return constantTimeEquals(name, a.Name) && constantTimeEquals(pass, a.Password), nil
}

func constantTimeEquals(a, b string) bool {
	return len(a) == len(b) && subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
