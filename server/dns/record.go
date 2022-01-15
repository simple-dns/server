package dns

import (
	"github.com/leganck/simple-dns/config"
	"github.com/miekg/dns"
	"github.com/patrickmn/go-cache"
	"reflect"
	"time"
)

type Record struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
}

var records = cache.New(10*time.Second, 30*time.Second)

func (a Record) IsEmpty() bool {
	return reflect.DeepEqual(a, Record{})
}

func AllRecord() map[string]Record {
	result := make(map[string]Record)

	for key, item := range records.Items() {
		if !item.Expired() {
			result[key] = item.Object.(Record)
		}
	}
	return result
}

func AddRecord(record Record) {
	record.Name = dns.Fqdn(dns.Fqdn(record.Name) + config.GetDomainSuffix())
	records.SetDefault(record.Name, record)
}
