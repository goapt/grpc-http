package contract

import "net/http"

type Codec interface {
	Encode(w http.ResponseWriter, r *http.Request, v any)
	Decode(r *http.Request, v any) error
	Validate(v any) error
}
