package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/fountains", rt.listFountains)
	rt.router.POST("/fountains", rt.createFountain)

	rt.router.GET("/fountains/:id", rt.getFountain)
	rt.router.DELETE("/fountains/:id", rt.deleteFountain)

	return rt.router
}
