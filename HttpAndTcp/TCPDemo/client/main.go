package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/howeyc/gopass"

	"TCPDemo/client/common"
	"TCPDemo/client/login"
)

func main() {
	var option common.Option
	var loop bool
	var user common.User
	for {
		color.Green("----------------- Golang Chat ------------------")
		color.Green("\t\t1 登录\t\t\t")
		color.Green("\t\t2 注册\t\t\t")
		color.Green("\t\t3 退出\t\t\t")

		fmt.Scanf("%d\n", &option)

		switch option {
		case common.LOGIN:
			color.Cyan("登录中......\n")
			loop = false
		case common.LOGUP:
			color.Cyan("跳转中......\n")
			loop = false
		case common.EXIT:
			color.Cyan("退出中......\n")
			loop = false
		default:
			color.Red("输入有误")
		}

		if option == common.LOGIN {

			color.Yellow("Input You Uid please......\n")
			fmt.Scanf("%s\n", &user.Uid)
			color.Yellow("Input You Passwd please......\n")
			temp, _ := gopass.GetPasswdMasked()
			user.Pwd = string(temp)

			login.LogIn(user)

		} else if option == common.LOGUP {

		}
		fmt.Println(loop)
	}

}
