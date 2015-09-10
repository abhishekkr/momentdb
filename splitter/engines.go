package momentdbSplitter

import "fmt"

/*
StartEngines will start all the goshare datastore engines based on config,
if any engine failed, a failover local will be created and assigned
*/
func StartEngines(configPath string) {
	usableEngines, _ := LoadEngineCollection(configPath)

	fmt.Println("Start all the engines configured...")
	for idx := range usableEngines.Engines {
		fmt.Println(usableEngines.Engines[idx])
		go ZmqSmartProxy(&(usableEngines.Engines[idx]))
	}
}

/*
StopEngines stop all started engines or their failovers if original failed.
*/
func StopEngines(engines EngineCollection) {
	fmt.Println("Stop all the started engines...")
}
