package controllers

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/letsgo-framework/examples/socket-io/log"
)

var SocketServer, _ = socketio.NewServer(nil)

func SocketHandler ( c  * gin.Context ) {
	SocketServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Debug("connected: %s", s.ID())
		return nil
	})
	SocketServer.OnError("/", func(e error) {
		log.Error("connected: %s", e.Error())
	})
	SocketServer.OnDisconnect("/", func(s socketio.Conn, msg string) {
		log.Error("connected: %s",msg)
	})

	go SocketServer.Serve()
	defer SocketServer.Close()

	SocketServer.ServeHTTP ( c.Writer, c.Request )
}