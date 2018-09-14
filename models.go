package main

import (
	"gopkg.in/mgo.v2/bson"
)

//User is the stuct to record user information,
//user account password,
//user search history

type User struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	FirstName     string        `bson:"first_name" json:"first_name"`
	LastName      string        `bson:"last_name" json:"last_name"`
	SearchHistory []string      `bson:"search_history" json:"search_history"`
	Email         string        `bson:"email" json:"email"`
	Password      string        `bson:"password" json:"password"`
	//	PasswordEncrypt [32]byte
	Username string `bson:"username" json:"username"`
	Identity string `bson:"identity" json:"identity"`
	//identity: guest, user, vip, admin
}

// Word is the stuct to record word,
// meaning and part of speech of the word, level of the word
// Level of a word is a label to mark the frequency of searchcing,
// difficulty of a word
// Meaning is a map. The key of the map is part of speech,
// the value is possible meanings under this part.
// Definition is the english explanation of a word
type Word struct {
	ID         bson.ObjectId     `bson:"_id" json:"id"`
	Spell      string            `bson:"spell" json:"spell"`
	Phonetic   string            `bson:"phonetic" json:"phonetic"`
	Definition string            `bson:"definition" json:"definition"`
	Meaning    map[string]string `bson:"meaning" json:"meaning"`
	Level      int               `bson:"level" json:"level"`
}

// WordMessage is the struct to store message from api request
// of word information
// type WordMessage struct {
// 	ID      bson.ObjectId     `bson:"_id" json:"id"`
// 	Spell   string            `bson:"spell" json:"spell"`
// 	Meaning map[string]string `bson:"meaning" json:"meaning"`
// 	Level   int               `bson:"level" json:"level"`
// }
