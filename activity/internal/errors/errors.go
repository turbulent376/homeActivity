// this file is generated by servgen util based on a template at 2021-06-26 10:37:24 +0300 MSK
package errors

import (
	"context"
	"github.com/turbulent376/kit/er"
	pb "github.com/turbulent376/proto/activity"
)

var (
	ErrActivityIdIsEmpty = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeActivityIdEmpty, "empty id").C(ctx).Err()
	}
	ErrActivityTypeIsEmpty = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeActivityTypeEmpty, "Activity type can't be empty").C(ctx).Err()
	}
	ErrActivityOwnerIsEmpty = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeActivityOwnerEmpty, "Owner can not be empty").C(ctx).Err()
	}
	ErrActivityNotFound = func(ctx context.Context, id string) error {
		return er.WithBuilder(pb.ErrCodeActivityNotFound, "not found").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrActivityUserIdIsEmpty = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeActivityUserIdIsEmpty, "empty user id").C(ctx).Err()
	}
	ErrActivityFamilyIdIsEmpty = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeActivityFamilyIdIsEmpty, "empty family id").C(ctx).Err()
	}
	ErrActivityNameIsEmpty = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeActivityNameIsEmpty, "empty activity name").C(ctx).Err()
	}
	ErrActivityDescriptionIsEmpty = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeActivityDescriptionIsEmpty, "empty activity name").C(ctx).Err()
	}
	ErrActivityTimeIsEmpty = func(ctx context.Context) error {
		return er.WithBuilder(pb.ErrCodeActivityTimeEmpty, "time can not be empty").C(ctx).Err()
	}
	ErrActivityStorageCreate = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeActivityStorageCreate, "").C(ctx).Err()
	}
	ErrActivityStorageGetDb = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeActivityStorageGetDb, "").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrActivityStorageGetCache = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeActivityStorageGetCache, "").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrActivityStorageSetCache = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeActivityStorageSetCache, "").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrActivityStorageUpdate = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeActivityStorageUpdate, "").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrActivityDeleteFail = func(cause error, ctx context.Context, id string) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeActivityDeleteFail, "").F(er.FF{"id": id}).C(ctx).Err()
	}
	ErrActivityByOwnerSearch = func(cause error, ctx context.Context) error {
		return er.WrapWithBuilder(cause, pb.ErrCodeActivityByOwnerSearch, "").C(ctx).Err()
	}
)
