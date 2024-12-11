package api

import (
	"bytes"
	"log/slog"
	"net"
	"rediska/internal/service"
)

type API struct {
	log    *slog.Logger
	conn   net.Conn
	router *Router

	checker         Checker
	storageModifier StorageModifier
	valuesProvider  ValuesProvider
}

//go:generate go run github.com/vektra/mockery/v2@v2.50.0 --name=Checker
type Checker interface {
	PING() (bytes.Buffer, error)
	ECHO(string) (bytes.Buffer, error)
}

type StorageModifier interface {
	SET(key string, value string, opts ...service.SetOptionFunc) (bytes.Buffer, error)
}

type ValuesProvider interface {
	GET(key string) (bytes.Buffer, error)
}

func New(
	log *slog.Logger,
	checker Checker,
	storageModifier StorageModifier,
	valuesProvider ValuesProvider,
) *API {
	return &API{
		log:             log,
		checker:         checker,
		storageModifier: storageModifier,
		valuesProvider:  valuesProvider,
		router:          NewRouter(),
	}
}

func (a *API) setConn(conn net.Conn) {
	a.conn = conn
}

func (a *API) RegisterCommands() {
	a.router.AddRoute("PING", a.PING)
	a.router.AddRoute("ECHO", a.ECHO)
	a.router.AddRoute("SET", a.SET)
	a.router.AddRoute("GET", a.GET)
}
