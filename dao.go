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

// Search word
func (u *UsersDAO) SearchWordBySpell(spell string) (Word, error) {
	var word Word
	err := db.C(COLLECTION_WORD).Find(bson.M{"spell": spell}).One(&word)
	return word, err
}

// Search word by id
func (u *UsersDAO) SearchWordById(id bson.ObjectId) (Word, error) {
	var word Word
	err := db.C(COLLECTION_WORD).FindId(id).One(&word)
	return word, err
}

// Insert word
func (u *UsersDAO) InsertWord(word Word) error {
	err := db.C(COLLECTION_WORD).Insert(&word)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Update word
func (u *UsersDAO) UpdateWord(word Word) error {
	err := db.C(COLLECTION_WORD).UpdateId(word.ID, &word)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Delete word
func (u *UsersDAO) DeleteWord(ID bson.ObjectId) error {
	err := db.C(COLLECTION_WORD).RemoveId(ID)
	if err != nil {
		log.Println(err)
		return err
	}
}
