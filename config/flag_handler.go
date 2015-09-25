package momentdConfig

import (
	"flag"
	"fmt"

	"github.com/abhishekkr/gol/golconfig"
	"github.com/abhishekkr/goshare"
)

// flags
var (
	flagType     = flag.String("type", "", "type of momentdb system (store,splitter,...)")
	flatSplitter = flag.String("splitter", "", "the path to configure splitter logic")
)

/* getNodeType assigns type value based on values available and priority of node-type */
func getNodeType(config *(golconfig.FlatConfig)) {
	if (*config)["splitter"] != "" {
		(*config)["type"] = "splitter"
	} else {
		(*config)["type"] = "goshare"
		(*config) = mergeConfig((*config), goshare.ConfigFromFlags())
	}
}

/*
ConfigFromFlags configs from values provided to flags.
*/
func ConfigFromFlags() golconfig.FlatConfig {
	var config golconfig.FlatConfig
	config = make(golconfig.FlatConfig)

	flag.Parse()
	config["type"] = *flagType
	config["splitter"] = *flatSplitter

	getNodeType(&config)

	fmt.Printf("MomentDB base config:\n")
	for cfg, val := range config {
		fmt.Printf("[ %v : %v ]\n", cfg, val)
	}
	return config
}

/*
mergeConfig
*/
func mergeConfig(cfgs ...golconfig.FlatConfig) golconfig.FlatConfig {
	var config golconfig.FlatConfig
	config = make(golconfig.FlatConfig)

	for _, cfg := range cfgs {
		for k, v := range cfg {
			config[k] = v
		}
	}
	return config
}
