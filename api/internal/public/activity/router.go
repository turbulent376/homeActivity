package activity

import (
	kitHttp "github.com/turbulent376/kit/http"
	"github.com/gorilla/mux"
)

type Router struct {
	c Controller
}

func (r *Router) Set(authRouter, noAuthRouter *mux.Router) {
	authRouter.HandleFunc("/api/activity/activity", r.c.CreateActivity).Methods("POST")
	authRouter.HandleFunc("/api/activity/activity/{id}", r.c.UpdateActivity).Methods("PUT")
	authRouter.HandleFunc("/api/activity/activity/{id}", r.c.GetActivity).Methods("GET")
	authRouter.HandleFunc("/api/activity/list/{owner}", r.c.ListActivities).Methods("GET")
	authRouter.HandleFunc("/api/activity/listfamily/{family}", r.c.ListActivitiesByFamily).Methods("GET")
	authRouter.HandleFunc("/api/activity/activity/{id}", r.c.DeleteActivity).Methods("DELETE")
	authRouter.HandleFunc("/api/activity/activitytype", r.c.CreateActivityType).Methods("POST")
	authRouter.HandleFunc("/api/activity/activitytype/{id}", r.c.UpdateActivityType).Methods("PUT")
	authRouter.HandleFunc("/api/activity/activitytype/{id}", r.c.GetActivityType).Methods("GET")
	authRouter.HandleFunc("/api/activity/activitytype/{id}", r.c.DeleteActivityType).Methods("DELETE")
	authRouter.HandleFunc("/api/activity/listactivitytypes/{family}", r.c.ListActivityTypes).Methods("GET")
}

func NewRouter(c Controller) kitHttp.RouteSetter {
	return &Router{
		c: c,
	}
}
