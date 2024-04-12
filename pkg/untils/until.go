package untils

import (
	"encoding/json"
	"io"
	"net/http"

)

func ParseBody(r *http.Request, X interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, X)
	if err != nil {
		return
	}
}
