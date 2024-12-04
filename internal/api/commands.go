package api

import (
	"fmt"
	"net"
	"rediska/internal/Commands"
	"strings"
	"time"
)

func (a *API) RegisterCommands() {
	ping := func(conn net.Conn, args []any) {
		fmt.Println("ponging...")
		Commands.PING(a.conn)
	}
	echo := func(conn net.Conn, args []any) {
		fmt.Println("echoing...")
		phrase := args[0].(string)
		Commands.ECHO(a.conn, phrase)
	}
	set := func(conn net.Conn, args []any) {
		fmt.Println("setting...")
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
				fmt.Println("failed to parse duration")
			}
			Commands.SET(a.conn, key, value, Commands.WithTTL(ttl))
			return
		}
		Commands.SET(a.conn, key, value)
	}
	get := func(conn net.Conn, args []any) {
		fmt.Println("getting...")
		key := args[0].(string)
		Commands.GET(a.conn, key)
	}

	a.router.AddRoute("PING", ping)
	a.router.AddRoute("ECHO", echo)
	a.router.AddRoute("SET", set)
	a.router.AddRoute("GET", get)
}
