package routes

import (
	"net/http"

	"github.com/qaiswardag/go_backend_api/internal/controller/authcontroller"
	"github.com/qaiswardag/go_backend_api/internal/controller/homecontroller"
	"github.com/qaiswardag/go_backend_api/internal/controller/userregistercontroller"
	"github.com/qaiswardag/go_backend_api/internal/controller/usersessionscontroller"
	"github.com/qaiswardag/go_backend_api/internal/middleware"
	"github.com/qaiswardag/go_backend_api/internal/qrouter"
)

type RouteHandler struct{}

func ChainMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func MainRouter() http.Handler {

	// Route::middleware(['first', 'second'])->group(function () {
	//   Route::get('/', function () {
	//   });
	//   Route::get('/user/profile', function () {
	//   });
	// });

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qrouter.NewRouter(w, r)
	})
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

func MainRouter2() http.Handler {
	mux := http.NewServeMux()

	// TODO: Add GET method for this route
	mux.Handle("/", middleware.Cors(
		middleware.GlobalMiddleware(
			http.HandlerFunc(homecontroller.Show),
		),
	))

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
