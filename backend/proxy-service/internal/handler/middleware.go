package handler

import (
	"github.com/Enthreeka/proxy-service/pkg/logger"
	"net/http"
	"time"
)

// https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e
func MiddlewareLogger(log *logger.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Info("%s - %s (%v), proto: %s", r.Method, r.URL.Path, time.Since(startTime), r.Proto)
		h.ServeHTTP(w, r)
	})
}
