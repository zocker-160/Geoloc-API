package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zocker-160/Geoloc-API/impl"
)

const PORT = 9001
const VERSION = "0.3"

func handleRequest(
		rType string, db *sql.DB,
		w http.ResponseWriter, r *http.Request) (*impl.DBRow, error) {

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

	row, err := impl.FindEntryDB(string(body), db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return nil, fmt.Errorf(" -> %s", err)
	}

	fmt.Printf(" -> found in %s", time.Since(startTime))

	return row, nil
}


func handleCountryRequest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := handleRequest("country", db, w, r)
		if err != nil {
			fmt.Println(err)
			return
		}

		res := data.Country

		fmt.Fprint(w, res)
		fmt.Printf(" (%s) \n", res)
	}
}

func handleCoordsRequest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := handleRequest("coords", db, w, r)
		if err != nil {
			fmt.Println(err)
			return
		}

		res := fmt.Sprintf("%f,%f", data.Latitude, data.Longitude)

		fmt.Fprint(w, res)
		fmt.Printf(" (%s - %s) \n", res, data.Country)
	}
}

func handleAllJsonRequest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := handleRequest("all", db, w, r)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonencoder := json.NewEncoder(w)
		if err := jsonencoder.Encode(data); err != nil {
			fmt.Println(err)
			return
		}

		//w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json,charset=utf-8")

		fmt.Printf(" (all json - %s) \n", data.Country)
	}
}


func getDatabase() string {
	if len(os.Args) < 2 {
		log.Fatalln("Please specify path to ip-locations.sqlite")
	}

	return os.Args[1]
}

func main() {
	fmt.Printf("GeoLocAPI v%s \n", VERSION)

	filePath := getDatabase()
	db, err := sql.Open("sqlite", filePath)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	/* just for testing
	entry, err := impl.FindEntryDB("223.255.227.0", db)
	if err != nil {
		panic(err)
	}

	fmt.Println(entry)
	fmt.Println(entry.Country)

	return
	*/

	http.HandleFunc("/country", handleCountryRequest(db))
	http.HandleFunc("/coords", handleCoordsRequest(db))
	http.HandleFunc("/all/json", handleAllJsonRequest(db))

	fmt.Printf("Listening on port %d \n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}