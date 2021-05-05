package models

type (
	PortInfo struct {
		Symbol     string
		Name       string
		Country    string
		City       string
		Province   string
		Alias      []string
		Regions    []string
		Timezones  []string
		Unlocks    []string
		Code       string
		Coordinate *Coordinate
	}

	Coordinate struct {
		Latitude  float32
		Longitude float32
	}
)
