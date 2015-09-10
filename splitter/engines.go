package momentdbSplitter

import "fmt"

/*
displayDetail to show engine internal mapping to split ends in human terms.
*/
func displayDetail(engineIndex int, engine EngineDetail) {
	detail := fmt.Sprintf("[+] Engine#%d Detail\n", (engineIndex + 1))
	detail = fmt.Sprintf("%s |  Source: %s:%v\n\n", detail, engine.SourceIP, engine.SourcePorts)
	for destination_idx, destination := range engine.Destinations {
		detail = fmt.Sprintf("%s |[+]  Destination#%d: %s:%v\n", detail, (destination_idx + 1), destination.DestinationIP, destination.DestinationPorts)
		detail = fmt.Sprintf("%s | |     Split Mode: %v\n", detail, destination.SplitterMode)
		detail = fmt.Sprintf("%s | |     Split Pattern: %v\n", detail, destination.SplitterPattern)
		detail = fmt.Sprintf("%s | |     Split Type: %v\n", detail, destination.SplitterType)
		detail = fmt.Sprintf("%s | |     Request Pattern: %v\n\n", detail, destination.RequestPattern)
	}
	fmt.Println(detail)
}

/*
StartEngines will start all the goshare datastore engines based on config,
if any engine failed, a failover local will be created and assigned
*/
func StartEngines(configPath string) {
	usableEngines, _ := LoadEngineCollection(configPath)

	fmt.Println("Start all the engines configured...")
	for engineIndex := range usableEngines.Engines {
		displayDetail(engineIndex, usableEngines.Engines[engineIndex])
		go ZmqSmartProxy(&(usableEngines.Engines[engineIndex]))
	}
}

/*
StopEngines stop all started engines or their failovers if original failed.
*/
func StopEngines(engines EngineCollection) {
	fmt.Println("Stop all the started engines...")
}
