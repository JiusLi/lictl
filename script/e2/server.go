package e2

import (
	"fmt"
	"net/http"
)

type Server struct {
	// addr
	Addr   string
	Router MethodMaps
}

type Route struct {
	Method  string
	Path    string
	Handler HandlerFunc
}

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

func NewSRV(Addr string) *Server {
	return &Server{
		Addr,
		NewRouter(),
	}
}

// 实现Handler接口，匹配方法以及路径
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//转发给doHandler进行执行
	s.doHandler(w, req)
}

// 判断需要执行的Http Method，从而查找对应的接口并且执行
func (s *Server) doHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("ths...", req.Method)
	switch req.Method {
	case http.MethodPost:
		{
			if hm, ok := s.Router[POST][req.URL.RequestURI()]; ok {
				hm(w, req)
			}

		}
	default:
		{

		}
	}
}

func (s *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		switch r.Method {
		case "POST":
			s.Router.POST(r.Path, r.Handler)
		default:
			continue
		}
	}
}

func (s *Server) Start() {
	fmt.Printf("start listen and server :%s \n", s.Addr)
	http.ListenAndServe(s.Addr, s)
}
