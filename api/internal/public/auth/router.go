package auth

import (
	kitHttp "git.jetbrains.space/orbi/fcsd/kit/http"
	"github.com/gorilla/mux"
)

type Router struct {
	c Controller
}

func (r *Router) Set(authRouter, noAuthRouter *mux.Router) {
	noAuthRouter.HandleFunc("/api/auth/login", r.c.AuthUserByEmail).Methods("POST")
	noAuthRouter.HandleFunc("/api/auth/firebase", r.c.AuthUserByFirebase).Methods("POST")
	noAuthRouter.HandleFunc("/api/auth/user/new", r.c.CreateUser).Methods("POST")
	authRouter.HandleFunc("/api/auth/refresh", r.c.RefreshToken).Methods("POST")
	authRouter.HandleFunc("/api/auth/user/{userId}", r.c.UserInfo).Methods("GET")
	authRouter.HandleFunc("/api/auth/user/{userId}", r.c.DeleteUser).Methods("DELETE")
	authRouter.HandleFunc("/api/auth/user/{userId}", r.c.UpdateUser).Methods("PUT")
	authRouter.HandleFunc("/api/auth/logout", r.c.Logout).Methods("GET")
	authRouter.HandleFunc("/api/auth/user/notify/token", r.c.SaveFCMToken).Methods("POST")
}

func NewRouter(c Controller) kitHttp.RouteSetter {
	return &Router{
		c: c,
	}
}
