package maxmind

import (
	"errors"
	"net/netip"
	"strings"
)

var errFailedToParseCSV = errors.New("failed to parse csv file")
var errFailedToParsePrefixCSV = errors.New("failed to parse Prefix from csv")
var errBadAsnProvided = errors.New("bad asn format in request")

func parsePrefix(str string) (netip.Prefix, error) {
	p := strings.Split(str, "/")
	if len(p) != 2 {
		return netip.Prefix{}, errFailedToParsePrefixCSV
	}
	network := p[0]
	if len(strings.Split(p[0], ".")) == 3 {
		network += ".0"
	}
	network += "/" + p[1]

	prefix, err := netip.ParsePrefix(network)
	if err != nil {
		return netip.Prefix{}, err
	}
	return prefix, nil
}
