package maxmind

import (
	"time"
)

const (
	mxmAsnFile = "GeoLite2-ASN-Blocks-IPv4.csv"
	//mxmCityFile    = "GeoLite2-City-Locations-en.csv"
	//mxmCountryFile = "GeoLite2-Country-Locations-en.csv"
)

// Init set interval in hours for database update.
// if 0 then no dynamic updates
func Init(interval uint) error {
	if interval == 0 {
		return parseCSV(mxmAsnFile)
	}

	for {
		select {
		case <-time.Tick(time.Duration(interval) * time.Hour):
			if err := parseCSV(mxmAsnFile); err != nil {
				return err
			}
		}
	}
	return nil
}
