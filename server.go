package web

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler

	// Start 开始方法
	Start(addr string) error

	// AddRoute 注册路由
	addRoute(method string, path string, handler HandleFunc)
}

type HTTPServer struct {
	router
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: newRouter(),
	}
}

// 确保 HTTPServer 肯定实现了 Server 接口
var _ Server = &HTTPServer{}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {
	// 查找路径 & 执行业务
	node, ok := h.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok || node.handler == nil {
		ctx.Resp.WriteHeader(404)
		ctx.Resp.Write([]byte("Not Found"))
		return
	}
	node.handler(ctx)
}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 这里可以做一些前置后置逻辑

	return http.Serve(l, h)
}

//func (h *HTTPServer) addRoute(method string, path string, handler HandleFunc) {
//	// 添加到路由树中
//}

func (h *HTTPServer) Get(path string, handler HandleFunc) {
	h.addRoute(http.MethodGet, path, handler)
}

func (h *HTTPServer) Post(path string, handler HandleFunc) {
	h.addRoute(http.MethodPost, path, handler)
}

func (h *HTTPServer) Put(path string, handler HandleFunc) {
	h.addRoute(http.MethodPut, path, handler)
}

func (h *HTTPServer) Delete(path string, handler HandleFunc) {
	h.addRoute(http.MethodDelete, path, handler)
}
