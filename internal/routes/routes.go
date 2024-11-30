package routes

import (
	"net/http"

	"github.com/qaiswardag/go_backend_api_jwt/internal/controller/home"
	"github.com/qaiswardag/go_backend_api_jwt/internal/controller/usercontroller"
	"github.com/qaiswardag/go_backend_api_jwt/internal/controller/usersettingscontroller"
	"github.com/qaiswardag/go_backend_api_jwt/internal/middleware"
)

type RouteHandler struct{}

func (h *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	home.Show(w, r)
}

func (h *RouteHandler) Home(w http.ResponseWriter, r *http.Request) {
	home.Show(w, r)
}

func (h *RouteHandler) LoginCreate(w http.ResponseWriter, r *http.Request) {
	usercontroller.Create(w, r)
}

func (h *RouteHandler) UserSettingsShow(w http.ResponseWriter, r *http.Request) {
	usersettingscontroller.Show(w, r)
}

func MainRouter() http.Handler {

	handler := &RouteHandler{}
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home)
	mux.Handle("/login", middleware.Cors(
		middleware.GlobalMiddleware(
			http.HandlerFunc(handler.LoginCreate),
		),
	),
	)

	mux.Handle("/user/settings",
		middleware.Cors(
			middleware.GlobalMiddleware(
				middleware.RequireSessionMiddleware(
					http.HandlerFunc(handler.UserSettingsShow),
				),
			),
		),
	)

	return mux
}
