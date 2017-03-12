package router

import (
	// "fmt"
	"github.com/any626/webapp/controller"
	"github.com/any626/webapp/service"
	"github.com/gorilla/mux"
	"net/http"
	// "github.com/any626/webapp/shared"
	// "github.com/garyburd/redigo/redis"
	// redistore "gopkg.in/boj/redistore.v1"
)

var SessionKey string = "session-key"

type Router struct {
	controller *controller.Controller
	Service    *service.Service
}

func NewRouter(c *controller.Controller, s *service.Service) *mux.Router {

	r := &Router{controller: c, Service: s}

	mRouter := mux.NewRouter()

	mRouter.HandleFunc("/", c.GetIndex).Methods("GET")
	mRouter.HandleFunc("/register", r.Authenticated(c.GetRegister)).Methods("GET")
	mRouter.HandleFunc("/register", r.Authenticated(c.PostRegister)).Methods("POST")
	mRouter.HandleFunc("/sign-in", r.Authenticated(c.GetSignIn)).Methods("GET")
	mRouter.HandleFunc("/sign-in", r.Authenticated(c.PostSignIn)).Methods("POST")
	mRouter.HandleFunc("/sign-out", c.Logout)

	mRouter.HandleFunc("/home", use(c.GetHome, r.Authenticate)).Methods("GET")

	return mRouter
}

// Middleware chainer
func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func (rt *Router) Authenticated(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rt.Service.IsAuth(w, r) {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func (rt *Router) Authenticate(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !rt.Service.IsAuth(w, r) {
			http.Redirect(w, r, "/sign-in", http.StatusFound)
			return
		}

		h.ServeHTTP(w, r)
	})
}
