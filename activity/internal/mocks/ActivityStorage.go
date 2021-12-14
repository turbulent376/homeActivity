// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/turbulent376/homeactivity/activity/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// TimetableStorage is an autogenerated mock type for the TimetableStorage type
type TimetableStorage struct {
	mock.Mock
}

// CreateTimetable provides a mock function with given fields: ctx, Timesheet
func (_m *TimetableStorage) CreateTimetable(ctx context.Context, Timesheet *domain.Timesheet) (*domain.Timesheet, error) {
	ret := _m.Called(ctx, Timesheet)

	var r0 *domain.Timesheet
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Timesheet) *domain.Timesheet); ok {
		r0 = rf(ctx, Timesheet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Timesheet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Timesheet) error); ok {
		r1 = rf(ctx, Timesheet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTimetable provides a mock function with given fields: ctx, id
func (_m *TimetableStorage) DeleteTimetable(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTimetable provides a mock function with given fields: ctx, id
func (_m *TimetableStorage) GetTimetable(ctx context.Context, id string) (bool, *domain.Timesheet, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *domain.Timesheet
	if rf, ok := ret.Get(1).(func(context.Context, string) *domain.Timesheet); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Timesheet)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SearchTimetable provides a mock function with given fields: ctx, rq
func (_m *TimetableStorage) SearchTimetable(ctx context.Context, rq *domain.SearchTimesheetRequest) (bool, *domain.Timesheet, error) {
	ret := _m.Called(ctx, rq)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *domain.SearchTimesheetRequest) bool); ok {
		r0 = rf(ctx, rq)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *domain.Timesheet
	if rf, ok := ret.Get(1).(func(context.Context, *domain.SearchTimesheetRequest) *domain.Timesheet); ok {
		r1 = rf(ctx, rq)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Timesheet)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *domain.SearchTimesheetRequest) error); ok {
		r2 = rf(ctx, rq)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateTimetable provides a mock function with given fields: ctx, Timesheet
func (_m *TimetableStorage) UpdateTimetable(ctx context.Context, Timesheet *domain.Timesheet) (*domain.Timesheet, error) {
	ret := _m.Called(ctx, Timesheet)

	var r0 *domain.Timesheet
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Timesheet) *domain.Timesheet); ok {
		r0 = rf(ctx, Timesheet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Timesheet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Timesheet) error); ok {
		r1 = rf(ctx, Timesheet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
