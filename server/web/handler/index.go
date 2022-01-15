package handler

import (
	"encoding/json"
	"github.com/leganck/simple-dns/server/dns"
	"github.com/leganck/simple-dns/server/web/files"
	"html/template"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	temp, err := template.New("index").
		Funcs(template.FuncMap{"removeDot": removeDot}).
		ParseFS(files.Template, "template/*")
	if err != nil {
		panic(err)
	}

	if err = temp.ExecuteTemplate(w, "index.gohtml", dns.AllRecord()); err != nil {
		panic(err)
	}

}

func resultProcess(w http.ResponseWriter, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

func removeDot(domain string) string {
	return strings.TrimSuffix(domain, ".")
}
