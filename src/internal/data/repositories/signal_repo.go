package repositories

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"github.com/pkg/errors"
)

// SignalDB interface
type SignalDB interface {
	Insert(context.Context, entities.Signal) (string, error)
	InsertMany(context.Context, []entities.Signal) (string, error)
}

// NewSignalRepository make new signal repo for db
func NewSignalRepository(db SignalDB) domain.SignalRepository {
	return &signalRepo{db}
}

type signalRepo struct {
	ds SignalDB
}

func (r *signalRepo) Add(ctx context.Context, s entities.Signal) (string, error) {
	id, err := r.ds.Insert(ctx, s)
	if err != nil {
		return "", errors.Wrap(err, "internal.data")
	}
	return id, nil
}

func (r *signalRepo) AddMany(ctx context.Context, ss []entities.Signal) (string, error) {
	ids, err := r.ds.InsertMany(ctx, ss)
	if err != nil {
		return "", errors.Wrap(err, "internal.data")
	}
	return ids, nil
}
