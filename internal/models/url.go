package models

import (
	"fmt"
	"log"

	db "url-shortener/internal/database"
	"url-shortener/utils"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	urlColl   = "url"
	defaultDB = "url-shortener"
)

// User defines the user in db
type URL struct {
	ID        primitive.ObjectID `bson:"_id"        json:"id"`
	FullPath  string             `bson:"full_path"  json:"full_path"`
	ShortPath string             `bson:"short_path" json:"short_path"`
	Count     int                `bson:"count"      json:"count"`
}

// Add creates a user record in the database
func (url *URL) Add() (*URL, error) {
	url.Log("creating")
	client, err := db.NewClient()
	defer func() {
		if err := client.C.Disconnect(client.Ctx); err != nil {
			log.Println("client.C.Disconnect", err)
		}
	}()
	if err != nil {
		return nil, err
	}

	dbURL, found, err := url.Find()
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

	urlCollection := client.C.Database(defaultDB).Collection(urlColl)
	if _, err := urlCollection.InsertOne(client.Ctx, bsonURL); err != nil {
		return nil, err
	}

	return url, nil
}

// Find searches the users collection using the email as key
func (url *URL) Find() (*URL, bool, error) {
	url.Log("searching")
	client, err := db.NewClient()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		return nil, false, fmt.Errorf("db.NewClient err: %v", err)
	}

	urlsCollection := client.C.Database(defaultDB).Collection(urlColl)
	cursor, err := urlsCollection.Find(client.Ctx, bson.M{"full_path": url.FullPath})
	if err != nil {
		return nil, false, fmt.Errorf("urlsCollection.Find err: %v", err)
	}
	urls := []*URL{}
	for cursor.Next(client.Ctx) {
		u := &URL{}
		err = cursor.Decode(&u)
		if err != nil {
			return nil, false, fmt.Errorf("cursor.Decode err: %v", err)
		}
		urls = append(urls, u)
	}

	if len(urls) == 0 {
		return nil, false, nil
	}

	return urls[0], true, nil
}

// Log ...
func (url URL) Log(header string) {
	if len(header) > 0 {
		fmt.Printf("%s - url\n", header)
	}
	log.Printf("url.ID --------- %s", url.ID.String())
	log.Printf("url.FullPath --- %s", url.FullPath)
	log.Printf("url.ShortPath -- %s", url.ShortPath)
	log.Printf("url.Count ------ %d", url.Count)
}
