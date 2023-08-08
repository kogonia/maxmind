package maxmind

import (
	"net/netip"
	"strconv"
	"strings"
	"sync"
)

type storage struct {
	sync.Mutex
	data map[string]*OrgInfo
}

type OrgInfo struct {
	OrgName string         `json:"org_name,omitempty"`
	ASN     string         `json:"asn,omitempty"`
	Prefix  []netip.Prefix `json:"prefix,omitempty"`
	//City     string `json:"city,omitempty"`
	//Country  string `json:"country,omitempty"`
}

var st = emptyStorage()

func emptyStorage() *storage {
	return &storage{data: make(map[string]*OrgInfo, 524288)}
}

func copyStorageData() *storage {
	if len(st.data) == 0 {
		return emptyStorage()
	}
	tmpStorage := emptyStorage()
	for k, v := range st.data {
		tmpStorage.data[k] = v
	}
	return tmpStorage
}

func (s *storage) save(oi *OrgInfo) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.data[oi.OrgName]; !ok {
		s.data[oi.OrgName] = oi
	} else {
		org := s.data[oi.OrgName]
		org.Prefix = append(s.data[oi.OrgName].Prefix, oi.Prefix...)
		s.data[oi.OrgName] = org
	}
}

func GetByIP(addr string) (*OrgInfo, error) {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return nil, err
	}
	st.Lock()
	defer st.Unlock()
	for _, org := range st.data {
		for _, prefix := range org.Prefix {
			if prefix.Contains(ip) {
				return org, nil
			}
		}
	}
	return nil, nil
}

func GetByASN(asn string) (*OrgInfo, error) {
	if strings.HasPrefix(strings.ToLower(asn), "as") {
		asn = strings.TrimPrefix(strings.ToLower(asn), "as")
	}
	if _, err := strconv.Atoi(asn); err != nil {
		return nil, errBadAsnProvided
	}
	for _, org := range st.data {
		if org.ASN == asn {
			return org, nil
		}
	}
	return nil, nil
}

func GetByOrgName(orgName string) *OrgInfo {
	if org, ok := st.data[orgName]; ok {
		return org
	}
	return nil
}
