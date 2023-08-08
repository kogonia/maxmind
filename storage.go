package maxmind

import (
	"bytes"
	"encoding/json"
	"net/netip"
	"strconv"
	"strings"
	"sync"
)

type OrgInfo struct {
	OrgName string         `json:"org_name,omitempty"`
	ASN     string         `json:"asn,omitempty"`
	Prefix  []netip.Prefix `json:"prefix,omitempty"`
	//City     string `json:"city,omitempty"`
	//Country  string `json:"country,omitempty"`
}

func (oi OrgInfo) String() string {
	b := bytes.NewBuffer(make([]byte, 0, 64))
	_ = json.NewEncoder(b).Encode(oi)
	return strings.TrimSpace(string(b.Bytes()))
}

func (oi OrgInfo) Json() string {
	js, _ := json.MarshalIndent(oi, "", "\t")
	return string(js)
}

var storage = emptyStorage()
var mu = sync.Mutex{}

func emptyStorage() map[string]OrgInfo {
	return make(map[string]OrgInfo, 524288)
}

func copyStorageData() map[string]OrgInfo {
	tmpStorage := emptyStorage()
	for k, v := range storage {
		tmpStorage[k] = v
	}
	return tmpStorage
}

func (oi OrgInfo) save() {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := storage[oi.OrgName]; !ok {
		storage[oi.OrgName] = oi
	} else {
		org := storage[oi.OrgName]
		org.Prefix = append(storage[oi.OrgName].Prefix, oi.Prefix...)
		storage[oi.OrgName] = org
	}
}

func GetByIP(addr string) (OrgInfo, error) {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return OrgInfo{}, err
	}
	for _, org := range storage {
		for _, prefix := range org.Prefix {
			if prefix.Contains(ip) {
				return org, nil
			}
		}
	}
	return OrgInfo{}, nil
}

func GetByASN(asn string) (OrgInfo, error) {
	if strings.HasPrefix(strings.ToLower(asn), "as") {
		asn = strings.TrimPrefix(strings.ToLower(asn), "as")
	}
	if _, err := strconv.Atoi(asn); err != nil {
		return OrgInfo{}, errBadAsnProvided
	}
	for _, org := range storage {
		if org.ASN == asn {
			return org, nil
		}
	}
	return OrgInfo{}, nil
}

func GetByOrgName(orgName string) OrgInfo {
	if org, ok := storage[orgName]; ok {
		return org
	}
	return OrgInfo{}
}
