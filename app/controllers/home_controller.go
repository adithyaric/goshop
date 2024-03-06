package controllers

import (
	// "fmt"
	"net/http"

	"github.com/unrolled/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to GoShop Home Page")
	render := render.New(render.Options{
		Layout:     "layout",
		Extensions: []string{".html", ".tmpl"},
	})

	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title": "Home Judul",
		"body":  "Home Deskripsi",
	})
}
