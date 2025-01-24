package routes

import (
	"encoding/json"
	"net/http"

	"github.com/qaiswardag/go_backend_api/internal/controller/authcontroller"
	"github.com/qaiswardag/go_backend_api/internal/controller/homecontroller"
	"github.com/qaiswardag/go_backend_api/internal/controller/userregistercontroller"
	"github.com/qaiswardag/go_backend_api/internal/controller/usersessionscontroller"
	"github.com/qaiswardag/go_backend_api/internal/middleware"
)

type RouteHandler struct{}

func ChainMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func MainRouter() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"message": "Method not allowed"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		homecontroller.Show(w, r)
	}))

	// TODO: Add POST method for this route
	mux.Handle("/user/sign-in", middleware.Cors(
		middleware.GlobalMiddleware(
			http.HandlerFunc(usersessionscontroller.Create),
		),
	))

	// TODO: Add POST method for this route
	mux.Handle("/user/sign-up", middleware.Cors(
		middleware.GlobalMiddleware(
			http.HandlerFunc(userregistercontroller.Create),
		),
	))

	// Add DELETE method for this route
	mux.Handle("/user/sign-out",
		middleware.Cors(
			middleware.GlobalMiddleware(
				middleware.RequireSessionMiddleware(
					http.HandlerFunc(authcontroller.Destroy),
				),
			),
		),
	)

	// Add GET method for this route
	mux.Handle("/user/user", ChainMiddlewares(
		http.HandlerFunc(authcontroller.Show),
		middleware.RequireSessionMiddleware,
		middleware.GlobalMiddleware,
		middleware.Cors,
	))

	return mux
}
