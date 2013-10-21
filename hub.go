package main

type hub struct {
	// Registered connections.
	connections map[string]*connection

	// Inbound messages from the connections.
	broadcast chan string

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

var h = hub{
	broadcast:   make(chan string),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[string]*connection),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c.username] = c
		case c := <-h.unregister:
			delete(h.connections, c.username)
			close(c.send)
		case m := <-h.broadcast:
			for u := range h.connections {
				c := h.connections[u]
				select {
				case c.send <- m:
				default:
					delete(h.connections, c.username)
					close(c.send)
					go c.ws.Close()
				}
			}
		}
	}
}
