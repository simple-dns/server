package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				errorResult(w, fmt.Errorf("%v", err))
			}
		}()
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Debugf(
			"%s %s %s %s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			time.Since(start).String(),
		)
	})
}

func errorResult(w http.ResponseWriter, err error) {
	log.Error("error %v", err)
	w.Header().Set("Content-Type", "text/html")
	_, err = fmt.Fprint(w, err.Error())
	if err != nil {
		log.Error("error %v", err)
	}
	flusher := w.(http.Flusher)
	flusher.Flush()
}
