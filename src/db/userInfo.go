package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Userinfo is table
type UserInfo struct {
	gorm.Model
	Id string
	Name string
}

// insert userInfo
func InsertUser(userInfo UserInfo){
	db := gormConnect()
	defer db.Close()
	
	db.NewRecord(userInfo)
	db.Create(&userInfo)
	log.Println("ユーザの登録をしました。ID:" + userInfo.Id + " Name:"+userInfo.Name)
}

// update user name
func UpdateUserName(userInfo UserInfo){
	db := gormConnect()
	defer db.Close()

	db.Model(&userInfo).Where("id = ?", userInfo.Id).Update("name", userInfo.Name)
	log.Println("ユーザの名前を変更しました。対象ID:" + userInfo.Id + " 変更名前:" + userInfo.Name)
}

// delete userInfo by id
func DeleteUserInfoById(id string){
	db := gormConnect()
	defer db.Close()

	d := UserInfo{Id:id}
	db.Delete(d)
	log.Println("ユーザを削除しました。対象ID:" + id)
}

func GetAllUserInfos() []UserInfo{
	db := gormConnect()
	defer db.Close()

	userInfos := []UserInfo{}
	db.Find(&userInfos)
	return userInfos
}