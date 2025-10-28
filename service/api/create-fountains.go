package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type Fountain struct {
	Id        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    string  `json:"status"`
}

func (f Fountain) IsLatitudeValid() bool {
	return f.Latitude >= -90 && f.Latitude <= 90
}

func (f Fountain) IsLongitudeValid() bool {
	return f.Longitude >= -180 && f.Longitude <= 180
}

func (f Fountain) IsIdValid() bool {
	return f.Id > 0
}

type SafeFountains struct {
	mu        sync.Mutex
	fountains []Fountain
}

var safeFountains SafeFountains

// listFountains is returning a JSON list of fountains
func (rt *_router) createFountain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var fountain Fountain
	err := json.NewDecoder(r.Body).Decode(&fountain)
	_ = r.Body.Close()
	if err != nil {
		rt.baseLogger.WithError(err).Warning("wrong JSON received")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !fountain.IsLatitudeValid() {
		rt.baseLogger.Warning("Latitude is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !fountain.IsLongitudeValid() {
		rt.baseLogger.Warning("Longitude is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !fountain.IsIdValid() {
		rt.baseLogger.Warning("Id is not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	safeFountains.mu.Lock()
	safeFountains.fountains = append(safeFountains.fountains, fountain)
	safeFountains.mu.Unlock()

	w.WriteHeader(http.StatusCreated)
}
