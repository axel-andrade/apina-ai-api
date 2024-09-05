package mongo_repositories

import (
	"fmt"

	"github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/mongo/mappers"
	"github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/mongo/models"
	"github.com/axel-andrade/opina-ai-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContentMongoRepository struct {
	Base          *BaseMongoRepository
	ContentMapper mappers.ContentMapper
}

const collection = "contents"

func BuildContentMongoRepository() *ContentMongoRepository {
	baseRepo := BuildBaseMongoRepository(collection)

	return &ContentMongoRepository{Base: baseRepo}
}

func (r *ContentMongoRepository) CreateContent(c *domain.Content) (*domain.Content, error) {
	model := r.ContentMapper.ToPersistence(*c)
	_, err := r.Base.Create(model)

	if err != nil {
		return nil, err
	}

	return r.ContentMapper.ToDomain(model), nil
}

func (r *ContentMongoRepository) FindContentByID(id string) (*domain.Content, error) {
	result := r.Base.FindByID(id)
	if result.Err() != nil {
		// Tratar erros, como document not found, aqui se necessário
		return nil, result.Err()
	}

	if result == nil {
		return nil, nil
	}

	var model models.Content

	if err := result.Decode(&model); err != nil {
		return nil, err
	}

	return r.ContentMapper.ToDomain(model), nil
}

func (r *ContentMongoRepository) UpdateContent(id string, c *domain.Content) error {
	// Convert the domain entity to a persistence model as a map
	updateModel := r.ContentMapper.ToUpdate(*c)

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Prepare the update document with $set
	updateData := bson.M{}

	// Iterate over the map and construct the update document
	for key, value := range updateModel {
		updateData[key] = value
	}

	// Perform the update operation
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updateData}
	_, err = r.Base.UpdateOne(filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("content not found")
		}

		return err
	}

	return nil
}
