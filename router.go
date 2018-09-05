package main

import (
	"github.com/gorilla/mux"
)

func CreateRouter(router *mux.Router) {

	//user api
	router.HandleFunc("/search/{spell}", SearchWord).Methods("GET")
	router.HandleFunc("/login", UserLogIn).Methods("POST")
	router.HandleFunc("/logout", UserLogOut).Methods("POST")
	router.HandleFunc("/signup", UserSignUp).Methods("POST")
	router.HandleFunc("/profile", GetProfile).Methods("GET")
	router.HandleFunc("/profile/change", ChangeProfile).Methods("PUT")
	router.HandleFunc("/profile/changepassword", ChangePassword).Methods("PUT")
	router.HandleFunc("/recommend", RecommendArticle).Methods("GET")
	router.HandleFunc("/battle", WordBattle).Methods("GET")
	router.HandleFunc("/battlehistory", GetBattleHistory).Methods("GET")

	//manage api
	router.HandleFunc("/manage/insert", InsertWord).Methods("POST")
	router.HandleFunc("/manage/modify", ModifyWord).Methods("PUT")
	router.HandleFunc("/manage/delete", DeleteWord).Methods("DELETE")
	router.HandleFunc("/manage/getwords", GetWords).Methods("GET")
	router.HandleFunc("/manage/getmodifedwords", GetModifiedWords).Methods("GET")
	router.HandleFunc("/manage/getdeletedwords", GetDeletedWords).Methods("GET")
	router.HandleFunc("/manage/getsearchwords", GetSearchedWords).Methods("GET")
}
