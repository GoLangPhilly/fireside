package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

type connection struct {
	username string
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan string
}

func (c *connection) reader() {
	for {
		var message string
		err := websocket.Message.Receive(c.ws, &message)
		if err != nil {
			break
		}

		/*
			//We could hook into to the message broadcasting here
			go func() {
				fmt.Println("Hello World")
			}()
		*/

		h.broadcast <- fmt.Sprintf("%s: %s", c.username, message)

	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := websocket.Message.Send(c.ws, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {

	if err := ws.Request().ParseForm(); err != nil {
		fmt.Println(err.Error())
	}

	username := ws.Request().Form.Get("username")

	if _, exists := h.connections[username]; exists {
		ws.Write([]byte("User with this username exists already!"))
		ws.Close()
		return
	}
	c := &connection{send: make(chan string, 256), ws: ws, username: username}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}
