package swagger

import (
	_ "github.com/turbulent376/homeactivity/api/docs"
	kitHttp "github.com/turbulent376/kit/http"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct{}

func (r *Router) Set(authRouter, noAuthRouter *mux.Router) {
	noAuthRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

func NewRouter() kitHttp.RouteSetter {
	return &Router{}
}
