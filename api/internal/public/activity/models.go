package activity

import (
	"time"
)

type Activity struct {
	Id      string `json:"id"` // Id
	Owner   string `json:"owner"`
	Family  string `json:"family"`
	Type    string `json:"type"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
}

type CreateActivityRequest struct {
	Owner   string `json:"owner"`
	Family  string `json:"family"`
	Type    string `json:"type"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
}

type UpdateActivityRequest struct {
	Id      string `json:"-"`
	Owner   string `json:"owner"`
	Family  string `json:"family"`
	Type    string `json:"type"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo time.Time `json:"dateTo"`
}

type ListActivitiesResponse struct {
	Result []*Activity              `json:"result"`
}

type ActivityType struct {
	Id          string    `json:"id"`
	Family      string    `json:"family"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type CreateActivityTypeRequest struct {
	Family      string    `json:"family"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type UpdateActivityTypeRequest struct {
	Id          string    `json:"id"`
	Family      string    `json:"family"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type ListActivityTypesResponse struct {
	Result []*ActivityType              `json:"result"`
}
