package infrastructure

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/yutaroyoshikawa/tipper-api/domain"
	"google.golang.org/api/iterator"
)

type Database struct {
	Firestore                *firestore.Client
	performanceCollectionRef *firestore.CollectionRef
	userCollectionRef        *firestore.CollectionRef
}

func NewDatabase(firebase *firebase.App) *Database {
	client, err := firebase.Firestore(context.Background())

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	f := new(Database)

	f.Firestore = client
	f.performanceCollectionRef = client.Collection("performances")
	f.userCollectionRef = client.Collection("users")

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

func (d *Database) CreatePerformance(performance domain.CreatePerfomanceInput) domain.Performance {
	docRef, _, err := d.performanceCollectionRef.Add(context.Background(), performance)

	if err != nil {
		log.Fatalf("error firestore: %v\n", err)
	}

	var registerPerformance domain.Performance

	snapshot, err := docRef.Get(context.Background())

	err = snapshot.DataTo(&registerPerformance)

	if err != nil {
		log.Fatalf("error firestore: %v\n", err)
	}

	return registerPerformance
}

func (d *Database) GetUserByUID(UId string) domain.User {
	var err error
	snapshot, err := d.userCollectionRef.Doc(UId).Get(context.Background())

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

func (d *Database) GetUserByUserID(userId string) (*domain.User, error) {
	iter := d.userCollectionRef.Where("id", "==", userId).Documents(context.Background())

	var user *domain.User

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		err = doc.DataTo(&user)

		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (d *Database) UpdateUser(userUId string, user domain.UpdateUserInput) error {
	_, err := d.userCollectionRef.Doc(userUId).Set(context.Background(), user)

	if err != err {
		return err
	}

	return nil
}

func (d *Database) UpdateUserID(UId string, newId string) error {
	_, err := d.userCollectionRef.Doc(UId).Set(context.Background(), map[string]string{
		"id": newId,
	})

	if err != nil {
		return err
	}

	return nil
}
