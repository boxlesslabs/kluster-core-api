//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================
 
/**
 * Define Mongo database connection
 * Return Existing collections for db operations
 **
 * @struct MongoDB
**/

package db

import (
	"github.com/klusters-core/api/config/secrets"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Connect() (StartMongoClient) {
	clientOptions := options.Client().ApplyURI(secrets.GetSecrets().DatabaseURL)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected")

	return &mongoDB{Client:client}
}

type StartMongoClient interface {
	SetCollection(col string) *mongo.Collection
	DropCollection(col string)
}

type mongoDB struct {
	Client 		*mongo.Client
}

func (db *mongoDB) DropCollection(col string) {
	err := db.Client.Database(secrets.GetSecrets().DatabaseName).Collection(col).Drop(nil)
	if err != nil {
		log.Println(err)
	}
}

func (db *mongoDB) SetCollection(col string) *mongo.Collection {
	return db.Client.Database(secrets.GetSecrets().DatabaseName).Collection(col)
}
