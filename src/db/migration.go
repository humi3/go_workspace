package db

import "log"

// is db maigration.
func StartDBMigration(){
	db := gormConnect()
	defer db.Close()

	log.Println("----DBのマイニングを実行-----")

	log.Println("table:user_info")
	db.AutoMigrate(&UserInfo{})

	log.Println("----DBのマイニングを終了-----")
}