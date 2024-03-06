package controllers

import (
	"net/http"

	"github.com/adithyaric/goshop/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout:     "layout",
		Extensions: []string{".html", ".tmpl"},
	})

	productModel := models.Product{}
	products, err := productModel.GetProducts(server.DB)
	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products": products,
	})
}
