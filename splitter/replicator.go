package momentdb_splitter

import "github.com/abhishekkr/goshare"

/* for PUSH/DELETE the request will be send to all Nodes */
func sendToAll(destinations []*EngineDestination, request []byte) (reply []byte) {
	for idx, _ := range destinations {
		destinations[idx].DestinationChannel <- request
		reply = <-destinations[idx].SourceChannel
		break
	}
	return reply
}

/* for READ the destinations will be inquired in LB mode */
func sendBalancedMode(destinations []*EngineDestination, goshare_packet goshare.Packet, request []byte) (reply []byte) {
	for idx, _ := range destinations {
		/* data check will go to logs to be done; and a channel with pused destinations if matched kin'a to make logic for LB */
		destinations[idx].DestinationChannel <- request
		reply = <-destinations[idx].SourceChannel
		break
	}
	return reply
}

/* manage replicated actions for DBTasks */
func Replicate(destinations []*EngineDestination, goshare_packet goshare.Packet, request []byte) (reply []byte) {
	return sendBalancedMode(destinations, goshare_packet, request)

	/*to be fixed*/
	if goshare_packet.DBAction == "read" {
		return sendBalancedMode(destinations, goshare_packet, request)
	}
	return sendToAll(destinations, request)
}
