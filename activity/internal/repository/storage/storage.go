package storage

import (
	"context"
	"encoding/json"
	"github.com/turbulent376/kit/log"
	"github.com/turbulent376/homeactivity/activity/internal/domain"
	"github.com/turbulent376/homeactivity/activity/internal/errors"
	"github.com/turbulent376/homeactivity/activity/internal/logger"
	"github.com/go-redis/redis"
	"time"
)

const (
	CacheKeyActivityId = "activity.id:"
	CacheKeyActivityTypeId     = "activity_type.id:"
)

type activity struct {
	Id       string    `gorm:"column:id"`
	Owner    string    `gorm:"column:owner"`
	Type     string    `gorm:"column:type"`
	Family   string    `gorm:"column:family"`
	DateFrom time.Time `gorm:"column:date_from"`
	DateTo   time.Time `gorm:"column:date_to"`
}

type activityType struct {
	Id          string    `gorm:"column:id"`
	Family      string    `gorm:"column:family"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
}

func (a *adapterImpl) l() log.CLogger {
	return logger.L().Cmp("activity-storage")
}

func (a *adapterImpl) setActivityCacheAsync(ctx context.Context, dto *activity) {

	go func() {

		l := a.l().Mth("set-cache").C(ctx).Dbg()

		keyId := CacheKeyActivityId + dto.Id

		j, err := json.Marshal(dto)
		if err != nil {
			l.E(err).St().Err()
		}
		dtoStr := string(j)

		// set cache for id key
		if err := a.container.Cache.Instance.Set(keyId, dtoStr, time.Hour).Err(); err != nil {
			l.E(errors.ErrActivityStorageSetCache(err, ctx, dto.Id)).St().Err()
		}
	}()
}

func (a *adapterImpl) CreateActivity(ctx context.Context, ac *domain.Activity) (*domain.Activity, error) {
	a.l().C(ctx).Mth("create")
	// save to DB
	if err := a.container.Db.Instance.Create(a.toActivityDto(ac)).Error; err != nil {
		return nil, errors.ErrActivityStorageCreate(err, ctx)
	}

	return ac, nil
}

func (a *adapterImpl) UpdateActivity(ctx context.Context, ac *domain.Activity) (*domain.Activity, error) {
	a.l().Mth("UpdateActivity").C(ctx).Dbg()

	// update DB
	if err := a.container.Db.Instance.Save(a.toActivityDto(ac)).Error; err != nil {
		return nil, errors.ErrActivityStorageUpdate(err, ctx, ac.Id)
	}

	// clear cache
	keys := []string{CacheKeyActivityId + ac.Id}
	a.container.Cache.Instance.Del(keys...)

	return ac, nil
}

func (a *adapterImpl) GetActivity(ctx context.Context, id string) (bool, *domain.Activity, error) {
	l := a.l().Mth("GetActivity").C(ctx).F(log.FF{"id": id}).Dbg()

	key := CacheKeyActivityId + id
	if j, err := a.container.Cache.Instance.Get(key).Result(); err == nil {
		// found in cache
		l.Dbg("found in cache")
		dto := &activity{}
		if err := json.Unmarshal([]byte(j), &dto); err != nil {
			return true, nil, err
		}
		return true, a.toActivityDomain(dto), nil
	} else {
		if err == redis.Nil {
			// not found in cache
			dto := &activity{Id: id}
			if res := a.container.Db.Instance.Limit(1).Find(&dto); res.Error == nil {
				l.DbgF("db: found %d", res.RowsAffected)
				if res.RowsAffected == 0 {
					return false, nil, nil
				} else {
					// set cache async
					a.setActivityCacheAsync(ctx, dto)
					return true, a.toActivityDomain(dto), nil
				}
			} else {
				return false, nil, errors.ErrActivityStorageGetDb(res.Error, ctx, id)
			}

		} else {
			return false, nil, errors.ErrActivityStorageGetCache(err, ctx, id)
		}
	}
}

func (a *adapterImpl) ListActivities(ctx context.Context, userId string) (bool, []*domain.Activity, error){
	a.l().Mth("ListActivities").C(ctx).F(log.FF{"id": userId}).Dbg()

	var dbres []*activity
	res := a.container.Db.Instance.Where("owner = ?", userId).Find(&dbres)
	if res.Error != nil {
		return false, nil, errors.ErrActivityByOwnerSearch(res.Error, ctx)
	}
	if res.RowsAffected == 0 {
		return false, nil, nil
	} else {
		return true, a.toActivitiesDomain(dbres), nil
	}

}

func (a *adapterImpl) ListActivitiesByFamily(ctx context.Context, familyId string) (bool, []*domain.Activity, error){
	a.l().Mth("ListActivities").C(ctx).F(log.FF{"id": familyId}).Dbg()

	var dbres []*activity
	res := a.container.Db.Instance.Where("family = ?", familyId).Find(&dbres)
	if res.Error != nil {
		return false, nil, errors.ErrActivityByOwnerSearch(res.Error, ctx)
	}
	if res.RowsAffected == 0 {
		return false, nil, nil
	} else {
		return true, a.toActivitiesDomain(dbres), nil
	}

}

func (a *adapterImpl) DeleteActivity(ctx context.Context, id string) error {
	a.l().C(ctx).Mth("DeleteActivity")
	// delete to DB
	key := CacheKeyActivityId + id

	a.container.Cache.Instance.Del(key)
	err := a.container.Db.Instance.Delete(&activity{}, id).Error
	if err != nil {
		return errors.ErrActivityDeleteFail(err, ctx, id)
	}
	return nil
}

func (a *adapterImpl) setActivityTypeCacheAsync(ctx context.Context, dto *activityType) {

	go func() {

		l := a.l().Mth("set-cache").C(ctx).Dbg()

		keyId := CacheKeyActivityTypeId + dto.Id

		j, err := json.Marshal(dto)
		if err != nil {
			l.E(err).St().Err()
		}
		dtoStr := string(j)

		// set cache for id key
		if err := a.container.Cache.Instance.Set(keyId, dtoStr, time.Hour).Err(); err != nil {
			l.E(errors.ErrActivityStorageSetCache(err, ctx, dto.Id)).St().Err()
		}
	}()
}

func (a *adapterImpl) CreateActivityType(ctx context.Context, at *domain.ActivityType) (*domain.ActivityType, error) {
	a.l().C(ctx).Mth("CreateActivityType")
	// save to DB
	if err := a.container.Db.Instance.Create(a.toActivityTypeDto(at)).Error; err != nil {
		return nil, errors.ErrActivityStorageCreate(err, ctx)
	}

	return at, nil
}

func (a *adapterImpl) UpdateActivityType(ctx context.Context, at *domain.ActivityType) (*domain.ActivityType, error) {
	a.l().Mth("UpdateActivityType").C(ctx).Dbg()

	// update DB
	if err := a.container.Db.Instance.Save(a.toActivityTypeDto(at)).Error; err != nil {
		return nil, errors.ErrActivityStorageUpdate(err, ctx, at.Id)
	}

	// clear cache
	keys := []string{CacheKeyActivityTypeId + at.Id}
	a.container.Cache.Instance.Del(keys...)

	return at, nil
}

func (a *adapterImpl) DeleteActivityType(ctx context.Context, id string) error {
	a.l().C(ctx).Mth("DeleteActivityType")
	key := CacheKeyActivityTypeId + id

	a.container.Cache.Instance.Del(key)
	err := a.container.Db.Instance.Delete(&activityType{}, id).Error
	if err != nil {
		return errors.ErrActivityDeleteFail(err, ctx, id)
	}
	return nil
}

func (a *adapterImpl) GetActivityType(ctx context.Context, id string) (bool, *domain.ActivityType, error) {
	l := a.l().Mth("GetActivityType").C(ctx).F(log.FF{"id": id}).Dbg()

	key := CacheKeyActivityTypeId + id
	if j, err := a.container.Cache.Instance.Get(key).Result(); err == nil {
		// found in cache
		l.Dbg("found in cache")
		dto := &activityType{}
		if err := json.Unmarshal([]byte(j), &dto); err != nil {
			return true, nil, err
		}
		return true, a.toActivityTypeDomain(dto), nil
	} else {
		if err == redis.Nil {
			// not found in cache
			dto := &activityType{Id: id}
			if res := a.container.Db.Instance.Limit(1).Find(&dto); res.Error == nil {
				l.DbgF("db: found %d", res.RowsAffected)
				if res.RowsAffected == 0 {
					return false, nil, nil
				} else {
					// set cache async
					a.setActivityTypeCacheAsync(ctx, dto)
					return true, a.toActivityTypeDomain(dto), nil
				}
			} else {
				return false, nil, errors.ErrActivityStorageGetDb(res.Error, ctx, id)
			}

		} else {
			return false, nil, errors.ErrActivityStorageGetCache(err, ctx, id)
		}
	}
}

func (a *adapterImpl) ListActivityTypes(ctx context.Context, familyId string) (bool, []*domain.ActivityType, error){
	a.l().Mth("ListActivityTypes").C(ctx).F(log.FF{"id": familyId}).Dbg()

	var dbres []*activityType
	res := a.container.Db.Instance.Where("family = ?", familyId).Find(&dbres)
	if res.Error != nil {
		return false, nil, errors.ErrActivityByOwnerSearch(res.Error, ctx)
	}
	if res.RowsAffected == 0 {
		return false, nil, nil
	} else {
		return true, a.toActivityTypesDomain(dbres), nil
	}

}