package publishing

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"github.com/pkg/errors"
)

// Service Publishing service interface
type Service interface {
	Publish(context.Context, []entities.Signal) error
}

// NewService return new instance of publishing service
func NewService(repo domain.SignalRepository, pub domain.SignalPublisher) Service {
	return &service{repo, pub}
}

type service struct {
	repo domain.SignalRepository
	pub  domain.SignalPublisher
}

func (svc *service) Publish(ctx context.Context, ss []entities.Signal) error {
	_, err := svc.repo.AddMany(ctx, ss)
	if err != nil {
		return errors.Wrap(err, "internal.domain.publishing")
	}

	for _, s := range ss {
		err = svc.pub.Publish(ctx, s)
		if err != nil {
			return errors.Wrap(err, "internal.domain.publishing")
		}
	}

	return nil
}
