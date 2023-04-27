package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

/*
http.HandleFunc 注册处理函数
http.Handle 注册处理器
*/

type mi func(handlerFunc http.HandlerFunc) http.HandlerFunc

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", TTT([]mi{LOG, LOG2}, index))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func LOG(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintln(writer, "LOG")
		handlerFunc.ServeHTTP(writer, request)
	}
}

func LOG2(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintln(writer, "LOG2")
		handlerFunc.ServeHTTP(writer, request)
	}
}

func TTT(ms []mi, hf http.HandlerFunc) http.HandlerFunc {
	for i := len(ms) - 1; i >= 0; i-- {
		hf = ms[i](hf)
	}

	//for i := 0; i <= len(ms)-1; i++ {
	//	hf = ms[i](hf)
	//}

	return hf
	//return hf
}
