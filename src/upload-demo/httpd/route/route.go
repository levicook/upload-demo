package route

import (
	"fmt"
	"net/http"
	"time"
	"upload-demo/log"

	"github.com/gorilla/mux"
)

type (
	Routes []Route

	Route struct {
		Name        string
		Method      string
		Path        string
		HandlerFunc http.HandlerFunc
	}
)

func (rt Route) Signature() string {
	return fmt.Sprintf("%s %s", rt.Method, rt.Path)
}

func (rt Route) logFailure(r *http.Request, duration time.Duration, err interface{}) {
	log.Printf("%s %s | %s %s | error: %v", r.Method, r.RequestURI, rt.Name, duration, err)
}

func (rt Route) logVictory(r *http.Request, duration time.Duration) {
	log.Printf("%s %s | %s %s", r.Method, r.RequestURI, rt.Name, duration)
}

func (rt Route) log(r *http.Request, start time.Time) {
	if err := recover(); err != nil {
		rt.logFailure(r, time.Since(start), err)
		panic(err)
	}
	rt.logVictory(r, time.Since(start))
}

func (rt Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer rt.log(r, time.Now())
	rt.HandlerFunc.ServeHTTP(w, r)
}

func (routes Routes) Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route)
	}

	return router
}
