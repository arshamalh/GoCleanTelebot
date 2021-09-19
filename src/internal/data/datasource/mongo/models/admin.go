package models

import (
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// AdminMongoModel struct
type AdminMongoModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	TID       string
	Active    bool
	CreatedAt time.Time `bson:"created_at"`
}

// FromDomainAdmin will make appropriate admin for other usages
func FromDomainAdmin(s entities.Admin) *AdminMongoModel {
	return &AdminMongoModel{
		ID:        primitive.NewObjectID(),
		TID:       s.TID,
		Active:    s.Active,
		CreatedAt: s.CreatedAt,
	}
}
