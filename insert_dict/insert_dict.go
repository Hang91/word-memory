package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

const (
	// Server
	Server = "localhost:27017"

	// Database
	Database = "word_memory"

	// mongodb collection for users
	COLLECTION_USER = "users"

	// mongodb collection for words
	COLLECTION_WORD = "words"
)

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

var db *mgo.Database

func main() {
	log.Println("Insert dict to database...")
	log.Println("open ecdict.csv ...")
	file, err := os.Open("ecdict.csv")
	Connect()
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	for i := 0; ; i++ {
		log.Println(i)
		line, e := reader.Read()
		if e == io.EOF {
			break
		}
		if e != nil {
			log.Fatal(e)
		}
		if i == 0 {
			continue
		}
		var word Word
		word.ID = bson.NewObjectId()
		word.Spell = line[0]
		word.Phonetic = line[1]
		word.Definition = line[2]
		if IsWord(line[0]) {
			word.Meaning = MeaningStringToMap(line[3])
		} else {
			word.Meaning = map[string]string{"": line[3]}
		}
		word.Level = 1
		log.Println(word)
		InsertWord(word)
	}
	log.Println("Insert dict to database success!")
}

func Connect() {
	log.Println("Start connect database server...")
	session, err := mgo.Dial(Server)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Database", Database, "connected.")
	db = session.DB(Database)
}

// split string
// n. 鱼\nvi. 捕鱼\nvt. 捕鱼
func MeaningStringToMap(inputStr string) map[string]string {
	result := make(map[string]string)
	strs := strings.Split(inputStr, "\\n")
	for _, str := range strs {
		part := strings.SplitN(strings.TrimSpace(str), " ", 2)
		if len(part) == 1 {
			result[""] = part[0]
		} else {
			result[part[0]] = part[1]
		}
	}
	return result
}

// The string whose spell only contains letter is a word
func IsWord(inputStr string) bool {
	for _, ch := range inputStr {
		if !unicode.IsLetter(ch) {
			return false
		}
	}
	return true
}

// Insert word
func InsertWord(word Word) error {
	err := db.C(COLLECTION_WORD).Insert(&word)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
