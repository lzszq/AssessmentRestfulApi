package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func IMAGE(w http.ResponseWriter, statusCode int, data []byte) {
	w.WriteHeader(statusCode)
	w.Write(data)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
