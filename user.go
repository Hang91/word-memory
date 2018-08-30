package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SearchWord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("search word", params["spell"], "...")
	flag, word := dao.SearchWord(params["spell"])
	if !flag {
		log.Println("word not found")
		respondWithError(w, http.StatusBadRequest, "No such words")
		return
	}
	respondWithJson(w, http.StatusOK, word.Meaning)
}

func UserLogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func UserLogOut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func RecommendArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func WordBattle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func GetBattleHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}
