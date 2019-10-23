package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Book struct {
	Id          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("booksDB3").C("books")
	c.RemoveAll(nil)
	index := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	err = c.Insert(
		&Book{bson.NewObjectId(), "Book1", "Golang Book"},
		&Book{bson.NewObjectId(), "Book2", "Golang Web Development"},
		&Book{bson.NewObjectId(), "Book3", "Golang Rest Api"},
	)
	if err != nil {
		panic(err)
	}
	result := Book{}
	err = c.Find(bson.M{"name": "Book1"}).One(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Description:", result.Description)
	}
}
