package main

import (
	_ "bufio"
	"encoding/xml"
	"fmt"
	_ "io"
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
	fmt.Println("Server started")
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

func ParseIq(request string) (iq Iq, e error) {

	v := LoginReq{}

	err := xml.Unmarshal([]byte(request), &v)

	if err != nil {
		fmt.Printf("err: %v", err)
		return v.Iq, err
	}
	fmt.Printf("Id: %v\n", v.Iq.Id)
	return v.Iq, err
}

func Authenticate(username string, password string) (authXml string) {
	return "hey"
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Handling connection....")

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
			fmt.Println(string(output))
			fmt.Fprintf(conn, string(output))
		case "set":
			fmt.Println("** Case Set **")
            i := Iq{Id: iqData.Id, Type: "result"}
			output2, _ := xml.Marshal(i)
			fmt.Println(string(output2))
			fmt.Fprintf(conn, string(output2))
		default:
			//fmt.Println(string(line))
		}
	}

}
