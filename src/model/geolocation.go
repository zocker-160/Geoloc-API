package model

import "fmt"

type Geolocation struct {
	Latitude, Longitude float32
}

func (c *Geolocation) ToStringTuple() string {
	return fmt.Sprintf("%f,%f", c.Latitude, c.Longitude)
}