package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrDecodingFailed = errors.New("docoding failed")

func PathValues(r *http.Request, values ...string) (vMap map[string]string) {
	vMap = make(map[string]string, len(values))
	for _, value := range values {
		vMap[value] = r.PathValue(value)
	}
	return
}

func DecodeRequest[T any](request *T, w http.ResponseWriter, r *http.Request) (err error) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(request); err != nil {
		err = fmt.Errorf("%w: %s", ErrDecodingFailed, err)
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}
	return
}
