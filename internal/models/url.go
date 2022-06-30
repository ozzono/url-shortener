package models

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	ID        primitive.ObjectID `bson:"_id"       json:"id"`
	Source    string             `bson:"source"    json:"source"`
	Shortened string             `bson:"shortened" json:"shortened"`
	Count     int                `bson:"count"     json:"count"`
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
	log.Printf("url.Source ----- %s", url.Source)
	log.Printf("url.Shortened -- %s", url.Shortened)
	log.Printf("url.Count ------ %d", url.Count)
}
