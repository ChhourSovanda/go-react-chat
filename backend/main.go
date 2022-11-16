package main

import (
	"fmt"
	"net/http"

	"github.com/wanda168/go-react-chat/pkg/websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {

}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Wanda full stack chat project")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
