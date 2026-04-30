package codec

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONCodec struct{}

func (JSONCodec) Encode(w http.ResponseWriter, _ *http.Request, v any) {
	w.Header().Set("Content-Type", "application/json")

	if err, ok := v.(error); ok {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	_ = json.NewEncoder(w).Encode(v)
}

func (JSONCodec) Decode(r *http.Request, v any) error {
	if r.Body == nil {
		return nil
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}
	return nil
}

func (JSONCodec) Validate(_ any) error {
	return nil
}
