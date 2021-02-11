package repository

import (
	"fmt"
	"github.com/hariszaki17/go-api-clean/exception"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/hariszaki17/go-api-clean/entity"
	"github.com/hariszaki17/go-api-clean/config"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// NewUserRepository global expose
func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepositoryImpl{
		Collection: database.Collection("user"),
	}
}

type userRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *userRepositoryImpl) Insert(user entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"username": user.Username,
		"password": user.Password,
		"role": user.Role,
	})
	exception.PanicIfNeeded(err)
}

func (repository *userRepositoryImpl) FindAll() (users []entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)
	var docs []bson.M
	err = cursor.All(ctx, &docs)
	exception.PanicIfNeeded(err)
	for _, doc := range docs {
		users = append(users, entity.User{
			ID:			doc["_id"].(interface{}),
			Username:	doc["username"].(string),
			Password:	doc["password"].(string),
			Role:		doc["role"].(string),
		})
		fmt.Println("docs <<<<<<<<<<<<<<<<", docs)
	}

	return users
}

func (repository *userRepositoryImpl) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}

// Hash expose global
func (repository *userRepositoryImpl) Encrypt(plainText string) (chiper string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 14)
	exception.PanicIfNeeded(err)
    return string(bytes)
}

func (repository *userRepositoryImpl) Decrypt(password, chiper string) {
	err := bcrypt.CompareHashAndPassword([]byte(chiper), []byte(password))
    exception.PanicIfNeeded(err)
}

func(repository *userRepositoryImpl) ValidatePassword(username, password string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	var user entity.User
	repository.Collection.FindOne(ctx, bson.M{ "username": username }).Decode(&user)
	repository.Decrypt(password, user.Password)
}