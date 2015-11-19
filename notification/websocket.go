package notification

import (
	"github.com/francoishill/leeroyci/database"
	"github.com/francoishill/leeroyci/websocket"
)

func sendWebsocket(job *database.Job, event string) {
	msg := websocket.NewMessage(job, event)
	websocket.Send(msg)
}
