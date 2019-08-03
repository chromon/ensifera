package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(writer http.ResponseWriter, request *http.Request) {

}

func main() {

	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
