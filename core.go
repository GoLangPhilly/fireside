package main

import (
	"bufio"
	"fmt"
	"net"
)

func Serve(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	b := bufio.NewReader(conn)

	fmt.Fprintf(conn, `<?xml version='1.0'?>
		<stream:stream
			from='localhost'
			id='abc123'
			to='djworth@localhost'
			version='1.0'
			xml:lang='en'
			xmlns='jabber:server'
			xmlns:stream='http://etherx.jabber.org/streams'>`)

	fmt.Fprintf(conn, "<stream:features/>")
	for {
		line, _, err := b.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(line))
	}

}
