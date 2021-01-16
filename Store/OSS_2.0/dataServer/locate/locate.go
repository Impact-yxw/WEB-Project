package locate

import (
	"OSS/dataServer/rabbitmq"
	"log"
	"os"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate(rabbitmqAddr, url, dir string) {
	q := rabbitmq.New(rabbitmqAddr)
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		log.Println(string(msg.Body))
		if e != nil {
			panic(e)

		}
		if Locate(dir + "/objects/" + object) {
			q.Send(msg.ReplyTo, url)
		}
	}
}
