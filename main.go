package main

import (
	"log"
	"net/http"

	"../word-memory/manage"
	"../word-memory/user"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//user apit
	router.HandleFunc("/search", user.SearchWord).Methods("GET")
	router.HandleFunc("/login", user.UserLogIn).Methods("POST")
	router.HandleFunc("/logout", user.UserLogOut).Methods("POST")
	router.HandleFunc("/signup", user.UserSignUp).Methods("POST")
	router.HandleFunc("/profile", user.GetProfile).Methods("GET")
	router.HandleFunc("/profile/change", user.ChangeProfile).Methods("PUT")
	router.HandleFunc("/profile/changepassword", user.ChangePassword).Methods("PUT")
	router.HandleFunc("/recommend", user.RecommendArticle).Methods("GET")
	router.HandleFunc("/battle", user.WordBattle).Methods("GET")
	router.HandleFunc("/battlehistory", user.GetBattleHistory).Methods("GET")

	//manage api
	router.HandleFunc("/manage/insert", manage.InsertWord).Methods("POST")
	router.HandleFunc("/manage/modify", manage.ModifyWord).Methods("PUT")
	router.HandleFunc("/manage/delete", manage.DeleteWord).Methods("DELETE")
	router.HandleFunc("/manage/getwords", manage.GetWords).Methods("GET")
	router.HandleFunc("/manage/getmodifedwords", manage.GetModifiedWords).Methods("GET")
	router.HandleFunc("/manage/getdeletedwords", manage.GetDeletedWords).Methods("GET")
	router.HandleFunc("/manage/getsearchwords", manage.GetSearchedWords).Methods("GET")

	// start server
	log.Fatal(http.ListenAndServe(":4000", router))
}
