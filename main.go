package main

import (
	"flag"
)

var listen = flag.String("listen", ":5222", "xmpp service address")

func main() {
	flag.Parse()
	Serve(*listen)
}
