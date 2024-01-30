package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// JSON decodes json from request body to ptr
var (
	// ErrRequestContentTypeNotJSON is used when the request content type is not application/json.
	ErrRequestContentTypeNotJSON = errors.New("request content type is not application/json")
	// ErrRequestJSONInvalid is used when the request json is invalid.
	ErrRequestJSONInvalid = errors.New("request json invalid")
)

// JSON decodes json from request body to ptr
func JSON(r *http.Request, ptr any) (err error) {
	// check content type
	if r.Header.Get("Content-Type") != "application/json" {
		err = ErrRequestContentTypeNotJSON
		return
	}

	// get body
	err = json.NewDecoder(r.Body).Decode(ptr)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrRequestJSONInvalid, err)
		return
	}

	return
}