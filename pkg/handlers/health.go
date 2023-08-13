package handlers

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	data := "Help me if you can, I'm feeling down"
	RespondWithData(w, r, data)
}
