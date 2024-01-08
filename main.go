package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const DEFAULT_PORT = 9001
const VERSION = "0.4"

func handleRequest(rType string,
		w http.ResponseWriter, r *http.Request) (*IPEntry, error) {

	fmt.Printf(
		"[%s] request %s",
		time.Now().Format("2006.01.02 15:04:05 MST"), rType,
	)

	body, err := io.ReadAll(r.Body)
	if bLen := len(body); err != nil || bLen < 4 {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return nil, fmt.Errorf(" -> invalid request (%s, %d)", err, bLen)
	}

	startTime := time.Now()

	entry, err := ipDB.FindEntry(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return nil, fmt.Errorf(" -> %s", err)
	}

	fmt.Printf(" -> found in %s", time.Since(startTime))

	return entry, nil
}

func handleCountryRequest(w http.ResponseWriter, r *http.Request) {
	entry, err := handleRequest("/country", w, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	res := entry.Country

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)

	fmt.Printf(" (%s) \n", res)
}

func handleCoordsRequest(w http.ResponseWriter, r *http.Request) {
	entry, err := handleRequest("/coords", w, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	res := entry.Location.ToStringTuple()

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)

	fmt.Printf(" (%s) \n", res)
}

func handleAllJsonRequest(w http.ResponseWriter, r *http.Request) {
	entry, err := handleRequest("/all/json", w, r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(entry); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json,charset=utf-8")

	fmt.Printf(" (all json - %s) \n", entry.Country)
}


var ipDB *IPDB

func main() {
	fmt.Printf("GeoLocAPI v%s \n###\n", VERSION)

	if len(os.Args) < 2 {
		log.Fatalln("Please specify path to ip-locations.txt")
	}

	filePath := os.Args[1]
	ramopt := os.Getenv("GEOAPI_RAM_OPT") == "1"

	port := DEFAULT_PORT
	if p, err := strconv.Atoi(os.Getenv("GEOAPI_PORT")); err == nil {
		port = p
	}

	fmt.Println("input file:", filePath)
	fmt.Println("RAM_OPT:", ramopt)
	fmt.Println("PORT:", port)
	fmt.Println("###")

	db, err := NewIPDB(filePath, !ramopt)
	if err != nil {
		log.Fatalln("Failed to load IPDB:", err)
	}
	ipDB = db

	http.HandleFunc("/country", handleCountryRequest)
	http.HandleFunc("/coords", handleCoordsRequest)
	http.HandleFunc("/all/json", handleAllJsonRequest)

	fmt.Printf("Listening on port %d \n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
