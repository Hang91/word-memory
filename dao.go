package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Users data access object
type UsersDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// mongodb collection for users
	COLLECTION_USER = "users"

	// mongodb collection for words
	COLLECTION_WORD = "words"
)

// Connect to database
func (u *UsersDAO) Connect() {
	log.Println("Start connect database server...")
	session, err := mgo.Dial(u.Server)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Database", u.Database, "connected.")
	db = session.DB(u.Database)
}

// Searchword
func (u *UsersDAO) SearchWord(spell string) (bool, Word) {
	var word Word
	err := db.C(COLLECTION_WORD).Find(bson.M{"spell": spell}).One(&word)
	if err != nil {
		log.Println(err)
		return false, word
	}
	return true, word
}

// Insert word
func (u *UsersDAO) InsertWord(word Word) (bool, error) {
	err := db.C(COLLECTION_WORD).Insert(&word)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}
