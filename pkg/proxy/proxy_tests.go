package proxy

import (
	"bufio"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"
)



type Server struct {
	l net.Listener
}

func (s *Server) Start() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			continue
		}
		go s.serve(conn)
	}
}
func (s *Server) serve(conn net.Conn) {
	client, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = Proxy(conn, client)
	if err != nil {
		log.Fatal(err)
	}
}

func TestProxy(t *testing.T) {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}
	server := &Server{l: l}
	go server.Start()
	client, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com/", nil)
	if err != nil {
		t.Fatal(err)
	}
	err = req.Write(client)
	if err != nil {
		t.Fatal(err)
	}
	rb := bufio.NewReader(client)
	resp, err := http.ReadResponse(rb, req)
	if err != nil {
		t.Fatal(err)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bs))
}

