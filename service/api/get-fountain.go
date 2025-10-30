package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

// listFountains is returning a JSON list of fountains
func (rt *_router) getFountain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sFountainId := ps.ByName("id")
	fountainId, err := strconv.Atoi(sFountainId)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("wrong Id received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var fountain Fountain
	found := false
	for _, f := range safeFountains.fountains {
		if f.Id == fountainId {
			fountain = f
			found = true
		}
	}
	if !found {
		rt.baseLogger.WithError(err).Warning("Id not found")
		NotFound(w)
		return
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(fountain)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("listFountains returned an error on Encode()")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: "Internal server error"})
	}
}
