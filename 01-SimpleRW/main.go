package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Id          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
}

func main() {
	//Connecting to MongoDB Server and Obtaining a Session
	session, err := mgo.Dial("localhost")     
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Accessing a MongoDB Collection
	c := session.DB("booksDB2").C("books")

    //create book object
	book1 := &Book{bson.NewObjectId(), "book1", "This is Book 1"}

	// inserting a single object
	err = c.Insert(book1)
	if err != nil{
		panic(err)
	}
	//inserting multiple object
    err = c.Insert(&Book{bson.NewObjectId(), "book2", "This is Book 2"},
				  &Book{bson.NewObjectId(), "book3", "This is Book 3"})

	//Query all the records
	result:= Book{}
	iter := c.Find(nil).Iter() // or c.Find(nil).Sort("field_name").Iter() for sorting records
	for iter.Next(&result){
		fmt.Printf("Book:%s, Description:%s\n", result.Name, result.Description)
	}

	// query the specific book 
	// result:= Book{}

	// err = c.Find(bson.M{"name":"book1"}).One(&result)
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Printf("Book Name is :%s, Description:%s\n", result.Name, result.Description)




}
