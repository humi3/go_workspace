package main

import (
	"fmt"

	"github.com/humi3/go_workspace/src/user"
)

func main() {
	kurimoto := user.UserInfo{"0001","kurimoto"}

	fmt.Println(kurimoto.Id)
	fmt.Println(kurimoto.Name)
	fmt.Println("Hello world！")
}