package storage

import (
	"git.jetbrains.space/orbi/fcsd/auth/internal/domain"
)

func (s *sessionStorageImpl) toSessionDto(t *domain.Session) *session {
	s.l().Mth("toSessionDto")
	if t == nil {
		return nil
	}

	dto := &session{
		Id:            t.Id,
		UserId:        t.UserId,
		DeviceName:    t.DeviceName,
		ClientVersion: t.ClientVersion,
		RefreshToken:  t.RefreshToken,
		FcmToken:      t.FCMToken,
		CreatedAt:     t.CreatedAt,
	}

	return dto
}

func (s *sessionStorageImpl) toSessionDomain(dto *session) *domain.Session {
	s.l().Mth("toSessionDomain")
	if dto == nil {
		return nil
	}
	sess := &domain.Session{
		Id:            dto.Id,
		UserId:        dto.UserId,
		DeviceName:    dto.DeviceName,
		ClientVersion: dto.ClientVersion,
		RefreshToken:  dto.RefreshToken,
		FCMToken:      dto.FcmToken,
		CreatedAt:     dto.CreatedAt,
	}
	return sess
}

func (s *sessionStorageImpl) toSessionArrayDomain(rq []*session) []*domain.Session {
	s.l().Mth("toSessionArrayDomain")
	var result []*domain.Session

	for _, value := range rq {
		result = append(result, &domain.Session{
			Id:            value.Id,
			UserId:        value.UserId,
			DeviceName:    value.DeviceName,
			ClientVersion: value.ClientVersion,
			RefreshToken:  value.RefreshToken,
			FCMToken:      value.FcmToken,
			CreatedAt:     value.CreatedAt,
		})
	}

	return result
}
