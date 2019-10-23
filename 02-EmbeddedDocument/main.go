package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Task struct {
	Description string
	Due         time.Time
}
type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
	Tasks       []Task
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//get collection
	c := session.DB("taskdb").C("categories")
	//Embedding child collection
	doc := Category{
		bson.NewObjectId(),
		"Task Category",
		"Category open-source projects",
		[]Task{
			Task{" Create project in mgo", time.Date(2015, time.August, 10, 0, 0, 0, 0,
				time.UTC)},
			Task{"Create REST API", time.Date(2015, time.August, 20, 0, 0, 0, 0, time.UTC)},
		},
	}
	//insert a Category object with embedded Tasks
	err = c.Insert(&doc)
	if err != nil {
		log.Fatal(err)
	}
}
