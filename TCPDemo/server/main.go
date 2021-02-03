package main

import (
	"TCPDemo/server/module"
	"TCPDemo/server/process"
	"github.com/fatih/color"
	"log"
	"net"
	"time"
)

type Status int

const (
	ONLINE  Status = 1
	OFFLINE Status = 0
)

type User struct {
	Uid    string
	Name   string
	Status Status
}

func init() {
	module.InitPool(16, 0, 300*time.Second, "127.0.0.1:6379")
	module.MyUserDao = module.NewUserDao(module.Pool)
}

func main() {

	color.Green("服务器开始监听...")

	listener, err := net.Listen("tcp", "127.0.0.1:6066")
	defer listener.Close()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			log.Println(err)
			continue
		}

		go func(conn net.Conn) {
			color.Yellow("来自%v的访问\n", conn.RemoteAddr())
			defer conn.Close()
			processor := &process.Processor{
				Conn: conn,
			}
			err = processor.ServerProcessMess()

		}(conn)
	}

}
