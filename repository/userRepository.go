package repository

import (
	"context"
	"log"
	"puppy/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"puppy/database"
)

type UserRepository interface {
	GetUserList() ([]model.User, error)
	GetUserById(id string) (model.User, error)
	GetUserByUserNameAndPassword(username, password string) (model.User, error)
	InsertUser(user model.User) (string, error)
	UpdateUserById(user model.User) error
	DeleteUserById(id string) error
}

type userRepository struct {
	db database.Db
}

func NewUserRepository() UserRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	return userRepository{
		db: db,
	}
}

func (userRepository userRepository) GetUserList() ([]model.User, error) {

	userCollection := userRepository.db.GetUserCollection()

	cursor, err := userCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []model.User
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (userRepository userRepository) GetUserById(id string) (model.User, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, err
	}
	userCollection := userRepository.db.GetUserCollection()
	var userObject model.User
	// db.getCollection('users').find({"_id" : ObjectId("6297bdcc8d7757574658ed66")})
	err = userCollection.FindOne(context.TODO(), bson.D{
		{"_id", objectId},
	}).Decode(&userObject)

	if err != nil {
		return model.User{}, err
	}

	return userObject, nil

}

func (userRepository userRepository) GetUserByUserNameAndPassword(username, password string) (model.User, error) {

	userCollection := userRepository.db.GetUserCollection()
	var userObject model.User
	// db.getCollection('users').find({"_id" : ObjectId("6297bdcc8d7757574658ed66")})
	err := userCollection.FindOne(context.TODO(), bson.D{
		{"UserName", username},
		{"Password", password},
	}).Decode(&userObject)

	if err != nil {
		return model.User{}, err
	}

	return userObject, nil

}

func (userRepository userRepository) InsertUser(user model.User) (string, error) {
	userCollection := userRepository.db.GetUserCollection()
	res, err := userCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}
	objectId := res.InsertedID.(primitive.ObjectID).Hex()
	return objectId, nil
}

func (userRepository userRepository) UpdateUserById(user model.User) error {
	objectId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return err
	}
	user.Id = ""
	userCollection := userRepository.db.GetUserCollection()
	_, err = userCollection.UpdateOne(context.TODO(), bson.D{{"_id", objectId}}, bson.D{{"$set", user}})

	if err != nil {
		return err
	}

	return nil
}

func (userRepository userRepository) DeleteUserById(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	userCollection := userRepository.db.GetUserCollection()
	_, err = userCollection.DeleteOne(context.TODO(), bson.D{{"_id", objectId}})

	if err != nil {
		return err
	}

	return nil
}
