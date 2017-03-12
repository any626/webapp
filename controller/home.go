package controller

import (
	"fmt"
	"net/http"
	// "net/mail"
	// "html/template"
	// "github.com/any626/webapp/database"
)

// GetHome returns the the home page
func (c *Controller) GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}
