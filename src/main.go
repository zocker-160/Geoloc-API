package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/zocker-160/Geoloc-API/model"
)

const PORT = 9001
const VERSION = "0.2"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify path to ip-locations.txt")
		return
	}

	fmt.Printf("GeoLocAPI v%s \n", VERSION)

	filePath := os.Args[1]

	ipEntries := parseIPs(filePath)
	setupEndpoints(ipEntries)

	fmt.Printf("Listening on port %d \n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}


func parseIPs(filename string) []*model.IPEntry {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Printf("Started parsing: %s\n", filename)
	startTime := time.Now()

	var ipData []*model.IPEntry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		ipe, err := model.ParseLine(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}

		ipData = append(ipData, ipe)
	}

	fmt.Println("Finished parsing")
	fmt.Printf("number of entries: %d \n", len(ipData))
	fmt.Printf("elapsed time: %s \n", time.Since(startTime))

	return ipData
}

func setupEndpoints(ipEntries []*model.IPEntry) {
	http.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		entry, err := handleRequest("country", w, r, ipEntries)

		if err != nil {
			fmt.Println(err)
		} else {
			str := entry.Country

			fmt.Fprint(w, str)
			fmt.Printf(" (%s) \n", str)
		}
	})

	http.HandleFunc("/coords", func(w http.ResponseWriter, r *http.Request) {
		entry, err := handleRequest("coords", w, r, ipEntries)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Fprint(w, entry.Location.ToStringTuple())
		}
	})
}

func handleRequest(
	rType string, 
	w http.ResponseWriter, r *http.Request, 
	ipEntries []*model.IPEntry) (*model.IPEntry, error) {
	
	fmt.Printf(
		"[%s] got %s request", 
		time.Now().Format("2006.01.02 15:04:05 MST"), rType,
	)

	body, err := io.ReadAll(r.Body)
	if bLen := len(body); err != nil || bLen < 4 {
		http.Error(w, "invalid request", http.StatusBadRequest)

		return nil, fmt.Errorf(" -> invalid request (%s, %d)", err, bLen)
	}

	startTime := time.Now()

	entry, err := findEntry(string(body), ipEntries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return nil, fmt.Errorf(" -> %s", err)
	}

	fmt.Printf(" -> found in %s", time.Since(startTime))

	return entry, nil
}


func findEntry(ip string, entries []*model.IPEntry) (*model.IPEntry, error) {
	if strings.HasPrefix(ip, "127.0.0") {
		return nil, errors.New("you have got to be kidding me!!! ")
	}

	ipDec, err := ipToDecimal(ip)
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


func ipToDecimal(ip string) (uint32, error) {
	parsedIP := net.ParseIP(ip)
	ipBytes := parsedIP.To4()
	if ipBytes == nil {
		return 0, errors.New("invalid IP address")
	}

	ipDecimal := 
		uint32(ipBytes[0]) << 24 | 
		uint32(ipBytes[1]) << 16 |
		uint32(ipBytes[2]) << 8 | 
		uint32(ipBytes[3])

	return ipDecimal, nil
}