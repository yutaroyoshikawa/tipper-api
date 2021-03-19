package infrastructure

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/yutaroyoshikawa/tipper-api/domain"
)

type Database struct {
	Firestore *firestore.Client
}

func NewDatabase(firebase *firebase.App) *Database {
	client, err := firebase.Firestore(context.Background())

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	f := new(Database)

	f.Firestore = client

	return f
}

func (d *Database) GetPerformance(performanceId string) domain.Performance {
	snapshot, err := d.Firestore.Collection("performances").Doc(performanceId).Get(context.Background())

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	var performance domain.Performance

	snapshot.DataTo(&performance)

	return performance
}

func (d *Database) GetUser(userId string) domain.User {
	snapshot, err := d.Firestore.Collection("users").Doc(userId).Get(context.Background())

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	var user domain.User

	snapshot.DataTo(&user)

	return user
}
