package models

import (
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SignalMongoModel struct
type SignalMongoModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	Pair        string             `bson:""`
	Date        string             `bson:""`
	ImageURL    string             `bson:""`
	TimeFrame   string             `bson:""`
	EntryPrice  string             `bson:""`
	TargetPrice string             `bson:"target_price"`
	StopLoss    string             `bson:""`
	Risk2Reward string             `bson:""`
	TradeVolume string             `bson:""`
}

// FromDomainSignal will make appropriate signal from signal entity to use somewhere else.
func FromDomainSignal(s entities.Signal) *SignalMongoModel {
	return &SignalMongoModel{
		ID:          primitive.NewObjectID(),
		Pair:        s.Pair,
		Date:        s.Date,
		ImageURL:    s.ImageURL,
		TimeFrame:   s.TimeFrame,
		EntryPrice:  s.EntryPrice,
		TargetPrice: s.TargetPrice,
		StopLoss:    s.StopLoss,
		Risk2Reward: s.Risk2Reward,
		TradeVolume: s.TradeVolume,
	}
}
