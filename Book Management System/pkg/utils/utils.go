package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) { //what is interfacae
	if body, err := io.ReadAll((r.Body)); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil { //doubts
			return
		}
	}

}
