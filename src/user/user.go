package user

import "fmt"

// user is struct
type UserInfo struct {
	Id string
	Name string
}

// EchoUserInfo is UserInfo print
func EchoUserInfo (user UserInfo){
	fmt.Println(user.Id)
	fmt.Println(user.Name)
}