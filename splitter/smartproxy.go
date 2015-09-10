package momentdbSplitter

import (
	"fmt"
	"regexp"
	"strings"

	golzmq "github.com/abhishekkr/gol/golzmq"
	"github.com/abhishekkr/goshare"
)

/*
ZmqSmartProxy create a Proxy connection for given ZmqProxyConfig.
*/
func ZmqSmartProxy(engine *EngineDetail) {
	defaultLength := 0
	for idx := range engine.Destinations {
		if engine.Destinations[idx].SplitterMode == "default" {
			defaultLength++
		}
	}
	engine.DefaultDestinations = make([]*EngineDestination, defaultLength)
	defaultIndex := 0
	for idx := range engine.Destinations {
		if engine.Destinations[idx].SplitterMode == "default" {
			engine.DefaultDestinations[defaultIndex] = &(engine.Destinations[idx])
			engine.Destinations[idx].SplitterPattern = ".?"
			defaultIndex++
		}
		pattern, patternError := regexp.Compile(engine.Destinations[idx].SplitterPattern)
		if patternError != nil {
			fmt.Println("ERROR: Compilation of provided Regexp failed.", patternError)
			continue
		}

		engine.Destinations[idx].SourceChannel = make(chan []byte)
		engine.Destinations[idx].DestinationChannel = make(chan []byte)
		engine.Destinations[idx].RequestPattern = pattern
		go proxyDestination(&(engine.Destinations[idx]))
	}
	go proxySource(engine)
}

/* proxyDestination create a ZMQ Proxy Reader from source of Proxy. */
func proxyDestination(destination *EngineDestination) error {
	socket := golzmq.ZmqRequestSocket(destination.DestinationIP, destination.DestinationPorts)

	for {
		request := <-destination.DestinationChannel
		reply, errorRequest := golzmq.ZmqRequestByte(socket, request)
		if errorRequest != nil {
			fmt.Println("ERROR:", errorRequest)
			return errorRequest
		}
		destination.SourceChannel <- reply
	}
}

/* proxySource create a ZMQ Proxy Reader from source of Proxy */
func proxySource(engine *EngineDetail) error {
	socket := golzmq.ZmqReplySocket(engine.SourceIP, engine.SourcePorts)

	replyHandler := func(request []byte) []byte {
		return channelForRequest(engine, request)
	}

	for {
		errorReply := golzmq.ZmqReplyByte(socket, replyHandler)
		if errorReply != nil {
			fmt.Println("ERROR:", errorReply)
			return errorReply
		}
	}
	return nil
}

/* packetSuitsKeyPatterrn check for key-type destination based on pattern. */
func packetSuitsKeyPatterrn(gosharePacket goshare.Packet, destination EngineDestination) bool {
	switch gosharePacket.DBAction {
	case "push":
		for keyname := range gosharePacket.HashMap {
			if destination.RequestPattern.Match([]byte(keyname)) {
				return true
			}
		}
	case "read", "delete":
		if destination.RequestPattern.Match([]byte(gosharePacket.KeyList[0])) {
			return true
		}
	}
	return false
}

/* packetSuitsDestination check for destination based on SplitterType. */
func packetSuitsDestination(gosharePacket goshare.Packet, destination EngineDestination) bool {
	if destination.SplitterMode == "partial" && destination.SplitterType == "key" {
		return packetSuitsKeyPatterrn(gosharePacket, destination)
	}
	return false
}

/* requestSuitsDestination check if request suits destination. */
func requestSuitsDestination(request []byte, gosharePacket goshare.Packet, destination EngineDestination) bool {
	if gosharePacket.DBAction == "ERROR" && destination.RequestPattern.Match(request) {
		return true
	}
	return packetSuitsDestination(gosharePacket, destination)
}

/* channelForRequest is the main split logic, makes call for replication when required */
func channelForRequest(engine *EngineDetail, request []byte) (reply []byte) {
	destinations := make([]*EngineDestination, len(engine.Destinations))
	destinationIndex := 0

	requestFields := strings.Fields(string(request))
	gosharePacket := goshare.CreatePacket(requestFields)

	for idx := range engine.Destinations {
		if requestSuitsDestination(request, gosharePacket, engine.Destinations[idx]) {
			destinations[destinationIndex] = &(engine.Destinations[idx])
			destinationIndex++
		}
	}
	if destinations[0] == nil {
		for idx := range engine.DefaultDestinations {
			destinations[destinationIndex] = engine.DefaultDestinations[idx]
			destinationIndex++
		}
	}
	if destinations[0] == nil {
		fmt.Println(engine.Destinations[0])
		fmt.Println(len(engine.DefaultDestinations))
		fmt.Println("ERROR: no destinations found for this request")
		return reply
	}

	reply = Replicate(destinations, gosharePacket, request)
	return reply
}
