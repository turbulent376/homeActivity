package timesheet

import (
	"git.jetbrains.space/orbi/fcsd/api/internal/public"
	kitHttp "git.jetbrains.space/orbi/fcsd/kit/http"
	pb "git.jetbrains.space/orbi/fcsd/proto/timesheet"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
)

type Controller interface {
	CreateTimesheet(w http.ResponseWriter, r *http.Request)
	UpdateTimesheet(w http.ResponseWriter, r *http.Request)
	GetTimesheet(w http.ResponseWriter, r *http.Request)
	SearchTimesheet(w http.ResponseWriter, r *http.Request)
	DeleteTimesheet(w http.ResponseWriter, r *http.Request)
	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	GetEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	SearchEvents(w http.ResponseWriter, r *http.Request)
}

type ctrlImpl struct {
	kitHttp.BaseController
	timesheetRepo public.TimesheetRepository
}

func NewController(timesheetRepo public.TimesheetRepository) Controller {
	return &ctrlImpl{
		timesheetRepo: timesheetRepo,
	}
}

// CreateTimesheet godoc
// @Summary create timetable
// @Accept json
// @Router /timesheet/timetable [POST]
// @Param json body CreateTimesheetRequest true "request"
// @Produce json
// @Success 200 {object} Timesheet
// @Failure 500 {object} kitHttp.Error
// @tags timesheet
func (c *ctrlImpl) CreateTimesheet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request *CreateTimesheetRequest

	err := c.DecodeRequest(r, ctx, &request)

	if err != nil {
		c.RespondError(w, err)
		return
	}
	timesheet, err := c.timesheetRepo.CreateTimesheet(ctx, c.toCreateTimesheetPb(request))
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toTimesheetApi(timesheet))
}

// UpdateTimesheet godoc
// @Summary update timetable
// @Accept json
// @Router /timesheet/timetable/{id} [PUT]
// @Param json body UpdateTimesheetRequest true "request"
// @Param id path string true "id of timesheet"
// @Produce json
// @Success 200 {object} Timesheet
// @Failure 500 {object} kitHttp.Error
// @tags timesheet
func (c *ctrlImpl) UpdateTimesheet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request *UpdateTimesheetRequest

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.DecodeRequest(r, ctx, &request)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	request.Id = id

	timesheet, err := c.timesheetRepo.UpdateTimesheet(ctx, c.toUpdateTimesheetPb(request))
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toTimesheetApi(timesheet))
}

// GetTimesheet godoc
// @Summary getting timetable
// @Router /timesheet/timetable/{id} [GET]
// @Param id path string true "id"
// @Produce json
// @Success 200 {object} Timesheet
// @Failure 500 {object} kitHttp.Error
// @tags timesheet
func (c *ctrlImpl) GetTimesheet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	timesheet, err := c.timesheetRepo.GetTimesheet(ctx, &pb.TimesheetIdRequest{Id: id})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toTimesheetApi(timesheet))
}

// SearchTimesheet godoc
// @Summary getting timetable
// @Router /timesheet/timetable/{owner} [GET]
// @Param owner path string true "owner"
// @Param dateFromSearch formData string false "date from (type is time)"
// @Param dateToSearch formData string false "date to (type is time)"
// @Produce json
// @Success 200 {object} Timesheet
// @Failure 500 {object} kitHttp.Error
// @tags timesheet
func (c *ctrlImpl) SearchTimesheet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	owner, err := c.Var(r, ctx, "owner", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	dateFromSearch, err := c.FormValTime(r, ctx, "dateFromSearch", true)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	dateToSearch, err := c.FormValTime(r, ctx, "dateToSearch", true)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	timesheet, err := c.timesheetRepo.SearchTimesheet(ctx, &pb.SearchTimesheetRequest{
		Owner:          owner,
		DateFromSearch: timestamppb.New(*dateFromSearch),
		DateToSearch:   timestamppb.New(*dateToSearch),
	})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toTimesheetApi(timesheet))
}

// DeleteTimesheet godoc
// @Summary delete timetable
// @Router /timesheet/timetable/{id} [DELETE]
// @Param id path string true "id of timesheet"
// @Success 200
// @Failure 500 {object} kitHttp.Error
// @tags timesheet
func (c *ctrlImpl) DeleteTimesheet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.timesheetRepo.DeleteTimesheet(ctx, &pb.TimesheetIdRequest{Id: id})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, nil)
}

// CreateEvent godoc
// @Summary create event
// @Accept json
// @Router /timesheet/event [POST]
// @Param json body CreateEventRequest true "request"
// @Produce json
// @Success 200 {object} Event
// @Failure 500 {object} kitHttp.Error
// @tags event
func (c *ctrlImpl) CreateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request *CreateEventRequest

	err := c.DecodeRequest(r, ctx, &request)

	if err != nil {
		c.RespondError(w, err)
		return
	}
	event, err := c.timesheetRepo.CreateEvent(ctx, c.toCreateEventPb(request))
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toEventApi(event))
}

// UpdateEvent godoc
// @Summary update event
// @Accept json
// @Router /timesheet/event/{id} [PUT]
// @Param json body UpdateEventRequest true "request"
// @Param id path string true "id of event"
// @Produce json
// @Success 200 {object} Event
// @Failure 500 {object} kitHttp.Error
// @tags event
func (c *ctrlImpl) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request *UpdateEventRequest

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.DecodeRequest(r, ctx, &request)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	request.Id = id
	event, err := c.timesheetRepo.UpdateEvent(ctx, c.toUpdateEventPb(request))
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toEventApi(event))
}

// GetEvent godoc
// @Summary getting event
// @Router /timesheet/event/{id} [GET]
// @Param id path string true "id of event"
// @Produce json
// @Success 200 {object} Event
// @Failure 500 {object} kitHttp.Error
// @tags event
func (c *ctrlImpl) GetEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	event, err := c.timesheetRepo.GetEvent(ctx, &pb.EventIdRequest{
		Id: id,
	})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toEventApi(event))
}

// DeleteEvent godoc
// @Summary delete event
// @Router /timesheet/event/{id} [DELETE]
// @Param id path string true "id of event"
// @Success 200
// @Failure 500 {object} kitHttp.Error
// @tags event
func (c *ctrlImpl) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.timesheetRepo.DeleteEvent(ctx, &pb.EventIdRequest{Id: id})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, nil)
}

// SearchEvents godoc
// @Summary search events
// @Router /timesheet/event/{timesheetId} [GET]
// @Param timesheetId path string true "id of timesheet"
// @Produce json
// @Success 200 {object} EventSearchResponse
// @Failure 500 {object} kitHttp.Error
// @tags event
func (c *ctrlImpl) SearchEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	timesheetId, err := c.Var(r, ctx, "timesheetId", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	event, err := c.timesheetRepo.SearchEvents(ctx, &pb.EventIdRequest{Id: timesheetId})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toEventsApi(event))
}
