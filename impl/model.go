package impl

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


type Geolocation struct {
	Latitude, Longitude float32
}

func (c *Geolocation) ToStringTuple() string {
	return fmt.Sprintf("%f,%f", c.Latitude, c.Longitude)
}

////

type IPRange struct {
	StartIP, EndIP uint32
}

func (c *IPRange) IsInRange(ip uint32) bool {
	return c.StartIP <= ip && ip <= c.EndIP
}

////

type IPEntry struct {
	Range IPRange
	//CountryCode, Country, State, City string
	Country string
	Location Geolocation
}

func ParseIPs(filename string) []*IPEntry {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var ipData []*IPEntry
	//countryMap := make(map[string]int8)

	fmt.Printf("Started parsing: %s\n", filename)
	startTime := time.Now()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		ipe, err := ParseLine(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		ipData = append(ipData, ipe)
	}

	fmt.Println("Finished parsing")
	fmt.Println("number of entries:", len(ipData))
	fmt.Println("elapsed time:", time.Since(startTime))

	return ipData
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

	ipentry := &IPEntry{
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

	return ipentry, nil
}


func FindEntry(ip string, entries []*IPEntry) (*IPEntry, error) {
	if strings.HasPrefix(ip, "127.0.0") {
		return nil, errors.New("you have got to be kidding me!!! ")
	}

	ipDec, err := IpToDecimal(ip)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.Range.IsInRange(ipDec) {
			return entry, nil
		}
	}

	return nil, fmt.Errorf("%s not found in database", ip)
}

