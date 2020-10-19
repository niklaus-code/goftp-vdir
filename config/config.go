package config

import (
	"database/sql"
	"fmt"

	"github.com/go-ini/ini"
	_ "github.com/lib/pq"
)

var Dbname = "postgres"
var cfg, _ = ini.Load("conf/setting.ini")

func Db() *sql.DB {

	var ip = cfg.Section(Dbname).Key("ip").String()
	var port = cfg.Section(Dbname).Key("port").String()
	var user = cfg.Section(Dbname).Key("user").String()
	// var passwd = cfg.Section(Dbname).Key("passwd").String()
	var database = cfg.Section(Dbname).Key("database").String()

	conn := fmt.Sprintf("host=%s  user=%s  dbname=%s port=%s sslmode=disable", ip, user, database, port)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil
	}
	return db
}

func Download_rate() int {
	rate, err := cfg.Section("download").Key("rate").Int()
	if err != nil {
		return 0
	}
	return rate
}
