package httptools

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/halliday/go-errors"
)

var ErrInternal = &errors.RichError{Name: "internal", Code: 500, Desc: "An internal server error has occurred."}
var Logger = log.Default()

func ServeError(resp http.ResponseWriter, req *http.Request, err error) (unsafe error) {
	safe, unsafe := errors.Safe(err)
	resp.Header().Set("X-Content-Type-Options", "nosniff")
	resp.Header().Set("Content-Type", "application/json")
	if safe != nil {
		resp.WriteHeader(range1000(safe.(*errors.RichError).Code))
		ServeJSON(resp, req, safe)
	} else {
		resp.WriteHeader(ErrInternal.Code)
		ServeJSON(resp, req, ErrInternal)
	}
	if unsafe != nil {
		Logger.Printf("Unsafe E: %v", unsafe)
	}
	return unsafe
}

func range1000(code int) int {
	if code < 1000 {
		return code
	}
	for code > 1000 {
		code /= 10
	}
	return code
}

func ServeJSON(resp http.ResponseWriter, req *http.Request, data interface{}) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(data)
	if err != nil {
		ServeError(resp, req, err)
		return
	}
	resp.Header().Add("Content-Type", "application/json")
	resp.Header().Add("Content-Length", strconv.Itoa(b.Len()))
	resp.Write(b.Bytes())
}
