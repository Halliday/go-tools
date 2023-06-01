package httptools

import "net/http"

func UnsafeDisableCORS(resp http.ResponseWriter, req *http.Request) {
	requestMethod := req.Header.Get("Access-Control-Request-Method")
	if requestMethod != "" {
		resp.Header().Set("Access-Control-Allow-Methods", requestMethod)
	}
	requestHeaders := req.Header.Get("Access-Control-Request-Headers")
	if requestHeaders != "" {
		resp.Header().Set("Access-Control-Allow-Headers", requestHeaders)
	}

	resp.Header().Set("Access-Control-Allow-Origin", "*")
}

type Middleware struct {
	UnsafeDisableCORS bool
	http.Handler
}

func (m Middleware) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if m.UnsafeDisableCORS {
		UnsafeDisableCORS(resp, req)
	}
	m.Handler.ServeHTTP(resp, req)
}
