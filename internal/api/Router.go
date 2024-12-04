package api

import (
	"net"
)

type HandlerFunc func(conn net.Conn, args []any)

type MiddlewareFunc func(next HandlerFunc) HandlerFunc

type Router struct {
	routes      map[string]HandlerFunc
	middlewares []MiddlewareFunc
}

func NewRouter() *Router {
	return &Router{
		routes:      make(map[string]HandlerFunc),
		middlewares: make([]MiddlewareFunc, 0),
	}
}

func (r *Router) AddRoute(command string, handler HandlerFunc) {
	r.routes[command] = r.applyMiddlewaresToHandler(handler)
}

func (r *Router) RegisterMiddleware(mw MiddlewareFunc) {
	r.middlewares = append(r.middlewares, mw)
}

func (r *Router) applyMiddlewaresToHandler(handler HandlerFunc) HandlerFunc {
	wrapped := handler
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		wrapped = r.middlewares[i](wrapped)
	}
	return wrapped
}
