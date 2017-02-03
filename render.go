package render

import (
	"encoding/json"
	"net/http"

	"github.com/google/jsonapi"
)

// Render ...
type Render struct {
	w      http.ResponseWriter
	status int
}

// MarshalOne ...
func (r *Render) MarshalOne(model interface{}) error {
	if err := jsonapi.MarshalOnePayload(r.w, model); err != nil {
		return err
	}
	return nil
}

// MarshalMany ...
func (r *Render) MarshalMany(models interface{}) error {
	if err := jsonapi.MarshalManyPayload(r.w, models); err != nil {
		return err
	}
	return nil
}

// MarshalErrors ...
func (r *Render) MarshalErrors(errors []*jsonapi.ErrorObject) error {
	if err := jsonapi.MarshalErrors(r.w, errors); err != nil {
		return err
	}
	return nil
}

// JSON ...
func JSON(w http.ResponseWriter, status int) *Render {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(status)
	return &Render{w: w, status: status}
}

// Bind ...
func Bind(r *http.Request, model interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(model); err != nil {
		return err
	}
	return nil
}
