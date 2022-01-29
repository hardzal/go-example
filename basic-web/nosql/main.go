package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type student struct {
	Name  string `bson:"Name"`
	Grade int    `bson:"Grade"`
}

func connect() (*mgo.Session, error) {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}

	return session, nil
}

func insert() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	defer session.Close()

	var collection = session.DB("golang_first").C("student")
	err = collection.Insert(&student{"John", 3}, &student{"Erwin", 6})
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Insert success!")
}

func find() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	defer session.Close()

	var collection = session.DB("golang_first").C("student")
	var result = student{}
	var selector = bson.M{"Name": "Jon"}
	err = collection.Find(selector).One(&result)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Name:", result.Name)
	fmt.Println("Grade:", result.Grade)
}

func update() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	defer session.Close()

	var collection = session.DB("golang_first").C("student")
	var selector = bson.M{"Name": "John"}
	var changes = student{"John Cenna", 10}
	err = collection.Update(selector, changes)

	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Update success")
}

func remove() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	defer session.Close()

	var collection = session.DB("golang_first").C("student")
	var selector = bson.M{"Name": "Erwin"}
	err = collection.Remove(selector)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Remove success!")
}

func main() {
	if _, err := connect(); err != nil {
		fmt.Println("Error!", err.Error())
	} else {
		fmt.Println("Mongo DB Connection success!")
	}
	// insert()
	// find()
	// update()
	// remove()

}
