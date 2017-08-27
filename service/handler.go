package service

import (
	"net/http"

	"github.com/unrolled/render"
)

const (
	fakeMatchLocationResult = "/matches/5a003b78-409e-4452-b456-a6f0dcee05bd"
)

func createMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Location", "Some value")
		formatter.JSON(w,
			http.StatusCreated,
			struct{ Test string }{"This is a test"})
	}
}
