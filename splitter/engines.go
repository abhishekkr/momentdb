package momentdb_splitter

import "fmt"

/*
will start all the goshare datastore engines based on config
if any engine failed, a failover local will be created and assigned
*/
func StartEngines(config_path string) {
	usable_engines, _ := LoadEngineCollection(config_path)

	fmt.Println("Start all the engines configured...")
	for _, engine := range usable_engines.Engines {
		fmt.Println(engine)
		go ZmqSmartProxy(&engine)
	}
}

/*
	stop all started engines or their failovers if original failed
*/
func StopEngines(engines EngineCollection) {
	fmt.Println("Stop all the started engines...")
}
