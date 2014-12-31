package main

import (
	"fmt"
	"strings"
	"time"

	golzmq "github.com/abhishekkr/gol/golzmq"
)

func defaultF(messages []string) string {
	fmt.Println("get:", messages)
	return fmt.Sprintf("GET: %s", strings.Join(messages, " "))
}

func people(messages []string) string {
	fmt.Println("put:", messages)
	return fmt.Sprintf("people: %s", strings.Join(messages, " "))
}

func animal(messages []string) string {
	fmt.Println("post:", messages)
	return fmt.Sprintf("animal: %s", strings.Join(messages, " "))
}

func ZmqReply(ip string, ports []int, fn golzmq.RecieveArrayReturnString) {
	socket := golzmq.ZmqReplySocket(ip, ports)
	for {
		err := golzmq.ZmqReply(socket, fn)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	go ZmqReply("127.0.0.1", []int{8787, 8777, 8877}, defaultF)
	go ZmqReply("127.0.0.1", []int{8686, 8666, 8866}, people)
	go ZmqReply("127.0.0.1", []int{8585, 8555, 8855}, animal)
	time.Sleep(60 * time.Second)
}
