package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
	fmt.Println("adwadw")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")

	err = conn.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	for {
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
