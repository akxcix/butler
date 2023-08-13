package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func RespondWithData(w http.ResponseWriter, r *http.Request, data interface{}) {
	res := &response{
		Status: http.StatusOK,
		Data:   data,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Error().Err(err).Msg("Unable to marshall data json.")
		RespondWithError(w, r, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func RespondWithError(w http.ResponseWriter, r *http.Request, e error, status int) {
	res := &response{
		Status: status,
		Error:  e.Error(),
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Error().Err(err).Msg("Unable to marshall error json.")
		status = http.StatusInternalServerError
		err := errors.New("Something bad happened!")
		json = []byte(fmt.Sprintf("{\"status\":%d,\"error\":\"%s\"}", status, err.Error()))
	}

	log.Error().
		Interface("status", status).
		Interface("error", err.Error()).
		Msg(http.StatusText(status))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}
