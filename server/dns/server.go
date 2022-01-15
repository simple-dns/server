package dns

import (
	"fmt"
	"github.com/leganck/simple-dns/config"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"net"
)

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	q := r.Question[0]
	info := fmt.Sprintf("Question: Type=%s Class=%s Name=%s", dns.TypeToString[q.Qtype], dns.ClassToString[q.Qclass], q.Name)
	m := new(dns.Msg)
	m.SetReply(r)
	if record, ok := records.Get(q.Name); q.Qtype == dns.TypeA && q.Qclass == dns.ClassINET && ok {
		a := new(dns.A)
		a.Hdr = dns.RR_Header{Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 600}
		a.A = net.ParseIP(record.(Record).Ip)
		m.Answer = []dns.RR{a}
		log.Debug(info)
	} else {
		m.Rcode = 2
	}
	w.WriteMsg(m)

}

func Server() error {
	dnsConfig := config.GetDnsConfig()
	server := &dns.Server{Addr: ":" + dnsConfig.DnsPort, Net: "udp"}
	server.Handler = dns.HandlerFunc(handleRequest)
	log.Infof("Dns Server Listening on %s", dnsConfig.DnsPort)
	return server.ListenAndServe()
}
