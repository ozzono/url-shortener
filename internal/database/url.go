package database

import (
	"url-shortener/internal/models"
	"url-shortener/utils"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	urlColl   = "url"
	defaultDB = "url-shortener"
)

func (client *Client) AddURL(url *models.URL) (*models.URL, error) {
	url.Log("creating")

	dbURL, found, err := client.FindURL(url)
	if err != nil {
		return dbURL, err
	}
	if found {
		return nil, nil
	}
	url.ID = primitive.NewObjectID()
	bsonURL, err := utils.ToDoc(url)
	if err != nil {
		return nil, errors.Wrap(err, "utils.ToDoc")
	}

	_, err = client.C.
		Database(defaultDB).
		Collection(urlColl).
		InsertOne(client.Ctx, bsonURL)
	if err != nil {
		return nil, err
	}

	return url, nil
}

// Find searches the users collection using the email as key
func (client *Client) FindURL(url *models.URL) (*models.URL, bool, error) {
	url.Log("searching")

	cursor, err := client.C.
		Database(defaultDB).
		Collection(urlColl).
		Find(client.Ctx, bson.M{"full_path": url.FullPath})
	if err != nil {
		return nil, false, errors.Wrap(err, "client.C.Database().Collection().Find()")
	}

	urls := []*models.URL{}
	for cursor.Next(client.Ctx) {
		u := &models.URL{}
		err = cursor.Decode(&u)
		if err != nil {
			return nil, false, errors.Wrap(err, "cursor.Decode")
		}
		urls = append(urls, u)
	}

	if len(urls) == 0 {
		return nil, false, nil
	}

	return urls[0], true, nil
}
