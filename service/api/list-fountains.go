package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type JSONErrorMsg struct {
	Message string
}

// listFountains is returning a JSON list of fountains
func (rt *_router) listFountains(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	err := json.NewEncoder(w).Encode(safeFountains.fountains)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("listFountains returned an error on Encode()")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: "Internal server error"})
	}
}
