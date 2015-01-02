package momentdb_splitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	golzmq "github.com/abhishekkr/gol/golzmq"
	"github.com/abhishekkr/goshare"
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
	RequestPattern     *regexp.Regexp
}

type EngineDetail struct {
	SourceIP            string              `json:"source_ip"`
	SourcePorts         []int               `json:"source_ports"`
	Destinations        []EngineDestination `json:"destinations"`
	DefaultDestinations []*EngineDestination
}

type EngineCollection struct {
	Engines []EngineDetail
}

/* Create a Proxy connection for given ZmqProxyConfig */
func ZmqSmartProxy(engine *EngineDetail) {
	default_len := 0
	for idx, _ := range engine.Destinations {
		if engine.Destinations[idx].SplitterMode == "default" {
			default_len += 1
		}
	}
	engine.DefaultDestinations = make([]*EngineDestination, default_len)
	default_idx := 0
	for idx, _ := range engine.Destinations {
		if engine.Destinations[idx].SplitterMode == "default" {
			engine.DefaultDestinations[default_idx] = &(engine.Destinations[idx])
			engine.Destinations[idx].SplitterPattern = ".?"
			default_idx += 1
		}
		pattern, pattern_err := regexp.Compile(engine.Destinations[idx].SplitterPattern)
		if pattern_err != nil {
			fmt.Println("ERROR: Compilation of provided Regexp failed.", pattern_err)
			continue
		}

		engine.Destinations[idx].SourceChannel = make(chan []byte)
		engine.Destinations[idx].DestinationChannel = make(chan []byte)
		engine.Destinations[idx].RequestPattern = pattern
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
		return channelForRequest(engine, request)
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

/* check for key-type destination based on pattern */
func packetSuitsKeyPatterrn(goshare_packet goshare.Packet, destination EngineDestination) bool {
	switch goshare_packet.DBAction {
	case "push":
		for keyname, _ := range goshare_packet.HashMap {
			if destination.RequestPattern.Match([]byte(keyname)) {
				return true
			}
		}
	case "read", "delete":
		if destination.RequestPattern.Match([]byte(goshare_packet.KeyList[0])) {
			return true
		}
	}
	return false
}

/* check for destination based on SplitterType */
func packetSuitsDestination(goshare_packet goshare.Packet, destination EngineDestination) bool {
	if destination.SplitterMode == "partial" && destination.SplitterType == "key" {
		return packetSuitsKeyPatterrn(goshare_packet, destination)
	}
	return false
}

/* check if request suits destination */
func requestSuitsDestination(request []byte, goshare_packet goshare.Packet, destination EngineDestination) bool {
	if goshare_packet.DBAction == "ERROR" && destination.RequestPattern.Match(request) {
		return true
	}
	return packetSuitsDestination(goshare_packet, destination)
}

/* main split logic, makes call for replication when required */
func channelForRequest(engine *EngineDetail, request []byte) (reply []byte) {
	destinations := make([]*EngineDestination, len(engine.Destinations))
	destination_idx := 0

	request_fields := strings.Fields(string(request))
	goshare_packet := goshare.CreatePacket(request_fields)

	for idx, _ := range engine.Destinations {
		if requestSuitsDestination(request, goshare_packet, engine.Destinations[idx]) {
			destinations[destination_idx] = &(engine.Destinations[idx])
			destination_idx += 1
		}
	}
	if destinations[0] == nil {
		for idx, _ := range engine.DefaultDestinations {
			destinations[destination_idx] = engine.DefaultDestinations[idx]
			destination_idx += 1
		}
	}
	if destinations[0] == nil {
		fmt.Println(engine.Destinations[0])
		fmt.Println(len(engine.DefaultDestinations))
		fmt.Println("ERROR: no destinations found for this request")
		return reply
	}

	reply = Replicate(destinations, goshare_packet, request)
	return reply
}
