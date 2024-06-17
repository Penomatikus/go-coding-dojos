package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrDecodingFailed   = errors.New("docoding failed")
)

func methodAllowed(want string, w http.ResponseWriter, r *http.Request) (err error) {
	if want != r.Method {
		err = ErrMethodNotAllowed
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	return
}

func pathValues(r *http.Request, values ...string) (vMap map[string]string) {
	vMap = make(map[string]string, len(values))
	for _, value := range values {
		vMap[value] = r.PathValue(value)
	}
	return
}

func decodeRequest[T any](request *T, w http.ResponseWriter, r *http.Request) (err error) {
	if err = json.NewDecoder(r.Body).Decode(request); err != nil {
		err = fmt.Errorf("%w: %s", ErrDecodingFailed, err)
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}
	return
}
