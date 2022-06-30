package utils

import (
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	charSet = "_-abcdefghijklmnopqrstuwxyz"
)

func ToDoc(input interface{}) (bson.D, error) {
	data, err := bson.Marshal(input)
	if err != nil {
		return bson.D{}, fmt.Errorf("bson.Marshal err: %v", err)
	}

	doc := bson.D{}
	err = bson.Unmarshal(data, &doc)
	if err != nil {
		return bson.D{}, fmt.Errorf("bson.Unmarshal err: %v", err)
	}
	return doc, nil
}

func RString(min, max int) string {
	out := ""
	for i := 0; i < (RInt(max) + min); i++ {
		out += string(charSet[RInt(len(charSet))])
	}
	return out
}

func RInt(i int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(i)
}
