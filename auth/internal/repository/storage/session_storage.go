package storage

import (
	"context"
	"encoding/json"
	"time"

	"git.jetbrains.space/orbi/fcsd/auth/internal/domain"
	"git.jetbrains.space/orbi/fcsd/auth/internal/errors"
	"git.jetbrains.space/orbi/fcsd/auth/internal/logger"
	"git.jetbrains.space/orbi/fcsd/kit/log"
	"github.com/go-redis/redis"
)

const (
	CacheSessionsUserKey = "sessions.userId:"
	CacheSessionKey      = "sessionId:"
)

type sessionStorageImpl struct {
	container *container
}

type session struct {
	Id            string    `gorm:"column:id"`
	UserId        string    `gorm:"column:user_id"`
	DeviceName    string    `gorm:"column:device_name"`
	ClientVersion string    `gorm:"column:client_version"`
	RefreshToken  string    `gorm:"column:refresh_token"`
	FcmToken      string    `gorm:"column:fcm_token"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}

func newSessionStorage(c *container) *sessionStorageImpl {
	return &sessionStorageImpl{
		container: c,
	}
}

func (s *sessionStorageImpl) l() log.CLogger {
	return logger.L().Cmp("session-storage")
}

func (s *sessionStorageImpl) setUserSessionsCacheAsync(ctx context.Context, sessions []*session, keyId string) {
	go func() {
		l := s.l().Mth("setUserSessionsCacheAsync").C(ctx).Dbg()

		j, err := json.Marshal(sessions)
		if err != nil {
			l.E(err).St().Err()
		}
		dtoStr := string(j)
		if err := s.container.Cache.Instance.Set(keyId, dtoStr, time.Hour).Err(); err != nil {
			l.E(errors.ErrUserStorageSetCache(err, ctx, keyId)).St().Err()
		}

	}()
}

func (s *sessionStorageImpl) setSessionCacheAsync(ctx context.Context, session *session, keyId string) {
	go func() {
		l := s.l().Mth("setSessionCacheAsync").C(ctx).Dbg()

		j, err := json.Marshal(session)
		if err != nil {
			l.E(err).St().Err()
		}
		dtoStr := string(j)
		if err := s.container.Cache.Instance.Set(keyId, dtoStr, time.Hour).Err(); err != nil {
			l.E(errors.ErrUserStorageSetCache(err, ctx, keyId)).St().Err()
		}
	}()
}

func (s *sessionStorageImpl) CreateSession(ctx context.Context, session *domain.Session) error {
	s.l().C(ctx).Mth("CreateSession")
	// save to DB
	if err := s.container.Db.Instance.Create(s.toSessionDto(session)).Error; err != nil {
		return errors.ErrSessionStorageCreate(err, ctx)
	}

	return nil
}

func (s *sessionStorageImpl) UpdateSession(ctx context.Context, session *domain.Session) error {
	s.l().Mth("UpdateSession").C(ctx).Dbg()

	// update DB
	if err := s.container.Db.Instance.Save(s.toSessionDto(session)).Error; err != nil {
		return errors.ErrSessionStorageUpdate(err, ctx, session.Id)
	}

	// clear cache
	s.container.Cache.Instance.Del(CacheSessionsUserKey + session.UserId)

	return nil
}

func (s *sessionStorageImpl) DeleteSession(ctx context.Context, id string) error {
	s.l().Mth("DeleteSession").C(ctx)
	dto := &session{Id: id}
	err := s.container.Db.Instance.Delete(&dto).Error
	if err != nil {
		return errors.ErrSessionStorageDelete(err, ctx)
	}
	// clear cache
	s.container.Cache.Instance.Del(CacheSessionsUserKey + id)

	return nil
}

func (s *sessionStorageImpl) GetSession(ctx context.Context, id string) (*domain.Session, error) {
	l := s.l().Mth("GetSession").C(ctx).F(log.FF{"id": id}).Dbg()

	key := CacheSessionKey + id
	if j, err := s.container.Cache.Instance.Get(key).Result(); err == nil {
		// found in cache
		l.Dbg("found in cache")
		dto := &session{}
		if err := json.Unmarshal([]byte(j), &dto); err != nil {
			return nil, errors.ErrSessionStorageGetCache(err, ctx, id)
		}
		return s.toSessionDomain(dto), nil
	} else {
		if err == redis.Nil {
			// not found in cache
			dto := &session{Id: id}
			if res := s.container.Db.Instance.Limit(1).Find(&dto); res.Error == nil {
				if res.RowsAffected == 0 {
					return nil, nil
				} else {
					s.setSessionCacheAsync(ctx, dto, key)
					return s.toSessionDomain(dto), nil
				}
			} else {
				return nil, errors.ErrSessionStorageGetDb(res.Error, ctx, id)
			}
		}
		return nil, errors.ErrSessionStorageGetCache(err, ctx, id)
	}

}

func (s *sessionStorageImpl) GetUserSessions(ctx context.Context, id string) ([]*domain.Session, error) {
	l := s.l().Mth("GetUserSessions").C(ctx).Dbg()
	var sessions []*session
	key := CacheSessionsUserKey + id
	if j, err := s.container.Cache.Instance.Get(key).Result(); err == nil {
		// found in cache
		l.Dbg("found in cache")
		if err := json.Unmarshal([]byte(j), &sessions); err != nil {
			return nil, errors.ErrSessionNotFound(err, ctx)
		}
		l.Inf(sessions)
		return s.toSessionArrayDomain(sessions), nil
	} else {
		if err == redis.Nil {
			// not found in cache
			if res := s.container.Db.Instance.Find(&sessions, "user_id = ?", id); res.Error == nil {
				l.DbgF("db: found %d", res.RowsAffected)
				if res.RowsAffected == 0 {
					return nil, nil
				} else {
					// set cache async
					s.setUserSessionsCacheAsync(ctx, sessions, key)
					l.Inf(sessions)
					return s.toSessionArrayDomain(sessions), nil
				}
			} else {
				return nil, errors.ErrSessionStorageGetDb(err, ctx, id)
			}
		} else {
			return nil, errors.ErrSessionStorageGetCache(err, ctx, id)
		}
	}

}
