// package firebase provides generic Firebase functionality.

package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
)

// FirestoreContext client and context for DB
type FirestoreContext struct {
	client *firestore.Client
	ctx    context.Context
}

// TODO maybe change to constructor function
// Initialize lets app communicate with a cloud Firestore database.
// configFilePath refers to account credentials config file created
// when setting up Firestore in browser.
func (f *FirestoreContext) Initialize(configFilePath string) error {
	f.ctx = context.Background()

	// Load service account credentials:
	serviceAccount := option.WithCredentialsFile(configFilePath)
	app, err := firebase.NewApp(f.ctx, nil, serviceAccount)
	if err != nil {
		// TODO handle another way?
		// TODO error descriptions
		log.Println(err)
		return err
	}

	// Instantiate firebase client:
	f.client, err = app.Firestore(f.ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Close closes the client
func (f *FirestoreContext) Close() {
	// Close down client:
	// TODO remove defer from function
	defer func() {
		err := f.client.Close()
		if err != nil {
			// TODO only log the error
			log.Fatal("Could not close Firebase client. Error:", err)
		}
	}()
}

// AddDocument stores a new document in a Firestore collection (by name).
// Returns autogenerated document ID.
// TODO rename, and create a function to create doc with specified id
// TODO interface{} as data input
func (f *FirestoreContext) AddDocument(collection string, document map[string]interface{}) (string, error) {
	id, _, err := f.client.Collection(collection).Add(f.ctx, document)
	if err != nil {
		log.Println("Duplicate document id generated")
		return "", err
	}
	return id.ID, nil
}

// DeleteDocument deletes data with a specific id from a collection.
func (f *FirestoreContext) DeleteDocument(collection, id string) error {
	_, err := f.client.Collection(collection).Doc(id).Delete(f.ctx)
	if err != nil {
		// TODO more specific log
		log.Println("Error while deleting")
	}
	// TODO return something if document does not exist?
	return nil
}

// ReadDocument reads a specific document by id.
// TODO return interface{} in separate function ReadToObject, usin pointer to object and dataTo()
func (f *FirestoreContext) ReadDocument(collection, id string) (map[string]interface{}, error) {
	documentSnap, err := f.client.Collection(collection).Doc(id).Get(f.ctx)
	if err != nil {
		return nil, err
	}
	document := documentSnap.Data()
	// TODO remove
	fmt.Println("document data: ", document)

	return document, nil
}

// CountDocuments counts all docs in specified collection
func (f *FirestoreContext) CountDocuments(collection string) (int, error) {
	count := 0
	iter := f.client.Collection(collection).Documents(f.ctx)

	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}
