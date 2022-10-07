package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil { // read the request body
		if err := json.Unmarshal([]byte(body), x); err != nil { // unmarshal the request body to the interface
			return // return if there is an error
		}
	}
}
