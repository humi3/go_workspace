package db

import (
	"fmt"
	"log"
	"os"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	ini "gopkg.in/ini.v1"
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

var config DbConfig

func init() {
	projectPath, pErr := os.Getwd()
	
	if pErr != nil {
		fmt.Print(pErr)
		os.Exit(1)
	}

	c, err := ini.Load(projectPath + "/config/db_config.ini")
	
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		log.Println(err.Error())
		os.Exit(1)
	}

	config = DbConfig{
		User: c.Section("db").Key("user").String(),
		Pass: c.Section("db").Key("pass").String(),
		Port: c.Section("db").Key("port").String(),
		Ip: c.Section("db").Key("ip").String(),
		Dbms: c.Section("db").Key("dbms").String(),
		Dbname: c.Section("db").Key("dbname").String(),
	}
}

// db connect
func gormConnect() *gorm.DB {
  DBMS     := config.Dbms
  USER     := config.User
  PASS     := config.Pass
  PROTOCOL := "tcp("+ config.Ip + ":" + config.Port + ")"
	DBNAME   := config.Dbname
	
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
  db,err := gorm.Open(DBMS, CONNECT)
	
  if err != nil {
		fmt.Println(CONNECT)
    panic(err.Error())
  }
  return db
}