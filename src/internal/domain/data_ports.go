package domain

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
)

// SignalRepository interface
type SignalRepository interface {
	Add(context.Context, entities.Signal) (string, error)
	AddMany(context.Context, []entities.Signal) (string, error)
}

// AdminRepository interface
type AdminRepository interface {
	Add(context.Context, entities.Admin) (string, error)
	GetByTID(context.Context, string) (*entities.Admin, error)
}

// SignalPublisher interface
type SignalPublisher interface {
	Publish(ctx context.Context, s entities.Signal) error
}

// CommandRepository interface
type CommandRepository interface {
	InsertCmd(context.Context, entities.Command) (string, error)
	GetAllCmds(context.Context) ([]entities.Command, error)
	DeleteCmd(context.Context, string) (string, error)
}
