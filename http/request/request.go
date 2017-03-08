package request

import (
	"net/http"
	"github.com/gorilla/schema"
)

func Decode(r *http.Request, i interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	var decoder = schema.NewDecoder()
	err = decoder.Decode(i, r.Form)
	if err != nil {
		return err
	}

	return nil
}
