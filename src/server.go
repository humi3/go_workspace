package main

import (
	"log"

	"github.com/humi3/go_workspace/src/db"
)

func main() {
	log.Println("起動します。")

	user := db.UserInfo{Id: "test", Name: "test"}
	db.InsertUser(user)
}
