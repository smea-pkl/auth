package handlers

import (
	"auth-system/services"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var MatchService = services.NewMatchmaker()

func MatchmakingHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Gagal melakukan upgrade ke websocker:", err)
		return
	}
	defer ws.Close()

	username := r.URL.Query().Get("user")
	if username == "" {
		username = "Guest"
	}

	client := &services.Client{
		Conn: ws,
		Username: username,
	}

	MatchService.FindMatch(client)

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			MatchService.RemoveClient(client)
			break
		}
	}
}