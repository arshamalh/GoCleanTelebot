package mongo

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/data/datasource/mongo/models"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

// Mongo is an instance of MongoDB
type Mongo struct {
	db *mongodb.Database
}

// NewMongo Port new mongo client to other layers
func NewMongo(db *mongodb.Database) *Mongo {
	return &Mongo{db}
}

// Insert method to insert single signal to MongoDB
func (mng *Mongo) Insert(ctx context.Context, s entities.Signal) (string, error) {
	newSignal := models.FromDomainSignal(s)
	_, err := mng.db.Collection("signals").InsertOne(ctx, newSignal)
	if err != nil {
		fmt.Println("Error inserting to collection", err)
	}
	return "", nil
}

// InsertMany method to insert multiple signals to db, get argument as an array
func (mng *Mongo) InsertMany(ctx context.Context, sigs []entities.Signal) (string, error) {
	signalsColl := mng.db.Collection("signals")
	signals := make([]interface{}, len(sigs))
	for i, v := range sigs {
		signals[i] = models.FromDomainSignal(v)
	}
	if _, err := signalsColl.InsertMany(ctx, signals); err != nil {
		fmt.Println("Error inserting documents", err)
		return "", err
	}

	return "", nil
}

// Add method to insert new admin to db
func (mng *Mongo) Add(ctx context.Context, admin entities.Admin) (string, error) {
	newAdmin := models.FromDomainAdmin(admin)
	collection := mng.db.Collection("Admins")
	if res, e := collection.InsertOne(ctx, newAdmin); e != nil {
		fmt.Println("Error happened Inserting admin", e)
	} else {
		fmt.Println(res)
	}
	return "", nil
}

// GetByTID method to get admin by its telegram id
func (mng *Mongo) GetByTID(ctx context.Context, tid string) (*entities.Admin, error) {
	result := &entities.Admin{}
	collection := mng.db.Collection("Admins")
	err := collection.FindOne(ctx, bson.D{primitive.E{Key: "tid", Value: tid}}).Decode(result)
	if err != nil {
		fmt.Println("Error decoding result", err)
	}
	return result, err
}
