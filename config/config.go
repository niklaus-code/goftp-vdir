package config

import (
	"database/sql"
	"fmt"

	"github.com/go-ini/ini"
	_ "github.com/lib/pq"
)

func Db() *sql.DB {
	cfg, err := ini.Load("conf/setting.ini")
	if err != nil {
		fmt.Println(err)
	}
	var ip = cfg.Section("postgres").Key("ip").String()
	// var port = cfg.Section("postgres").Key("port").String()
	var user = cfg.Section("postgres").Key("user").String()
	// var passwd = cfg.Section("postgres").Key("passwd").String()
	var database = cfg.Section("postgres").Key("database").String()

	conn := fmt.Sprintf("host=%s  user=%s  dbname=%s  sslmode=disable", ip, user, database)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("222222")
		fmt.Println(err)
		return nil
	}
	return db
}

func Download_rate() int {
	cfg, err := ini.Load("conf/setting.ini")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	rate, err := cfg.Section("download").Key("rate").Int()
	if err != nil {
		return 0
	}
	return rate
}
