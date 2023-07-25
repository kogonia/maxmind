package maxmind

import "fmt"

const (
	mxmAsnFile = "GeoLite2-ASN-Blocks-IPv4.csv"
	//mxmCityFile    = "GeoLite2-City-Locations-en.csv"
	//mxmCountryFile = "GeoLite2-Country-Locations-en.csv"
)

func Init() error {
	parseCSV(mxmAsnFile)
	for _, oi := range storage {
		fmt.Println(oi.String())
	}
	return nil
}
