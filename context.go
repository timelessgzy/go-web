package go_web

import "net/http"

type Context struct {
	req  *http.Request
	resp http.ResponseWriter
}
