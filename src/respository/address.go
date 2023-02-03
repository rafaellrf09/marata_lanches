package respository

import (
	"context"
	"mlanches/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AddressRepository struct {
	collection *mongo.Collection
}

func NewAddessRepository(collection *mongo.Collection) *AddressRepository {
	return &AddressRepository{collection}
}

func (repositoryAddress AddressRepository) FindAll(userId string) ([]models.Address, error) {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return []models.Address{}, err
	}

	filter := bson.D{{Key: "userid", Value: objectId}}

	cursor, err := repositoryAddress.collection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return []models.Address{}, err
	}

	addresses := []models.Address{}

	for cursor.Next(context.TODO()) {
		var address models.Address
		err := cursor.Decode(&address)
		if err != nil {
			return []models.Address{}, err
		}
		addresses = append(addresses, address)
	}
	if err := cursor.Err(); err != nil {
		return []models.Address{}, err
	}

	defer cursor.Close(context.TODO())

	return addresses, nil
}

func (repositoryAddress AddressRepository) Create(address models.Address) error {
	_, err := repositoryAddress.collection.InsertOne(context.TODO(), address)
	if err != nil {
		return err
	}
	return nil
}

func (repositoryAddress AddressRepository) FindOne(id string) (models.Address, error) {
	return models.Address{}, nil
}
func (repositoryAddress AddressRepository) Delete(id string) (uint16, error) {
	return 0, nil
}
func (repositoryAddress AddressRepository) Update(id string, updateAddress models.Address) (uint16, error) {
	return 0, nil
}
