package main

import (
	"fmt"
	"time"

	golassert "github.com/abhishekkr/gol/golassert"
	golzmq "github.com/abhishekkr/gol/golzmq"
)

func ZmqRequest(ip string, ports []int, request string) {
	_socket := golzmq.ZmqRequestSocket(ip, ports)
	val, err := golzmq.ZmqRequest(_socket, request)
	golassert.AssertEqual(err, nil)
	fmt.Println(val)
}

func main() {
	fmt.Println("defaultr")
	go ZmqRequest("127.0.0.1", []int{8887}, "push default myname anonymous")
	go ZmqRequest("127.0.0.1", []int{8886}, "read default myname")
	go ZmqRequest("127.0.0.1", []int{8885}, "delete default myname")
	fmt.Println("people")
	go ZmqRequest("127.0.0.1", []int{8887}, "push default people netizens")
	go ZmqRequest("127.0.0.1", []int{8886}, "read default people")
	go ZmqRequest("127.0.0.1", []int{8885}, "delete default people")
	fmt.Println("animal")
	go ZmqRequest("127.0.0.1", []int{8887}, "push default animal netizens")
	go ZmqRequest("127.0.0.1", []int{8886}, "read default animal")
	go ZmqRequest("127.0.0.1", []int{8885}, "delete default animal")
	time.Sleep(2 * time.Second)
	fmt.Println("passed not panic")
}
