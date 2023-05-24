package web

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewHTTPServer()
	s.Get("/user", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello, user"))
	})
	err := s.Start(":8080")
	if err != nil {
		fmt.Println(err)
	}

}
