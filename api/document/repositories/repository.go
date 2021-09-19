package repositories

import (
	"context"
	"seafarer-backend/api/document/interfaces"
	"seafarer-backend/domain/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AFERepository struct {
	MongoDatabase *mongo.Database
}

func NewAFERepository(mongoDatabase *mongo.Database) interfaces.IAFERepository {
	return &AFERepository{MongoDatabase: mongoDatabase}
}

func (repo AFERepository) IsIDExist(id string) (isExist bool, err error) {

	count, err := repo.MongoDatabase.Collection(models.NewAFE().GetCollection()).CountDocuments(context.Background(), bson.D{{"_id", id}})
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, err
	}

	return true, err
}

func (repo AFERepository) Add(afe *models.AFE) (err error) {
	if _, err = repo.MongoDatabase.Collection(afe.GetCollection()).InsertOne(context.Background(), afe); err != nil {
		return err
	}

	return err
}

func (repo AFERepository) Update(afe *models.AFE) (err error) {
	if _, err = repo.MongoDatabase.Collection(afe.GetCollection()).UpdateOne(context.Background(), bson.D{{"_id", afe.UserID}}, bson.D{{"$set", afe}}); err != nil {
		return err
	}

	return err
}

func (repo AFERepository) Read(IDUser string) (afe models.AFE, err error) {
	if err = repo.MongoDatabase.Collection(afe.GetCollection()).FindOne(context.Background(), bson.M{"_id": IDUser}).Decode(&afe); err != nil {
		return afe, err
	}
	return afe, err
}
