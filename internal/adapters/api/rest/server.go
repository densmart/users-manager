package rest

import (
	"context"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type RestServer struct {
	server *http.Server
}

func NewRestServer(handler http.Handler) *RestServer {
	server := &http.Server{
		Addr:           ":" + viper.GetString("rest.port"),
		Handler:        handler,
		MaxHeaderBytes: viper.GetInt("rest.maxHeaderBytes"),
		ReadTimeout:    time.Duration(viper.GetInt("rest.readTimeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("rest.writeTimeout")) * time.Second,
	}
	return &RestServer{server: server}
}

func (h *RestServer) Run() error {
	return h.server.ListenAndServe()
}

func (h *RestServer) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
