package routes

import (
	"net/http"

	"github.com/qaiswardag/go_backend_api/internal/controller/authcontroller"
	"github.com/qaiswardag/go_backend_api/internal/controller/homecontroller"
	"github.com/qaiswardag/go_backend_api/internal/controller/userregistercontroller"
	"github.com/qaiswardag/go_backend_api/internal/middleware"
	"github.com/qaiswardag/go_backend_api/internal/router"
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

	// Wrap mux with a handler that calls SayHi() for every request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router.SayHi()      // This runs on every request
		mux.ServeHTTP(w, r) // Forward the request to actual router
	})

	// TODO: Add GET method for this route
	mux.Handle("/", middleware.Cors(
		middleware.GlobalMiddleware(
			http.HandlerFunc(homecontroller.Show),
		),
	))

	// TODO: Add POST method for this route

	// mux.Handle("/user/sign-in", middleware.Cors(
	// 	middleware.GlobalMiddleware(
	// 		http.HandlerFunc(usersessionscontroller.Create),
	// 	),
	// ))

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
