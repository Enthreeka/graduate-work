package http

import (
	"github.com/Enthreeka/elasticsearch-service/internal/delivery/http/handler"
	"github.com/Enthreeka/elasticsearch-service/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

type ServerOption struct {
	Addr string
}

func NewServer(log *logger.Logger, service Services, opts ServerOption) *http.Server {
	mux := chi.NewMux()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	mux.Use(middleware.RealIP,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
		handler.MiddlewareLogger(log),
	)

	mux.Mount("/", Router(log, service))

	return &http.Server{
		Addr:    opts.Addr,
		Handler: mux,
	}
}
