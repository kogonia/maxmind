package maxmind

const (
	mxmAsnFile = "GeoLite2-ASN-Blocks-IPv4.csv"
	//mxmCityFile    = "GeoLite2-City-Locations-en.csv"
	//mxmCountryFile = "GeoLite2-Country-Locations-en.csv"
)

func Init() error {
	return parseCSV(mxmAsnFile)
}
