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

// Search word by spell
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
	return err
}

// Update word
func (u *UsersDAO) UpdateWord(word Word) error {
	err := db.C(COLLECTION_WORD).UpdateId(word.ID, &word)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

// Delete word
func (u *UsersDAO) DeleteWord(ID bson.ObjectId) error {
	err := db.C(COLLECTION_WORD).RemoveId(ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

// Get All Words
func (u *UsersDAO) GetAllWords() ([]Word, error) {
	var words []Word
	err := db.C(COLLECTION_WORD).Find(bson.M{}).All(&words)
	return words, err
}

// Register user
func (u *UsersDAO) RegisterUser(user User) error {
	//插入encrypt函数
	err := db.C(COLLECTION_USER).Insert(&user)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

// User LogIn
// func (u *UsersDAO) LogInUser(username string)（User, error) {
// 	var user User
// 	err := db.C(COLLECTION_USER).Find(bson.M{"username": username}).One(&user)
// 	if err != nil {
// 		log.Println("Username not found!")
// 		return user, err
// 	}

// 	return user, err

//err := db.C(COLLECTION_USER).Find(bson.M{"Password": password})
// err := strings.Compare(user.Password, password)
// if err == -1 {
// 	log.Println("Password incorrect!")
// 	return
// }
//}

// User LogIn
func (u *UsersDAO) LogInUser(username string) (User, error) {
	var user User
	err := db.C(COLLECTION_USER).Find(bson.M{"username": username}).One(&user)
	if err != nil {
		log.Println("Username not found!")
		return user, err
	}
	return user, err
}
