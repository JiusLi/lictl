package e3

import (
	"net/http"
)

// 实现IOdServer的接口，以及http提供ServeHttp方法
type OdServer struct {
	router MethodMaps
}

type IOdServer interface {
	GET(url string, f HandlerFunc)
	POST(url string, f HandlerFunc)
	PUT(url string, f HandlerFunc)
	DELETE(url string, f HandlerFunc)
}

type HandlerMapped struct {
	f HandlerFunc
}

// 接口函数单位，即我们编写代码逻辑的函数
type HandlerFunc func(w http.ResponseWriter, req *http.Request)

func Default() *OdServer {
	return &OdServer{
		router: NewRouter(),
	}
}

// 实现Handler接口，匹配方法以及路径
func (o *OdServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//转发给doHandler进行执行
	o.doHandler(w, req)
}

// 判断需要执行的Http Method，从而查找对应的接口并且执行
func (o *OdServer) doHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		{
			if hm, ok := o.router.GetMapping(req.URL.RequestURI()); ok {
				hm.f(w, req)
			}
		}
	case http.MethodPost:
		{
			if hm, ok := o.router.PostMapping(req.URL.RequestURI()); ok {
				hm.f(w, req)
			}

		}
	case http.MethodDelete:
		{
			if hm, ok := o.router.DeleteMapping(req.URL.String()); ok {
				hm.f(w, req)
			}
		}
	case http.MethodPut:
		{
			if hm, ok := o.router.PutMapping(req.URL.String()); ok {
				hm.f(w, req)
			}
		}
	default:
		{

		}
	}
}

func (o *OdServer) GET(url string, f HandlerFunc) {
	o.router.GetAdd(url, HandlerMapped{f: f})
}
func (o *OdServer) POST(url string, f HandlerFunc) {
	o.router.PostAdd(url, HandlerMapped{f: f})
}
func (o *OdServer) PUT(url string, f HandlerFunc) {
	o.router.PutAdd(url, HandlerMapped{f: f})
}
func (o *OdServer) DELETE(url string, f HandlerFunc) {
	o.router.DeleteAdd(url, HandlerMapped{f: f})
}

/*
*
提供基本的路由功能，添加路由，查找路由
*/
const (
	GET = iota
	POST
	PUT
	DELETE
	CONNECTIBNG
	HEAD
	OPTIONS
	PATCH
	TRACE
)

func NewRouter() MethodMaps {
	return []handler{
		GET:    make(handler),
		POST:   make(handler),
		PUT:    make(handler),
		DELETE: make(handler),
	}
}

type MethodMaps []handler
type handler map[string]HandlerMapped

// 映射路由，获取Get方法下对应的接口
func (m MethodMaps) GetMapping(url string) (HandlerMapped, bool) {
	if hm, ok := m[GET][url]; ok {
		return hm, true
	}
	return HandlerMapped{}, false
}

// 映射路由，获取Post方法下对应的接口
func (m MethodMaps) PostMapping(url string) (HandlerMapped, bool) {
	if hm, ok := m[POST][url]; ok {
		return hm, true
	}
	return HandlerMapped{}, false
}

// 映射路由，获取Delete方法下对应的接口
func (m MethodMaps) DeleteMapping(url string) (HandlerMapped, bool) {
	if hm, ok := m[DELETE][url]; ok {
		return hm, true
	}
	return HandlerMapped{}, false
}

// 映射路由，获取Put方法下对应的接口
func (m MethodMaps) PutMapping(url string) (HandlerMapped, bool) {
	if hm, ok := m[PUT][url]; ok {
		return hm, true
	}
	return HandlerMapped{}, false
}

// 增加Get方法下的接口
func (m MethodMaps) GetAdd(url string, mapped HandlerMapped) {
	if _, ok := m.GetMapping(url); ok {
		panic("duplicate url with get method")
	}
	m[GET].SetUrl(url, mapped)
}

// 增加Post方法下的接口
func (m MethodMaps) PostAdd(url string, mapped HandlerMapped) {
	if _, ok := m.GetMapping(url); ok {
		panic("duplicate url with Post method")
	}
	m[POST].SetUrl(url, mapped)

}

// 增加Put方法下的接口
func (m MethodMaps) PutAdd(url string, mapped HandlerMapped) {
	if _, ok := m.GetMapping(url); ok {
		panic("duplicate url with Put method")
	}
	m[PUT].SetUrl(url, mapped)

}

// 增加Delete方法下的接口
func (m MethodMaps) DeleteAdd(url string, mapped HandlerMapped) {
	if _, ok := m.GetMapping(url); ok {
		panic("duplicate url with Delete method")
	}
	m[DELETE].SetUrl(url, mapped)
}
func (h handler) SetUrl(url string, mapped HandlerMapped) {
	h[url] = mapped
}
