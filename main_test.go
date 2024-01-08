package main

import (
	"testing"
)

func Test(t *testing.T) {
	db, err := NewIPDB("ip-locations.txt", false)
	if err != nil {
		t.Errorf("Failed to load IPDB: %s", err.Error())
	}

	entry, err := db.FindEntry("223.255.227.0")
	if err != nil {
		t.Error("IP not found")
	}

	if entry.Country != "Indonesia" {
		t.Error("wrong country")
	}
}
