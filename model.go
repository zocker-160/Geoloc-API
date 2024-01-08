package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type IPEntry struct {
	Range IPRange
	//CountryCode, Country, State, City string
	Country string
	Location Geolocation
}

type IPRange struct {
	StartIP, EndIP uint32
}
func (c IPRange) IsInRange(ip uint32) bool {
	return c.StartIP <= ip && ip <= c.EndIP
}

type Geolocation struct {
	Latitude, Longitude float32
}
func (c Geolocation) ToStringTuple() string {
	return fmt.Sprintf("%f,%f", c.Latitude, c.Longitude)
}


type IPDB struct {
	directRead bool
	filename string
	fileLock sync.Mutex
	data []*IPEntry
}
func NewIPDB(filename string, preload bool) (*IPDB, error) {
	if preload {
		data, err := parseIPs(filename)
		if err != nil {
			return nil, err
		}

		return &IPDB{
			directRead: false,
			filename: filename,
			data: data,
		}, nil
	}

	if _, err := os.Stat(filename); err != nil {
		return nil, err
	}

	return &IPDB{
		directRead: true,
		filename: filename,
	}, nil
}
func (c *IPDB) FindEntry(ip string) (*IPEntry, error) {
	if strings.HasPrefix(ip, "127.0.0") {
		return nil, errors.New("localhost? You have got to be kidding me")
	}

	ipDec, err := ipToDecimal(ip)
	if err != nil {
		return nil, err
	}

	if c.directRead {
		// very slow, but low memory usage
		c.fileLock.Lock()
		defer c.fileLock.Unlock()

		file, err := os.Open(c.filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ipe, err := parseLine(scanner.Text())
			if err != nil {
				continue
			}

			if ipe.Range.IsInRange(ipDec) {
				return ipe, nil
			}
		}

	} else {
		for _, entry := range c.data {
			if entry.Range.IsInRange(ipDec) {
				return entry, nil
			}
		}
	}

	return nil, fmt.Errorf("%s not found in database", ip)
}

func parseIPs(filename string) ([]*IPEntry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ipData []*IPEntry

	fmt.Println("Preloading database:", filename)
	startTime := time.Now()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ipe, err := parseLine(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		ipData = append(ipData, ipe)
	}

	fmt.Println("Preload done:")
	fmt.Println("- number of entries:", len(ipData))
	fmt.Println("- elapsed time:", time.Since(startTime))

	return ipData, nil
}

func parseLine(line string) (*IPEntry, error) {
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
