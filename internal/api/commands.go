package api

import (
	"bytes"
	"rediska/internal/domain/response"
	"rediska/internal/lib/logger/sl"
	"rediska/internal/service"
	"strings"
	"time"
)

func (a *API) PING(args []any) {
	log := a.log.With("op", "api.PING")
	log.Debug("ponging...")
	buf, err := a.checker.PING()
	a.sendResponse(buf, err)
}
func (a *API) ECHO(args []any) {
	log := a.log.With("op", "Api.Echo")

	log.Debug("echoing...")
	phrase := args[0].(string)
	buf, err := a.checker.ECHO(phrase)
	a.sendResponse(buf, err)
}
func (a *API) SET(args []any) {
	log := a.log.With("op", "api.SET")

	log.Debug("setting...")
	key := args[0].(string)
	value := args[1].(string)
	ttlType := args[2].(string)
	expire := args[3].(string)

	var (
		ttl time.Duration
		buf bytes.Buffer
		err error
	)
	if strings.EqualFold(ttlType, "px") {
		ttl, err = time.ParseDuration(expire + "ms")
		if err != nil {
			log.Error("failed to parse duration", sl.Err(err))
			a.sendResponse(bytes.Buffer{}, err)
			return
		}
		buf, err = a.storageModifier.SET(key, value, service.WithTTL(ttl))
		a.sendResponse(buf, err)
		return
	}
	buf, err = a.storageModifier.SET(key, value)
	a.sendResponse(buf, err)
	return
}

func (a *API) GET(args []any) {
	log := a.log.With("op", "api.GET")
	log.Debug("getting...")

	key := args[0].(string)
	buf, err := a.valuesProvider.GET(key)
	a.sendResponse(buf, err)
}

func (a *API) sendResponse(buf bytes.Buffer, err error) {
	log := a.log.With("op", "api.sendResponse")

	if err != nil {
		log.Debug("sending Error")
		buf = response.CreateError(err)
	}
	_, err = a.conn.Write(buf.Bytes())
	if err != nil {
		log.Error("failed to write response to client")
	}
}
