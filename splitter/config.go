package momentdbSplitter

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
)

/*
logic to help goshare key specific mode-type-* values to golzmq.LoadSplitter compatible values
*/

/*
EngineDestination is a structure to manage details of multiple momentdb-stores
or splitters balanced by this instance.
*/
type EngineDestination struct {
	DestinationIP    string `json:"destination_ip"`
	DestinationPorts []int  `json:"destination_ports"`
	SplitterMode     string `json:"mode"`
	SplitterType     string `json:"type"`
	SplitterPattern  string `json:"pattern"`
	SplitterYear     string `json:"year"`
	SplitterMonth    string `json:"month"`
	SplitterDay      string `json:"day"`
	SplitterHour     string `json:"hour"`
	SplitterMinute   string `json:"minute"`
	SplitterSecond   string `json:"second"`

	SourceChannel      chan []byte
	DestinationChannel chan []byte
	RequestPattern     *regexp.Regexp
}

/*
EngineDetail is a structure to manage details for this instance of momentdb
and EngineDestination it will balance.
*/
type EngineDetail struct {
	SourceIP            string              `json:"source_ip"`
	SourcePorts         []int               `json:"source_ports"`
	Destinations        []EngineDestination `json:"destinations"`
	DefaultDestinations []*EngineDestination
}

/*
EngineCollection is collection of EngineDetails.
*/
type EngineCollection struct {
	Engines []EngineDetail
}

/*
LoadEngineCollection unmarshall EngineCollection from Config file.
*/
func LoadEngineCollection(configPath string) (engines EngineCollection, err error) {
	jsonBytes, fileErr := ioutil.ReadFile(configPath)
	if fileErr != nil {
		err = fileErr
		return
	}
	err = json.Unmarshal(jsonBytes, &(engines.Engines))
	return
}
