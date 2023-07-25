package maxmind

import (
	"encoding/csv"
	"errors"
	"log"
	"net/netip"
	"os"
)

const delimiter = ','

var errFailedToParseCSV = errors.New("failed to parse csv file")
var errFailedToParsePrefixCSV = errors.New("failed to parse Prefix from csv")

func parseCSV(fileName string) (err error) {
	var f *os.File
	f, err = os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = delimiter
	data, err := reader.ReadAll()
	if err != nil {
		return
	}

	for i := range data {
		if len(data[i]) != 3 {
			return errFailedToParseCSV
		}

		if len(data[i][0]) > 0 && len(data[i][1]) > 0 && len(data[i][2]) > 0 {

			prefix, err := parsePrefix(data[i][0]) // Parse Prefix (network in csv)
			if err != nil {
				log.Printf("ERROR data: \"%s\" err: %v", data[i][0], err)
				continue
			}
			org := OrgInfo{
				OrgName: data[i][2],
				Prefix:  []netip.Prefix{prefix},
				ASN:     data[i][1],
			}
			org.save()
		}
	}
	return nil
}