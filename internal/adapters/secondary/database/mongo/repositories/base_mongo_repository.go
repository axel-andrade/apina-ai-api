package mongo_repositories

import (
	"context"

	mongo_database "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/mongo"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseMongoRepository struct {
	collection *mongo.Collection
}

func BuildBaseMongoRepository(collectionName string) *BaseMongoRepository {
	db := mongo_database.GetDB()
	collection := db.Collection(collectionName)

	return &BaseMongoRepository{collection: collection}
}

func (r *BaseMongoRepository) StartTransaction() error {
	// Not implemented. Maybe use uow pattern for transactions in mongodb
	return nil
}

func (r *BaseMongoRepository) CommitTransaction() error {
	// Not implemented. Maybe use uow pattern for transactions in mongodb
	return nil
}

func (r *BaseMongoRepository) CancelTransaction() error {
	// Not implemented. Maybe use uow pattern for transactions in mongodb
	return nil
}

func (r *BaseMongoRepository) NextEntityID() string {
	return uuid.NewV4().String()
}

func (r *BaseMongoRepository) Create(data any) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.Background(), data)
}

func (r *BaseMongoRepository) FindOne(filter any) *mongo.SingleResult {
	return r.collection.FindOne(context.Background(), filter)
}

func (r *BaseMongoRepository) FindByID(id string) *mongo.SingleResult {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	result := r.collection.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		// Handle other potential errors here if needed
	}
	return result
}
func (r *BaseMongoRepository) ReplaceOne(filter any, replacement any) (*mongo.UpdateResult, error) {
	return r.collection.ReplaceOne(context.Background(), filter, replacement, options.Replace().SetUpsert(false))
}

func (r *BaseMongoRepository) UpdateOne(filter any, update any) (*mongo.UpdateResult, error) {
	return r.collection.UpdateOne(context.Background(), filter, update)
}
