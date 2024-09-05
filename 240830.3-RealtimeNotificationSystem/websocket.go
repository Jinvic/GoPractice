package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{} // use default options
	conn, _ := upgrader.Upgrade(w, r, nil)
	// defer conn.Close()
	ch := subAllChannels()

	go func() {
		for msg := range ch {
			channel := msg.Channel
			var m Message
			json.Unmarshal([]byte(msg.Payload), &m)
			m.Channel = channel
			mJson, _ := json.Marshal(m)
			// fmt.Println(m)
			conn.WriteMessage(websocket.TextMessage, mJson)
		}
	}()
}
