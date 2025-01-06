package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// infra:movies-persistence
// collection:movies

type man struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}

type movie struct {
	ID            bson.ObjectId `bson:"_id"`
	skey          string        `bson:"skey"`
	filename      string        `bson:"filename"`
	title         string        `bson:"title"`
	playtime      string        `bson:"playtime"`
	photodatetime string        `bson:"photodatetime"`
}

func main() {
	session, err := mgo.Dial("mongodb://localhost/")
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	c := session.DB("persistence").C("movies")

	if err := c.Insert(movie{
		ID:            bson.NewObjectId(),
		skey:          "292941",
		filename:      "moV265.mp4",
		title:         "たいとる",
		playtime:      "39",
		photodatetime: "55603",
	}); err != nil {
		log.Fatal(err)
	}
	var movies []movie
	if err := c.Find(nil).All(&movies); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", movies)
}

func pattern01(session *mgo.Session) {
	c := session.DB("some-infra-name").C("man")

	if err := c.Insert(man{
		ID:   bson.NewObjectId(),
		Name: "Taro",
	}, man{
		ID:   bson.NewObjectId(),
		Name: "Jiro",
	}); err != nil {
		log.Fatal(err)
	}

	var men []man
	if err := c.Find(nil).All(&men); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", men)
}
