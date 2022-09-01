package webRouter

import "net/http"

type WebRouter interface {
	GET(uri string, f func(rw http.ResponseWriter, r *http.Request))
	POST(uri string, f func(rw http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(rw http.ResponseWriter, r *http.Request))
	USE(func(next http.Handler) http.Handler)
	SERVE(port string)
}
