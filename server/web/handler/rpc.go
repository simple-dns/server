package handler

import (
	"fmt"
	"github.com/leganck/simple-dns/config"
	"github.com/leganck/simple-dns/server/dns"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RpcAddRecord(w http.ResponseWriter, r *http.Request) {
	if config.GetRpcSecret() == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if config.GetRpcSecret() != r.Header.Get("X-Dns-Token") {
		// not authorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	record := dns.Record{
		Ip:   r.PostFormValue("ip"),
		Name: r.PostFormValue("name"),
	}
	if !record.IsEmpty() {
		dns.AddRecord(record)
		log.Debugf("record: %s Added successfully", record.Name)
		result := make(map[string]interface{})
		result["success"] = true
		result["msg"] = fmt.Sprintf("%s Added successfully", record.Name)
		resultProcess(w, result)
	}
}
