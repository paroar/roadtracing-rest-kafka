package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/paroar/roadtracing-rest-kafka/internal/kafka"
	"github.com/paroar/roadtracing-rest-kafka/internal/types"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// PositionHandler handler
func PositionHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		var _pos types.Position
		err := c.ReadJSON(&_pos)
		if err != nil {
			log.Println("read:", err)
			break
		}
		err = c.WriteJSON(_pos)
		if err != nil {
			log.Println(err)
			break
		}
		kafka.SavePositionToKafka(_pos)
	}
}
