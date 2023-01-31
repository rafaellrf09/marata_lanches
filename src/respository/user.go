package respository

import (
	"context"
	"mlanches/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection}
}

func (repositoryUser UserRepository) FindAll() ([]models.User, error) {
	cursor, err := repositoryUser.collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return []models.User{}, err
	}

	users := []models.User{}

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return []models.User{}, err
	}

	defer cursor.Close(context.TODO())

	return users, nil
}

func (repositoryUser UserRepository) Create(user models.User) error {
	_, err := repositoryUser.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (repositoryUser UserRepository) FindOne(id string) (models.User, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	var result models.User

	err := repositoryUser.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, err
		}
	}
	return result, nil
}

func (repositoryUser UserRepository) Delete(id string) (uint16, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	result, err := repositoryUser.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return uint16(result.DeletedCount), nil
}

func (repositoryUser UserRepository) Update(id string, userUpdate models.User) (uint16, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: userUpdate}}

	result, err := repositoryUser.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	return uint16(result.ModifiedCount), nil
}
