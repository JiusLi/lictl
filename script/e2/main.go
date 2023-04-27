package main

import (
	"fmt"
	"lictl/tools/errorx"
	"lictl/tools/httpx"
	"lictl/tools/response"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	srv := httpx.NewServer("")
	initRoute(srv)
	srv.Start()
}

func initRoute(srv *httpx.Server) {
	srv.RouteAdds(
		httpx.Middlewares([]httpx.Mi{LOG, TEST},
			[]httpx.Route{
				{
					Method:      http.MethodPost,
					Path:        "/index",
					HandlerFunc: IndexHandler(),
				},
				{
					Method:      http.MethodGet,
					Path:        "/client",
					HandlerFunc: WS(),
				},
			}...),
	)
}

// middleware

func LOG(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("LOG")
		handlerFunc(writer, request)
	}
}

func TEST(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("TEST")
		handlerFunc(writer, request)
	}
}

// handler

func IndexHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("这是控制器")
		// 可拓展到 logic
		//response.Response(writer, "", nil)

		_ = map[string]interface{}{
			"success": "恭喜你跑完一遍流程了",
		}
		//response.Response(writer, nil, errorx.New(3001, "奇葩错误", nil))
		response.Response(writer, nil, errorx.A)

	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WS() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Header.Get("Connection") != "Upgrade" {
			return
		}

		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Println(err)
		}

		// 连接成功
		log.Println("Client Connected")

		// 往客户端发送消息
		err = conn.WriteMessage(1, []byte("Hi Client!"))
		if err != nil {
			log.Println(err)
		}

		for {
			// 读取客户端消息
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			log.Println(string(p))

			err = conn.WriteMessage(messageType, p)
			if err != nil {
				log.Println(err)
				return
			}
		}

	}
}
