package go_web

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx Context)

type Server interface {
	http.Handler

	// Start 开始方法
	Start(addr string) error

	// AddRoute 注册路由
	AddRoute(method string, path string, handler HandleFunc)
}

type HTTPServer struct {
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := Context{
		req:  request,
		resp: writer,
	}
	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx Context) {
	// 查找路径 & 执行业务
}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 这里可以做一些前置后置逻辑

	return http.Serve(l, h)
}

func (h *HTTPServer) AddRoute(method string, path string, handler HandleFunc) {
	// 添加到路由树中
}

func (h *HTTPServer) Get(path string, handler HandleFunc) {
	h.AddRoute(http.MethodGet, path, handler)
}

func (h *HTTPServer) Post(path string, handler HandleFunc) {
	h.AddRoute(http.MethodPost, path, handler)
}

func (h *HTTPServer) Put(path string, handler HandleFunc) {
	h.AddRoute(http.MethodPut, path, handler)
}

func (h *HTTPServer) Delete(path string, handler HandleFunc) {
	h.AddRoute(http.MethodDelete, path, handler)
}
