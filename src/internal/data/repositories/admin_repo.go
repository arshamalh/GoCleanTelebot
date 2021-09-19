package repositories

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
)

// AdminDB interface
type AdminDB interface {
	Add(context.Context, entities.Admin) (string, error)
	GetByTID(context.Context, string) (*entities.Admin, error)
}

// NewAdminRepository make new admin repository
func NewAdminRepository(db AdminDB) domain.AdminRepository {
	return &adminRepo{db}
}

type adminRepo struct {
	ds AdminDB
}

func (r *adminRepo) Add(ctx context.Context, a entities.Admin) (string, error) {
	id, err := r.ds.Add(ctx, a)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *adminRepo) GetByTID(ctx context.Context, tid string) (*entities.Admin, error) {
	return r.ds.GetByTID(ctx, tid)
}
