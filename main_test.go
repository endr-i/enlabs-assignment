package main

import (
	"assignment/server"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"net"
	"testing"
)

func serve(handler fasthttp.RequestHandler) *fasthttputil.InmemoryListener {
	ln := fasthttputil.NewInmemoryListener()
	go func() {
		err := fasthttp.Serve(ln, handler)
		if err != nil {
			panic(err)
		}
	}()

	return ln
}

// TODO: create test cases
func TestServer(t *testing.T) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("/ping")
	req.SetHost("localhost")
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")

	resp := fasthttp.AcquireResponse()

	router := server.Router()
	ln := serve(router.Handler)
	defer ln.Close()

	client := fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}
	err := client.Do(req, resp)
	if err != nil {
		t.Fatal(err)
	}
	if string(resp.Body()) != "pong" {
		t.Fatal("cannot handle routing")
	}
}
