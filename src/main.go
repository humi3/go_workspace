package main

import (
	"fmt"

	db "github.com/humi3/go_workspace/src/db"
)

func main() {
	// db.StartDBMigration()


	// user1 := db.UserInfo{Id:"2",Name:"test1"} 
	// user2 := db.UserInfo{Id:"3",Name:"test2"} 
	// user3 := db.UserInfo{Id:"4",Name:"test3"} 

	// db.InsertUser(user1)
	// db.InsertUser(user2)
	// db.InsertUser(user3)

	userInfos := db.GetAllUserInfos()
	for _,user := range userInfos {
		fmt.Println("id:"+ user.Id + " name:" + user.Name)
	}
}