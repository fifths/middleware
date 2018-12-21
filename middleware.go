package middleware

import (
	"github.com/julienschmidt/httprouter"
)

type RoutesRequest interface {
	Use(RouteMiddleware)
	Handle(httprouter.Handle) httprouter.Handle
}

type RouteMiddleware func(httprouter.Handle) httprouter.Handle

type RoutesRequests struct {
	routeMiddleware []RouteMiddleware
}

func New() *RoutesRequests {
	return &RoutesRequests{}
}

func (s *RoutesRequests) Use(rmw RouteMiddleware) {
	s.routeMiddleware = append([]RouteMiddleware{rmw}, s.routeMiddleware...)
}

func (s *RoutesRequests) Handle(handle httprouter.Handle) httprouter.Handle {
	for _, v := range s.routeMiddleware {
		handle = v(handle)
	}
	return handle
}
