package tools

import (
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/halliday/go-errors"
)

var ErrInternal = &errors.RichError{
	Name: "internal_server_error",
	Code: 500,
	Desc: "An internal server error has occurred.",
}

func ServeError(resp http.ResponseWriter, err error) (unsafe error) {
	safe, unsafe := errors.Safe(err)
	resp.Header().Set("X-Content-Type-Options", "nosniff")
	resp.Header().Set("Content-Type", "application/json")
	if safe != nil {
		resp.WriteHeader(range1000(safe.(*errors.RichError).Code))
		ServeJSON(resp, safe)
	} else {
		resp.WriteHeader(ErrInternal.Code)
		ServeJSON(resp, ErrInternal)
	}
	if unsafe != nil {
		log.Printf("[     ] Unsafe: %v", unsafe)
	}
	return unsafe
}

func range1000(code int) int {
	for i := 1000; i < math.MaxInt/10; i *= 10 {
		if code < i {
			return code / (i / 1000)
		}
	}
	return code // unreachable
}

func ServeJSON(resp http.ResponseWriter, data interface{}) {
	resp.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(resp)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(data)
	if err != nil {
		log.Printf("tools.ServeJSON: can not marshal response: %v\n%v", err, data)
	}
}

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
