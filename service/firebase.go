package service

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
)

func InitializeFirebase(ctx context.Context) *firebase.App {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}
