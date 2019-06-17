package repo

import (
	"context"
	"fmt"

	"github.com/NguyenHoaiPhuong/betta-house/utils"
	"github.com/logrusorgru/aurora"
	_ "github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	_ "github.com/mongodb/mongo-go-driver/mongo/options"
)

// MongoDB include db client
type MongoDB struct {
	client *mongo.Client
}

// NewMongoDB returns a pointer to MongoDB
func NewMongoDB(serverHost string, port string) *MongoDB {
	var err error
	db := new(MongoDB)

	connection := "mongodb://" + serverHost + ":" + port
	db.client, err = mongo.Connect(context.TODO(), connection)
	utils.CheckError(err)

	err = db.client.Ping(context.TODO(), nil)
	utils.CheckError(err)

	fmt.Println(aurora.Red("Connected to MongoDB!"))

	return db
}

// Close the connection to server host
func (db *MongoDB) Close() {
	err := db.client.Disconnect(context.TODO())
	utils.CheckError(err)

	fmt.Println(aurora.Red("Connection to MongoDB closed."))
}

// ConnectToCollection return the collection respective to the given db name and collection name
func (db *MongoDB) ConnectToCollection(dbName string, colName string) *mongo.Collection {
	return db.client.Database(dbName).Collection(colName)
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
