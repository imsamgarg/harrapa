package server

import (
	"encoding/json"
	"io"
)

func (s *Server) JsonDecodeAndValidate(r io.Reader, v any) error {

	if err := json.NewDecoder(r).Decode(v); err != nil {
		return err
	}

	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}
