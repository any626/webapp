package router

import (
    // "fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/any626/webapp/controller"
    "github.com/any626/webapp/service"
    // "github.com/garyburd/redigo/redis"
    redistore "gopkg.in/boj/redistore.v1"
)

var SessionKey string = "session-key"

type Router struct {
    controller *controller.Controller
    Service *service.Service
    RediStore *redistore.RediStore
}

func NewRouter(c *controller.Controller, s *service.Service, redisStore *redistore.RediStore) *mux.Router {

    r := &Router{controller: c, Service: s, RediStore: redisStore}

	mRouter := mux.NewRouter()

	mRouter.HandleFunc("/", c.GetIndex).Methods("GET")
    mRouter.HandleFunc("/register", c.GetRegister).Methods("GET")
    mRouter.HandleFunc("/register", c.PostRegister).Methods("POST")
    mRouter.HandleFunc("/signin", c.GetSignIn).Methods("GET")
    mRouter.HandleFunc("/signin", c.PostSignIn).Methods("POST")
    mRouter.HandleFunc("/logout", c.Logout)
    


    // mRouter.HandleFunc("/home", use(http.HandlerFunc(c.GetHome), r.validateJWT)).Methods("GET")
    // mRouter.HandleFunc("/home", Middleware(http.HandlerFunc(c.GetHome))).Methods("GET")
    mRouter.HandleFunc("/home", use(c.GetHome, r.Authenticate)).Methods("GET")
    // mRouter.HandleFunc("/home", log(HomeHandler)).Methods("GET")
    // mRouter.HandleFunc("/home", use(http.HandlerFunc(c.GetHome), Csrf))

	return mRouter
}

// Middleware chainer
func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
    for _, m := range middleware {
        h = m(h)
    }
    return h
}

//
func (rt *Router) Authenticate(h http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, err := rt.RediStore.Get(r, SessionKey)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        _, ok := session.Values["userId"]
        if !ok {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        h.ServeHTTP(w, r)
    })
}