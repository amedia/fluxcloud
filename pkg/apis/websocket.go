package apis

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

//HandleWebsocket Handle Flux WebSocket connections
func HandleWebsocket(config APIConfig) error {
	config.Server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("WS Request for:", r.URL)
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("ws upgrade:", err)
			return
		}
		defer func() {
			log.Println("client disconnected")
			c.Close()
		}()

		log.Println("client connected!")

		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			log.Printf("ws recv: %s", message)
			err = c.WriteMessage(mt, message)

			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})

	return nil
}
