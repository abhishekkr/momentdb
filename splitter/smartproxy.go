package momentdb_splitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	golzmq "github.com/abhishekkr/gol/golzmq"
)

func LoadEngineCollection(config_path string) (engines EngineCollection, err error) {
	jsonBytes, fileErr := ioutil.ReadFile(config_path)
	if fileErr != nil {
		err = fileErr
		return
	}
	err = json.Unmarshal(jsonBytes, &(engines.Engines))
	return
}

type EngineDestination struct {
	DestinationIP      string `json:"destination_ip"`
	DestinationPorts   []int  `json:"destination_ports"`
	SplitterMode       string `json:"mode"`
	SplitterType       string `json:"type"`
	SplitterPattern    string `json:"pattern"`
	SourceChannel      chan []byte
	DestinationChannel chan []byte
}

type EngineDetail struct {
	SourceIP     string              `json:"source_ip"`
	SourcePorts  []int               `json:"source_ports"`
	Destinations []EngineDestination `json:"destinations"`
}

type EngineCollection struct {
	Engines []EngineDetail
}

/* Create a Proxy connection for given ZmqProxyConfig */
func ZmqSmartProxy(engine *EngineDetail) {

	for idx, _ := range engine.Destinations {
		engine.Destinations[idx].SourceChannel = make(chan []byte)
		engine.Destinations[idx].DestinationChannel = make(chan []byte)
		go proxyDestination(&(engine.Destinations[idx]))
	}
	go proxySource(engine)
}

/* Create a ZMQ Proxy Reader from source of Proxy */
func proxyDestination(destination *EngineDestination) error {
	socket := golzmq.ZmqRequestSocket(destination.DestinationIP, destination.DestinationPorts)

	for {
		request := <-destination.DestinationChannel
		reply, err_request := golzmq.ZmqRequestByte(socket, request)
		if err_request != nil {
			fmt.Println("ERROR:", err_request)
			return err_request
		}
		destination.SourceChannel <- reply
	}
}

/* Create a ZMQ Proxy Reader from source of Proxy */
func proxySource(engine *EngineDetail) error {
	socket := golzmq.ZmqReplySocket(engine.SourceIP, engine.SourcePorts)

	reply_handler := func(request []byte) []byte {
		return ChannelForRequest(engine, request)
	}

	for {
		err_reply := golzmq.ZmqReplyByte(socket, reply_handler)
		if err_reply != nil {
			fmt.Println("ERROR:", err_reply)
			return err_reply
		}
	}
	return nil
}

/* main split logic */
func ChannelForRequest(engine *EngineDetail, request []byte) []byte {
	var reply []byte
	for _, destination := range engine.Destinations {
		if destination.SplitterMode == "default" {
			fmt.Println("need more logic here")
			destination.DestinationChannel <- request
			reply = <-destination.SourceChannel
			break
		}
	}
	fmt.Println("when no handler found send back default")
	return reply
}
