package httpx

import (
	"fmt"
	"net/http"
)

type Server struct {
	// ...

	// 监听的地址和端口
	// 默认 127.0.0.1:8080 :8080
	Addr     string
	routeMap routeMap
}

// NewServer 申请一个服务
func NewServer(Addr string) *Server {
	if Addr == "" {
		Addr = "127.0.0.1:8080"
	}

	return &Server{
		Addr,
		newRouteMap(),
	}
}

// Start 启动服务
func (s *Server) Start() {
	fmt.Printf("start listen and server :%s \n", s.Addr)
	http.ListenAndServe(s.Addr, s)
}

// 实现Handler接口，匹配方法以及路径
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//转发给doHandler进行执行
	s.doHandler(w, req)
}

// 判断需要执行的Http Method，从而查找对应的接口并且执行
func (s *Server) doHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case http.MethodPost:
		{
			if hm, ok := s.routeMap[post][req.URL.RequestURI()]; ok {
				hm(w, req)
			}
		}
	case http.MethodGet:
		{
			if hm, ok := s.routeMap[get][req.URL.RequestURI()]; ok {
				hm(w, req)
			}
		}
	default:
		{

			return
		}
	}
}

// RouteAdds 注册路由
func (s *Server) RouteAdds(routes []Route) {
	for _, route := range routes {
		switch route.Method {
		case http.MethodPost:
			s.routeMap.post(route.Path, route.HandlerFunc)
		case http.MethodGet:
			s.routeMap.get(route.Path, route.HandlerFunc)
		default:
			continue
		}
	}
}
