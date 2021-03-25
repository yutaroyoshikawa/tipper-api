package infrastructure

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/yutaroyoshikawa/tipper-api/domain"
)

type Database struct {
	Firestore                *firestore.Client
	performanceCollectionRef *firestore.CollectionRef
}

func NewDatabase(firebase *firebase.App) *Database {
	client, err := firebase.Firestore(context.Background())

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	f := new(Database)

	f.Firestore = client
	f.performanceCollectionRef = client.Collection("performances")

	return f
}

func (d *Database) UpdatePerformance(performanceId string, performance domain.Performance) domain.Performance {
	_, err := d.performanceCollectionRef.Doc(performanceId).Set(context.Background(), performance)

	if err != nil {
		log.Fatalf("error firestore: %v\n", err)
	}

	return performance
}

func (d *Database) GetPerformance(performanceId string) domain.Performance {
	var err error
	snapshot, err := d.performanceCollectionRef.Doc(performanceId).Get(context.Background())

	if err != nil {
		log.Fatalf("error firestore: %v\n", err)
	}

	var performance domain.Performance

	err = snapshot.DataTo(&performance)

	if err != nil {
		log.Fatalf("error firestore: %v\n", err)
	}

	return performance
}

func (d *Database) DeletePerformance(performanceId string) string {
	_, err := d.performanceCollectionRef.Doc(performanceId).Delete(context.Background())

	if err != nil {
		log.Fatalf("error firestore: %v\n", err)
	}

	return performanceId
}

func (d *Database) GetUserByUID(userId string) domain.User {
	var err error
	snapshot, err := d.Firestore.Collection("users").Doc(userId).Get(context.Background())

	if err != nil {
		log.Fatalf("error firestore: %v\n", err)
	}

	var user domain.User

	err = snapshot.DataTo(&user)

	if err != nil {
		log.Fatalf("error firestore: %v\n", err)
	}

	return user
}
