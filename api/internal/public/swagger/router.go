package swagger

import (
	_ "git.jetbrains.space/orbi/fcsd/api/docs"
	kitHttp "git.jetbrains.space/orbi/fcsd/kit/http"
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
