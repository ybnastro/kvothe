package router

import (
	"net/http"

	"github.com/SurgicalSteel/kvothe/service"
	"github.com/gorilla/mux"
)

//RouterWrap is a custom wrapper type for router, you can add more configuration fields here
type RouterWrap struct {
	Router *mux.Router
}

var routeWrap *RouterWrap

// RegisterHandler is a RouterWrap 'method' to register your API endpoints.
// Usually handler calls services module
func (rw *RouterWrap) RegisterHandler() {
	rw.Router.HandleFunc("/ping", service.HandlePing).Methods(http.MethodGet)
}
