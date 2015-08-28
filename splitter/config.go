package momentdb_splitter

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
)

/*
logic to help goshare key specific mode-type-* values to golzmq.LoadSplitter compatible values
*/

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

/* Unmarshall EngineCollection from Config file */
func LoadEngineCollection(config_path string) (engines EngineCollection, err error) {
	jsonBytes, fileErr := ioutil.ReadFile(config_path)
	if fileErr != nil {
		err = fileErr
		return
	}
	err = json.Unmarshal(jsonBytes, &(engines.Engines))
	return
}
