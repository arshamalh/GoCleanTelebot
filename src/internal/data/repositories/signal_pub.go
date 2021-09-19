package repositories

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"github.com/pkg/errors"
)

// SignalPub interface
type SignalPub interface {
	Publish(ctx context.Context, s entities.Signal) error
}

// NewSignalPublisher make new signal publisher repository
func NewSignalPublisher(spub SignalPub) domain.SignalPublisher {
	return &signalPublisher{spub}
}

type signalPublisher struct {
	pub SignalPub
}

func (spub *signalPublisher) Publish(ctx context.Context, s entities.Signal) error {
	err := spub.pub.Publish(ctx, s)
	if err != nil {
		return errors.Wrap(err, "internal.data")
	}
	return nil
}
