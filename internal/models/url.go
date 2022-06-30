package models

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	ID        primitive.ObjectID `bson:"_id"        json:"id"`
	FullPath  string             `bson:"full_path"  json:"full_path"`
	ShortPath string             `bson:"short_path" json:"short_path"`
	Count     int                `bson:"count"      json:"count"`
	Debug     bool
}

func (url URL) Log(header string) {
	if !url.Debug {
		return
	}
	if len(header) > 0 {
		fmt.Printf("%s - url\n", header)
	}
	log.Printf("url.ID --------- %s", url.ID.String())
	log.Printf("url.FullPath --- %s", url.FullPath)
	log.Printf("url.ShortPath -- %s", url.ShortPath)
	log.Printf("url.Count ------ %d", url.Count)
}
