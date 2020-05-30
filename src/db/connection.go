package db

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

// DbConfig is config
type DbConfig struct {
	//User is DB user name
		User string
		// Pass is DB password
		Pass string
		// Port is DB port
		Port string
		// Ip is DB ip
		Ip string
		// Dbms is DB type
		Dbms string
		// DBname is databese name
		Dbname string
}

var Config DbConfig

func init() {
	c, err := ini.Load("./db_config.ini")

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	Config = DbConfig{
		User: c.Section("db").Key("user").MustInt(),
		Pass: c.Section("db").Key("pass").MustInt(),
		Port: c.Section("db").Key("port").MustInt(),
		Ip: c.Section("db").Key("ip").MustInt(),
		Dbms: c.Section("db").Key("dbms").MustInt(),
		Dbname: c.Section("db").Key("dbname").MustInt(),
	}
}