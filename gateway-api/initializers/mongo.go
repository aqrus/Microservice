package initializers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
)

func ConnectToMongoDB() (*mgo.Session, error) {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to get env, %v", err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		config.MongoDBUser,
		config.MongoDBPassword,
		config.MongoDBHost,
		config.MongoDBPort,
		config.MongoDBName,
	)

	session, err := mgo.DialWithTimeout(mongoURI, 5*time.Second)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}

func NewMongoDB() (*mgo.Database, error) {
	session, err := ConnectToMongoDB()
	if err != nil {
		return nil, err
	}

	db := session.DB("")

	return db, nil
}

func WithMongoDB(ctx context.Context, f func(context.Context, *mgo.Database) error) error {
	db, err := NewMongoDB()
	if err != nil {
		return err
	}

	defer db.Session.Close()

	return f(ctx, db)
}
