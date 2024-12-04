package api

import (
	"rediska/internal/Commands"
	"strings"
	"time"
)

func (a *API) PING(args []any) error {
	a.log.Debug("ponging...")
	//err := Commands.PING(a.conn)
	//TODO: return error from service when it will be implemented
	Commands.PING(a.conn)

	return nil
}
func (a *API) ECHO(args []any) error {
	a.log.Debug("echoing...")
	phrase := args[0].(string)
	Commands.ECHO(a.conn, phrase)

	//err := Commands.ECHO(a.conn)
	//TODO: return error from service  when it will be implemented

	return nil
}
func (a *API) SET(args []any) error {
	a.log.Debug("setting...")
	key := args[0].(string)
	value := args[1].(string)
	ttlType := args[2].(string)
	expire := args[3].(string)

	var (
		ttl time.Duration
		err error
	)
	if strings.EqualFold(ttlType, "px") {
		ttl, err = time.ParseDuration(expire + "ms")
		if err != nil {
			a.log.Error("failed to parse duration")
			return err
		}
		Commands.SET(a.conn, key, value, Commands.WithTTL(ttl))
		//err = Commands.SET(a.conn, key, value, Commands.WithTTL(ttl))
	}
	Commands.SET(a.conn, key, value)
	//err = Commands.SET(a.conn, key, value)
	//TODO: return error from service  when it will be implemented
	return nil
}

func (a *API) GET(args []any) error {
	a.log.Debug("getting...")
	key := args[0].(string)
	Commands.GET(a.conn, key)

	//err := Commands.GET(a.conn, key)
	//TODO: return error from service when it will be implemented

	return nil
}
func (a *API) RegisterCommands() {
	a.router.AddRoute("PING", a.PING)
	a.router.AddRoute("ECHO", a.ECHO)
	a.router.AddRoute("SET", a.SET)
	a.router.AddRoute("GET", a.GET)
}
