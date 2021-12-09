package timesheet

import (
	"time"
)

type Timesheet struct {
	Id      string `json:"id"` // Id
	Owner   string `json:"owner"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
}

type CreateTimesheetRequest struct {
	Owner   string `json:"owner"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
}

type UpdateTimesheetRequest struct {
	Id      string `json:"-"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
}

type Event struct {
	Id          string    `json:"id"`
	TimesheetId string    `json:"timesheetId"`
	Subject     string    `json:"subject"`
	WeekDay     string    `json:"weekDay"`
	TimeStart   time.Time `json:"timeStart"`
	TimeEnd     time.Time `json:"timeEnd"`
}

type CreateEventRequest struct {
	TimesheetId string    `json:"timesheetId"`
	Subject     string    `json:"subject"`
	WeekDay     string    `json:"weekDay"`
	TimeStart   time.Time `json:"timeStart"`
	TimeEnd     time.Time `json:"timeEnd"`
}

type UpdateEventRequest struct {
	Id          string    `json:"-"`
	Subject     string    `json:"subject"`
	WeekDay     string    `json:"weekDay"`
	TimeStart   time.Time `json:"timeStart"`
	TimeEnd     time.Time `json:"timeEnd"`
}

type EventSearchResponse struct {
	Result []*Event              `json:"result"`
}
