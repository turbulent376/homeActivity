package storage

import (
	"git.jetbrains.space/orbi/fcsd/auth/internal/domain"
	kitStorage "git.jetbrains.space/orbi/fcsd/kit/db"
)

func (u *userStorageImpl) toUserDto(t *domain.User) *user {
	u.l().Mth("toUserDto")
	if t == nil {
		return nil
	}

	dto := &user{
		BaseDto: kitStorage.BaseDto{
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			DeletedAt: t.DeletedAt,
		},
		Id:           t.Id,
		Email:        t.Email,
		Password:     t.Password,
		Name:         t.Name,
		Surname:      t.Surname,
		Avatar:       t.Avatar,
		FirebaseUUID: t.FirebaseUUID,
		KundelikId:   t.KundelikId,
		CountryCode:  t.CountryCode,
	}

	return dto
}

func (u *userStorageImpl) toUserDomain(dto *user) *domain.User {
	u.l().Mth("toUserDomain")
	if dto == nil {
		return nil
	}
	cl := &domain.User{
		Id:           dto.Id,
		Email:        dto.Email,
		Password:     dto.Password,
		Name:         dto.Name,
		Surname:      dto.Surname,
		Avatar:       dto.Avatar,
		FirebaseUUID: dto.FirebaseUUID,
		KundelikId:   dto.KundelikId,
		CountryCode:  dto.CountryCode,
	}
	return cl
}
