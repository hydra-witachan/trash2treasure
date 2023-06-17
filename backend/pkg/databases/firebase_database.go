package databases

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func NewFirebaseBucket() (bucket *storage.BucketHandle, err error) {
	ctx := context.Background()

	credentialsFile := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	app, err := firebase.NewApp(ctx, &firebase.Config{
		StorageBucket: os.Getenv("FIREBASE_BUCKET"),
	}, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return
	}

	storageApp, err := app.Storage(ctx)
	if err != nil {
		return
	}

	bucket, err = storageApp.DefaultBucket()
	return
}
