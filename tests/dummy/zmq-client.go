package main

import (
	"fmt"
	"time"

	golassert "github.com/abhishekkr/gol/golassert"
	golzmq "github.com/abhishekkr/gol/golzmq"
)

func ZmqRequest(ip string, ports []int) {
	_socket := golzmq.ZmqRequestSocket(ip, ports)

	ports_str := fmt.Sprintf("%v", ports)
	val, err := golzmq.ZmqRequest(_socket, "get from", ports_str)
	golassert.AssertEqual(err, nil)
	fmt.Println(val)
}

func main() {
	go ZmqRequest("127.0.0.1", []int{8887})
	go ZmqRequest("127.0.0.1", []int{8886})
	go ZmqRequest("127.0.0.1", []int{8885})
	time.Sleep(2 * time.Second)
	fmt.Println("passed not panic")
}
