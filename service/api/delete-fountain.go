package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// listFountains is returning a JSON list of fountains
func (rt *_router) deleteFountain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sFountainId := ps.ByName("id")
	fountainId, err := strconv.Atoi(sFountainId)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("wrong Id received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	found := false
	for i, f := range safeFountains.fountains {
		if f.Id == fountainId {
			safeFountains.mu.Lock()
			safeFountains.fountains[i] = safeFountains.fountains[len(safeFountains.fountains)-1]
			safeFountains.fountains = safeFountains.fountains[:len(safeFountains.fountains)-1]
			safeFountains.mu.Unlock()
			found = true
		}
	}
	if !found {
		rt.baseLogger.WithError(err).Warning("Id not found")
		NotFound(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
