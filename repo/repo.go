package repo

import (
	"context"
	"fmt"

	"github.com/NguyenHoaiPhuong/betta-house/utils"
	_ "github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	_ "github.com/mongodb/mongo-go-driver/mongo"
	_ "github.com/mongodb/mongo-go-driver/mongo/options"
)

// MongoDB include db client
type MongoDB struct {
	client *mongo.Client
}

// NewMongoDB returns a pointer to MongoDB
func NewMongoDB(serverHost string, port string) *MongoDB {
	connection := "mongodb://" + serverHost + ":" + port
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	utils.CheckError(err)

	err = client.Ping(context.TODO(), nil)
	utils.CheckError(err)

	fmt.Println(aurora.Red("Connected to MongoDB!"))
	return client
}

// Close the connection to server host
func (db *MongoDB) Close(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	utils.CheckError(err)

	fmt.Println(aurora.Red("Connection to MongoDB closed."))
}

// ConnectToCollection : function
func ConnectToCollection(client *mongo.Client, dbName string, colName string) *mongo.Collection {
	return client.Database(dbName).Collection(colName)
}

// InsertOneDocument : function
func InsertOneDocument(collection *mongo.Collection, document interface{}) {
	insertResult, err := collection.InsertOne(context.TODO(), document)
	utils.CheckError(err)
	fmt.Println(aurora.Green("Inserted a single document: "), aurora.Green(insertResult.InsertedID))
}

// InsertManyDocuments : function
func InsertManyDocuments(collection *mongo.Collection, documents []interface{}) {
	insertManyResult, err := collection.InsertMany(context.TODO(), documents)
	utils.CheckError(err)
	fmt.Println(aurora.Green("Inserted many documents: "), aurora.Green(insertManyResult.InsertedIDs))
}
