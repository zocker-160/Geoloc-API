package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zocker-160/Geoloc-API/impl"
)



func setupEndpoints(ipEntries []*impl.IPEntry) {

	http.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		entry, err := handleRequest("country", w, r, ipEntries)
		if err != nil {
			fmt.Println(err)
			return
		}

		str := entry.Country

		fmt.Fprint(w, str)
		fmt.Printf(" (%s) \n", str)
	})

	http.HandleFunc("/coords", func(w http.ResponseWriter, r *http.Request) {
		entry, err := handleRequest("coords", w, r, ipEntries)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprint(w, entry.Location.ToStringTuple())
		fmt.Println()
	})
}

func handleRequest(
		rType string, 
		w http.ResponseWriter, r *http.Request, 
		ipEntries []*impl.IPEntry) (*impl.IPEntry, error) {
	
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

	entry, err := impl.FindEntry(string(body), ipEntries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return nil, fmt.Errorf(" -> %s", err)
	}

	fmt.Printf(" -> found in %s", time.Since(startTime))

	return entry, nil
}


const PORT = 9001
const VERSION = "0.2"


func getIPFile() string {
	if len(os.Args) < 2 {
		log.Fatalln("Please specify path to ip-locations.txt")
	}

	return os.Args[1]
}

func main() {
	fmt.Printf("GeoLocAPI v%s \n", VERSION)

	filePath := getIPFile()
	ipEntries := impl.ParseIPs(filePath)

	setupEndpoints(ipEntries)

	fmt.Printf("Listening on port %d \n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}