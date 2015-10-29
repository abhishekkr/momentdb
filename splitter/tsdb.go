package momentdbSplitter

import (
	"regexp"

	"github.com/abhishekkr/goshare"
)

func doesNotMatch(txt int, patternTxt string) bool {
	pattern, patternError := regexp.Compile(patternTxt)
	if patternError != nil {
		return false
	}

	return pattern.Match([]byte(string(txt)))
}

/* packetSuitsKeyPatterrn check for key-type destination based on pattern. */
func packetSuitsTimeSeries(gosharePacket goshare.Packet, destination EngineDestination) bool {
	if destination.SplitterYear != "" && doesNotMatch(gosharePacket.TimeDot.Year, destination.SplitterYear) {
		return false
	}
	if destination.SplitterMonth != "" && doesNotMatch(gosharePacket.TimeDot.Month, destination.SplitterMonth) {
		return false
	}
	if destination.SplitterDay != "" && doesNotMatch(gosharePacket.TimeDot.Day, destination.SplitterDay) {
		return false
	}
	if destination.SplitterHour != "" && doesNotMatch(gosharePacket.TimeDot.Hour, destination.SplitterHour) {
		return false
	}
	if destination.SplitterMinute != "" && doesNotMatch(gosharePacket.TimeDot.Min, destination.SplitterMinute) {
		return false
	}
	if destination.SplitterSecond != "" && doesNotMatch(gosharePacket.TimeDot.Sec, destination.SplitterSecond) {
		return false
	}
	return true
}
