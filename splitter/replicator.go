package momentdbSplitter

import "github.com/abhishekkr/goshare"

/* sendToAll for PUSH/DELETE the request will be send to all Nodes */
func sendToAll(destinations []*EngineDestination, request []byte) (reply []byte) {
	for idx := range destinations {
		destinations[idx].DestinationChannel <- request
		reply = <-destinations[idx].SourceChannel
		break
	}
	return reply
}

/* sendBalancedMode for READ the destinations will be inquired in LB mode */
func sendBalancedMode(destinations []*EngineDestination, request []byte) (reply []byte) {
	for idx := range destinations {
		/* data check will go to logs to be done; and a channel with pused destinations if matched kin'a to make logic for LB */
		destinations[idx].DestinationChannel <- request
		reply = <-destinations[idx].SourceChannel
		break
	}
	return reply
}

/*
Replicate manage replicated actions for DBTasks.
*/
func Replicate(destinations []*EngineDestination, gosharePacket goshare.Packet, request []byte) (reply []byte) {
	return sendBalancedMode(destinations, request)

	/*to be fixed*/
	if gosharePacket.DBAction == "read" {
		return sendBalancedMode(destinations, request)
	}
	return sendToAll(destinations, request)
}
