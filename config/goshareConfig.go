package momentdConfig

import "github.com/abhishekkr/gol/golconfig"

func defaultGoshareConfig(config *golconfig.FlatConfig) {
	assignIfEmpty(config, "DBEngine", "leveldb")
	assignIfEmpty(config, "DBPath", "Moment.db")
	assignIfEmpty(config, "server-uri", "0.0.0.0")
	assignIfEmpty(config, "http-port", "10000")
	assignIfEmpty(config, "rep-ports", "11000,11001")
}
