package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"net"
_    "io"
)

type Stream struct {
    XMLName xml.Name `xml:"stream:stream"`
    Iq
}

type Query struct {
    XMLName xml.Name `xml:"query"`
    Xmlns string `xml:"xmlns,attr"`
    Username string `xml:"username"`
    Password string `xml:"password"`
    Digest string `xml:"digest"`
    Resource string `xml:"resource"`
}


type Iq struct {
    XMLName xml.Name `xml:"iq"`
	Type string `xml:"type,attr"`
	Id   string `xml:"id,attr"`
    Query
}


type LoginReq struct {
	//<?xml version='1.0' ?><stream:stream to='localhost' xmlns='jabber:client'
	//xmlns:stream='http://etherx.jabber.org/streams' version='1.0'>
	//<iq type='get' id='purple3d9c697f'><query xmlns='jabber:iq:auth'>
	//<username>james</username></query></iq></stream:stream>
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

func handleLogin(loginRequest string) (x string, e error) {

	v := LoginReq{}

	err := xml.Unmarshal([]byte(loginRequest), &v)

	if err != nil {
		fmt.Printf("err: %v", err)
		return "", err
	}
	fmt.Printf("Id: %v\n", v.Iq.Id)
    return v.Iq.Id, err
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Handling connection....")

	b := bufio.NewReader(conn)

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
		line, _, err := b.ReadLine()
		if err != nil {
			break
		}
		loginId, err := handleLogin(string(line))
        fmt.Println(loginId)
		fmt.Println(string(line))
        r := &Stream{}
        r.Iq = Iq{Id: loginId, Type: "result"}
        r.Iq.Query = Query{Xmlns: "jabber:iq:auth"}
        output, err := xml.Marshal(r)
        fmt.Println(string(output))
        
        
        line, _, err = b.ReadLine()
        loginId2, err := handleLogin(string(line))
        i := &Stream{}
        i.Iq = Iq{Id: loginId2, Type: "result"}
        output2, err := xml.Marshal(i)
        fmt.Println(string(output2))
        //fmt.Fprintf(conn, string(output))
	}


}
