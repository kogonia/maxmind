package maxmind

import (
	"net/netip"
	"strings"
)

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
