package router

import (
    "fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/any626/webapp/controller"
    "github.com/any626/webapp/service"
)

type Router struct {
    controller *controller.Controller
    Service *service.Service
}

func NewRouter(c *controller.Controller, s *service.Service) *mux.Router {

    // r := &Router{controller: c, Service: s}

	mRouter := mux.NewRouter()

	mRouter.HandleFunc("/", c.GetIndex).Methods("GET")
    mRouter.HandleFunc("/register", c.GetRegister).Methods("GET")
    mRouter.HandleFunc("/register", c.PostRegister).Methods("POST")

    // mRouter.HandleFunc("/home", use(http.HandlerFunc(c.GetHome), r.validateJWT)).Methods("GET")
    // mRouter.HandleFunc("/home", Middleware(http.HandlerFunc(c.GetHome))).Methods("GET")
    mRouter.HandleFunc("/home", use(c.GetHome, ValidateJWT)).Methods("GET")
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
func ValidateJWT(h http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // cookie, err := r.Cookie("Auth")
        // if err != nil {
        //     w.WriteHeader(http.StatusInternalServerError)
        //     return
        // }
        // err = rt.Service.ParseJwt(cookie.Value)
        // if err != nil {
        //     w.WriteHeader(http.StatusInternalServerError)
        //     return
        // }
        fmt.Println("first")
        return
        h.ServeHTTP(w, r)
        fmt.Println("second")
    })
}