package client

import (
	"fmt"
	"net"
	"os"
	"slices"
)

type Client struct {
	srvAddr string
	port    int
}

func NewClient(addr string, port int) *Client {
	return &Client{
		srvAddr: addr,
		port:    port,
	}
}

func (c *Client) SendQuery(query []byte) []byte {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", c.srvAddr, c.port))
	if err != nil {
		fmt.Printf("Dial err: %v\n", err)
		os.Exit(-1)
	}
	defer conn.Close()

	if _, err = conn.Write(query); err != nil {
		fmt.Printf("Write err: %v\n", err)
		os.Exit(-1)
	}

	resp := make([]byte, 1024)
	length, err := conn.Read(resp)
	if err != nil {
		fmt.Printf("Read err: %v\n", err)
		os.Exit(-1)
	}

	if !hasTheSameID(query, resp) {
		fmt.Printf("response does not have the same ID of the query - q: %v, r: %v\n", query, resp)
		os.Exit(-1)
	}

	return resp[:length]
}

func hasTheSameID(query []byte, response []byte) bool {
	return slices.Equal(query[:2], response[:2])
}
