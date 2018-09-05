package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func InsertWord(w http.ResponseWriter, r *http.Request) {
	var word Word
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&word)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request!")
		log.Println("Invalid request!")
		return
	}
	log.Println(word)
	daoErr := dao.InsertWord(word)
	if daoErr != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println("InsertWord function failed!")
		return
	}
	log.Println(word.Spell, "insert success!")
	RespondWithJson(w, http.StatusCreated, word)
}

func ModifyWord(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating word...")
	var word Word
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&word)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request!")
		log.Println("Invalid request!")
		return
	}
	log.Println(word)
	daoErr := dao.UpdateWord(word)
	if daoErr != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println("UpdateWord function failed!")
		return
	}
	log.Println(word.Spell, "update success!")
	RespondWithJson(w, http.StatusOK, word)
}

func DeleteWord(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting word...")
	type Id struct {
		ID bson.ObjectId `bson:"_id" json:"id"`
	}
	var id Id
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request!")
		log.Println("Invalid request!")
		return
	}
	word, err := dao.SearchWordById(id.ID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request!")
		log.Println("Cannot find word!")
		return
	}
	daoErr := dao.DeleteWord(id.ID)
	if daoErr != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println("DeleteWord function failed!")
		return
	}
	log.Println(word.Spell, "delete success!")
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func GetWords(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all words...")
	words, err := dao.GetAllWords()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println("DeleteWord function failed!")
		return
	}
	log.Println("Get all words success!")
	RespondWithJson(w, http.StatusOK, words)
}

// func GetModifiedWords(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "not implemented yet !")
// }

// func GetDeletedWords(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "not implemented yet !")
// }

// func GetSearchedWords(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "not implemented yet !")
// }
