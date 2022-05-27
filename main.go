package main

import (
	"fmt"
	"github.com/leganck/simple-dns/config"
	"github.com/leganck/simple-dns/server/dns"
	"github.com/leganck/simple-dns/server/web"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"runtime"
)

func init() {
	var logLevel log.Level
	if config.GetDebug() {
		logLevel = log.DebugLevel
	} else {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2000-01-01 00:00:00",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return fmt.Sprintf("%s:%d", frame.Function, frame.Line), ""
		},
	})
}
func main() {
	var g errgroup.Group
	g.Go(dns.Server)
	g.Go(web.Server)
	err := g.Wait()
	if err != nil {
		log.Fatal(err.Error())
	}
}
