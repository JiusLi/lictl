package httpx

import (
	"fmt"
	"net/http"
)

const (
	get = iota
	post
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// 依据请求方式分组
type routeMap []handler

// url 对应 http handlerFunc
type handler map[string]http.HandlerFunc

// http handlerFunc
type HandlerFunc func(w http.ResponseWriter, req *http.Request)
type Middleware func(next HandlerFunc) HandlerFunc

// 初始化路由
func newRouteMap() routeMap {
	return []handler{
		get:  make(handler),
		post: make(handler),
	}
}

// 添加post分组请求
func (m routeMap) post(url string, fn http.HandlerFunc) {
	if _, ok := m[post][url]; ok {
		panic(fmt.Sprintf("path: %s is exit", url))
	}
	m[post][url] = fn
}

// 添加get分组请求
func (m routeMap) get(url string, fn http.HandlerFunc) {
	if _, ok := m[get][url]; ok {
		panic(fmt.Sprintf("path: %s is exit", url))
	}
	m[get][url] = fn
}
