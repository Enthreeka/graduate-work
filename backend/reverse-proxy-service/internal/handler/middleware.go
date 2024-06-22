package handler

import (
	"github.com/Enthreeka/reverse-proxy-service/pkg/logger"
	"github.com/Enthreeka/reverse-proxy-service/pkg/metric"
	"net/http"
	"strconv"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

// https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e
func MiddlewareLogger(log *logger.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		h.ServeHTTP(rec, r)

		duration := time.Since(start)
		log.Info("%s - %s (%v), proto: %s, status: %d", r.Method, r.URL.Path, duration, r.Proto, rec.status)

		metric.SummeryTimes.WithLabelValues(strconv.Itoa(rec.status), r.Method, r.URL.Path).Observe(duration.Seconds())
		metric.IncRequest.WithLabelValues(strconv.Itoa(rec.status), r.URL.Path).Add(1)
	})
}

func EnableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h.ServeHTTP(w, r)
	})
}
