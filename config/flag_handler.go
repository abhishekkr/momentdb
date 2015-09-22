package momentdConfig

import (
	"flag"
	"fmt"

	"github.com/abhishekkr/gol/golconfig"
)

/* assignIfEmpty assigns val to *key only if it's empty */
func assignIfEmpty(config *golconfig.FlatConfig, key string, val string) {
	if *config[key] == "" {
		*config[key] = val
	}
}

/* getNodeType assigns type value based on values available and priority of node-type */
func getNodeType(config *golconfig.FlatConfig) {
	if *config["splitter"] != "" {
		*config["type"] = "splitter"
	} else {
		*config["type"] = "goshare"
	}
}

/*
ConfigFromFlags configs from values provided to flags.
*/
func ConfigFromFlags() golconfig.FlatConfig {
	var config golconfig.FlatConfig
	config = make(golconfig.FlatConfig)

	flag.Parse()
	flagConfig := flag.String("config", "", "the path to overriding config file")
	if *flagConfig != "" {
		configFile := golconfig.GetConfigurator("json")
		configFile.ConfigFromFile(*flagConfig, &config)
	}
	config["type"] = flag.String("type", "", "type of momentdb system (store,splitter,...)")
	config["splitter"] = flag.String("splitter", "", "the path to configure splitter logic")

	flagCPUProfile := flag.String("cpuprofile", "", "write cpu profile to file")
	assignIfEmpty(&config, "cpuprofile", *flagCPUProfile)

	getNodeType(&config)
	defaultGoshareConfig(&config)

	fmt.Printf("MomentDB base config:")
	for cfg, val := range config {
		fmt.Printf("[ %v : %v ]:", cfg, config["cfg"])
	}
	fmt.Println("\n\n+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	return config
}
