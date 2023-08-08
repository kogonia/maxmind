package maxmind

import (
	"encoding/csv"
	"net/netip"
	"os"
)

const delimiter = ','

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

	tmpStorage := copyStorageData()
	for i := range data {
		if len(data[i]) != 3 {
			return errFailedToParseCSV
		}

		if len(data[i][0]) > 0 && len(data[i][1]) > 0 && len(data[i][2]) > 0 {

			prefix, err := parsePrefix(data[i][0]) // Parse Prefix (network in csv)
			if err != nil {
				//log.Printf("ERROR data: \"%s\" err: %v", data[i][0], err)
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
	if len(storage) == 0 {
		storage = tmpStorage
	}
	return nil
}
