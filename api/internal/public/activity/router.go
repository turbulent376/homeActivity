package timesheet

import (
	kitHttp "git.jetbrains.space/orbi/fcsd/kit/http"
	"github.com/gorilla/mux"
)

type Router struct {
	c Controller
}

func (r *Router) Set(authRouter, noAuthRouter *mux.Router) {
	authRouter.HandleFunc("/api/timesheet/timetable", r.c.CreateTimesheet).Methods("POST")
	authRouter.HandleFunc("/api/timesheet/timetable/{id}", r.c.UpdateTimesheet).Methods("PUT")
	authRouter.HandleFunc("/api/timesheet/timetable/{id}", r.c.GetTimesheet).Methods("GET")
	authRouter.HandleFunc("/api/timesheet/timetable/{owner}", r.c.SearchTimesheet).Methods("GET")
	authRouter.HandleFunc("/api/timesheet/timetable/{id}", r.c.DeleteTimesheet).Methods("DELETE")
	authRouter.HandleFunc("/api/timesheet/event", r.c.CreateEvent).Methods("POST")
	authRouter.HandleFunc("/api/timesheet/event/{id}", r.c.UpdateEvent).Methods("PUT")
	authRouter.HandleFunc("/api/timesheet/event/{id}", r.c.GetEvent).Methods("GET")
	authRouter.HandleFunc("/api/timesheet/event/{id}", r.c.DeleteEvent).Methods("DELETE")
	authRouter.HandleFunc("/api/timesheet/event/{timesheetId}", r.c.SearchEvents).Methods("GET")
}

func NewRouter(c Controller) kitHttp.RouteSetter {
	return &Router{
		c: c,
	}
}
