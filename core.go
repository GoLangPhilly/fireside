package main

import (
	"encoding/xml"
	"fmt"
	"net"
)

type Query struct {
	XMLName  xml.Name `xml:"query"`
	Xmlns    string   `xml:"xmlns,attr"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
	Digest   string   `xml:"digest"`
	Resource string   `xml:"resource"`
}

type Iq struct {
	XMLName xml.Name `xml:"iq"`
	Type    string   `xml:"type,attr"`
	Id      string   `xml:"id,attr"`
	Query   Query    `xml:"query"`
}

type Stream struct {
	XMLName xml.Name `xml:"stream:stream"`
	Iq      Iq       `xml:"iq"`
}

type LoginReq struct {
	Iq Iq `xml:"iq"`
}

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

    b := xml.NewDecoder(conn)
	fmt.Fprintf(conn, `<?xml version='1.0'?>
		<stream:stream
			from='localhost'
			id='abc123'
			to='james@localhost'
			version='1.0'
			xml:lang='en'
			xmlns='jabber:server'
			xmlns:stream='http://etherx.jabber.org/streams'>`)

	fmt.Fprintf(conn, "<stream:features/>")
	for {
        iqData := new(Iq)
        b.Decode(iqData)
		switch iqData.Type {
            case "get":
                r := &Iq{Id: iqData.Id, Type: "result"}
                r.Query = Query{Xmlns: "jabber:iq:auth"}
                output, _ := xml.Marshal(r)
                fmt.Fprintf(conn, string(output))
            case "set":
                // Need to perform auth lookup here
                i := Iq{Id: iqData.Id, Type: "result"}
                output, _ := xml.Marshal(i)
                fmt.Fprintf(conn, string(output))
            default:
                // Nothing
		}
	}

}
