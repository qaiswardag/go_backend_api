package routes

import (
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/internal/controller/home"
	"github.com/qaiswardag/go_backend_api_jwt/internal/controller/usersessionscontroller"
	"github.com/qaiswardag/go_backend_api_jwt/internal/controller/usersettingscontroller"
	"github.com/qaiswardag/go_backend_api_jwt/internal/middleware"
)

type RouteHandler struct{}

func MainRouter() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		home.Show(w, r)
	}))

	mux.Handle("/login", middleware.Cors(
		middleware.GlobalMiddleware(
			http.HandlerFunc(usersessionscontroller.Create),
		),
	))

	mux.Handle("/user/settings",
		middleware.Cors(
			middleware.GlobalMiddleware(
				middleware.RequireSessionMiddleware(
					http.HandlerFunc(usersettingscontroller.Show),
				),
			),
		),
	)

	return mux
}
