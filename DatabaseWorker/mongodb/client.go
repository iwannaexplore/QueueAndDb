package mongodb

import (
	"QueueAndDb/pkg/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

// MongoDB configuration
const (
	uri            = "mongodb://localhost:27017" // Default MongoDB URI
	databaseName   = "mydatabase"                // Replace with your database name
	collectionName = "mycollection"              // Replace with your collection name
)

// ItemRepository interface for person data operations
type ItemRepository interface {
	Insert(item models.Item) (interface{}, error)
	FindByNumber(number int) (models.Item, error)
}

// MongoItemRepository struct to implement ItemRepository
type MongoItemRepository struct {
	collection *mongo.Collection
	client     *mongo.Client
}

// NewMongoItemRepository function to create a new MongoItemRepository
func NewMongoItemRepository() *MongoItemRepository {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pinged MongoDB successfully.")
	collection := client.Database(databaseName).Collection(collectionName)

	return &MongoItemRepository{collection: collection}
}

// Insert method to insert a person into the database
func (repo *MongoItemRepository) Insert(item models.Item) (interface{}, error) {
	insertResult, err := repo.collection.InsertOne(context.TODO(), item)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

// FindByNumber method to find a person by name
func (repo *MongoItemRepository) FindByNumber(number int) (models.Item, error) {
	var result models.Item
	filter := models.Item{NumberProperty: number}
	err := repo.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.Item{}, err
	}
	return result, nil
}

func (repo *MongoItemRepository) Close() {
	if err := repo.client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection closed.")
}
