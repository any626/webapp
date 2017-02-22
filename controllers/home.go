package controllers

import (
	"net/http"
	"fmt"
)

type HomeController struct {
}

func (c *HomeController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home");
}