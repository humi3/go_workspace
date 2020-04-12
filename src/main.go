package main

import (
	"fmt"

	"github.com/humi3/go_workspace/src/user"
)

func main() {
	user := user.UserInfo{"0001","kurimoto"}

	user.EchoUserInfo(user)

	fmt.Println("Hello world！")
}