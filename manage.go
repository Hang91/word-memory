package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func InsertWord(w http.ResponseWriter, r *http.Request) {
	var word Word
	var message WordMessage
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request!")
		log.Println("Invalid request!")
		return
	}
	word.ID = bson.NewObjectId()
	word.Spell = message.Spell
	word.Level = message.Level
	word.Meaning = MeaningStringToMap(message.Meaning)
	log.Println(word)
	flag, err := dao.InsertWord(word)
	if !flag {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println("InsertWord function failed!")
		return
	}
	log.Println(word.Spell, "insert success!")
	RespondWithJson(w, http.StatusCreated, word)
}

func ModifyWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetWords(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetModifiedWords(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetDeletedWords(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetSearchedWords(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
