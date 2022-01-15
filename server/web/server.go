package web

import (
	"github.com/leganck/simple-dns/config"
	"github.com/leganck/simple-dns/server/web/route"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Server() error {
	httpConfig := config.GetHttpConfig()
	s := &http.Server{
		Addr:    ":" + httpConfig.HttpPort,
		Handler: route.NewRouter(),
	}
	log.Infof("Http Server Listening on %s", httpConfig.HttpPort)
	return s.ListenAndServe()
}
