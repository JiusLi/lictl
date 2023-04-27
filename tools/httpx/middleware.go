package httpx

import (
	"net/http"
)

// 中间件

type Mi func(handlerFunc http.HandlerFunc) http.HandlerFunc

func middleware(ms Mi, rs ...Route) []Route {
	routes := make([]Route, len(rs))

	for i := range rs {
		route := rs[i]
		routes[i] = Route{
			Method:      route.Method,
			Path:        route.Path,
			HandlerFunc: ms(route.HandlerFunc),
		}
	}
	return routes
}

// 添加中间件
func Middlewares(ms []Mi, rs ...Route) []Route {
	for i := len(ms) - 1; i >= 0; i-- {
		rs = middleware(ms[i], rs...)
	}
	return rs

}
