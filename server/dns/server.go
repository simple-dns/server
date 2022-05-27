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
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true
	if record, ok := records.Get(q.Name); q.Qtype == dns.TypeA && q.Qclass == dns.ClassINET && ok {
		a := new(dns.A)
		a.Hdr = dns.RR_Header{Name: q.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 600}
		a.A = net.ParseIP(record.(Record).Ip)
		m.Answer = []dns.RR{a}
		info := fmt.Sprintf("Question: Type=%s Class=%s Name=%s Record=%s", dns.TypeToString[q.Qtype], dns.ClassToString[q.Qclass], q.Name, a.A)
		log.Debug(info)
	} else {
		if config.GetDnsConfig().ForwarderServer != "" {
			client := dns.Client{Net: "udp"}
			res, _, err := client.Exchange(r, config.GetDnsConfig().ForwarderServer)
			if err != nil {
				errorInfo := fmt.Sprintf("Question:  Name=%s errInfo=%s", q.Name, err)
				log.Error(errorInfo)
			} else {
				m = res
			}
		} else {
			m.Rcode = 2
		}
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
