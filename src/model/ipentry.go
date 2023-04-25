package model

import (
	"fmt"
	"strconv"
	"strings"
)


type IPEntry struct {
	Range IPRange
	//CountryCode, Country, State, City string
	Country string
	Location Geolocation
}

func ParseLine(line string) (*IPEntry, error) {
	entries := strings.Split(line, ",")
	if (len(entries) < 8) {
		return nil, fmt.Errorf("invalid data input: %s", line)
	}

	var startIP, endIP uint64
	startIP, _ = strconv.ParseUint(entries[0], 10, 32)
	endIP, _ = strconv.ParseUint(entries[1], 10, 32)

	var lat, long float64
	lat, _ = strconv.ParseFloat(entries[len(entries) - 2], 32)
	long, _ = strconv.ParseFloat(entries[len(entries) - 1], 32)

	ipentry := IPEntry{
		Range: IPRange{
			StartIP: uint32(startIP), 
			EndIP: uint32(endIP),
		},
		//CountryCode: entries[2],
		Country: entries[3],
		//State: entries[4],
		//City: entries[len(entries)-3],
		Location: Geolocation{
			Latitude: float32(lat),
			Longitude: float32(long),
		},
	}

	return &ipentry, nil
}
