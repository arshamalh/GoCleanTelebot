package authing

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"github.com/pkg/errors"
	"time"
)

// Service interface
type Service interface {
	Authorize(context.Context, string) (bool, error)
	NewAdmin(context.Context, string) error
}

// NewService make new authing service
func NewService(repo domain.AdminRepository) Service {
	return &service{repo}
}

type service struct {
	repo domain.AdminRepository
}

func (svc *service) Authorize(ctx context.Context, tid string) (bool, error) {
	admin, err := svc.repo.GetByTID(ctx, tid)
	if err != nil || admin == nil {
		return false, errors.Wrap(err, "internal.domain.authing")
	}
	return admin.Active, nil
}

func (svc *service) NewAdmin(ctx context.Context, tid string) error {
	_, err := svc.repo.Add(ctx, entities.Admin{
		ID:        "",
		TID:       tid,
		CreatedAt: time.Now(),
		Active:    true,
	})
	if err != nil {
		return errors.Wrap(err, "internal.domain.authing")
	}
	return nil
}
