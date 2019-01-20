package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

func SearchWord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("searching word \"", params["spell"], "\"...")
	word, err := dao.SearchWordBySpell(params["spell"])
	if err != nil {
		log.Println("word not found")
		RespondWithError(w, http.StatusBadRequest, "No such words")
		return
	}
	log.Println(word)
	RespondWithJson(w, http.StatusOK, word)
}

func UserLogIn(w http.ResponseWriter, r *http.Request) {
	var user1 User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user1)
	log.Println(user1.Email, user1.Password)

	if err != nil {
		log.Println("Invalid request!")
		return
	}
	user2 := dao.FindUserByEmail(user1.Email)
	// passwordLogIn := sha256.Sum256([]byte(password))
	// err := reflect.DeepEqual(user.Password, passwordLogIn)
	//err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	passwordLogIn, err1 := bcrypt.GenerateFromPassword([]byte(user1.Password), bcrypt.DefaultCost)
	if err1 != nil {
		log.Println("Password fault!")
	}
	passwordString := string(passwordLogIn[:])
	err2 := strings.Compare(user2.Password, passwordString)
	if err2 == -1 {
		log.Println("Password incorrect!")
		return
	} else {
		log.Println("Password correct!")
	}
	RespondWithJson(w, http.StatusOK, user2)

}

func UserLogOut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	log.Println(user)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request!")
		log.Println("Invalid request!")
		return
	}
	// passwordEncrypt := sha256.Sum256([]byte(user.Password))
	// user.PasswordEncrypt = PasswordEncrypt
	passwordEncrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	user.Password = string(passwordEncrypt[:])
	user.ID = bson.NewObjectId()
	daoErr := dao.RegisterUser(user)
	if daoErr != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println("Register failed!")
		return
	}
	log.Println(user.Email, "register success!")
	RespondWithJson(w, http.StatusCreated, user)

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

// func RegistrationStringToMap(inputStr string) map[string]string {
// 	partition := make(map[string]string])
// 	strs := strings.Split(inputStr, ";")
// 	for _, str := range strs {
// 		part := strings.Split(str, ":")
// 		partition[part[0]] = part[1]
// 	}

// 	user User := {
// 		email = partition[0],
// 		password = partition[1],
// 		username = partition[2],
// 		identity = partition[3]
// 	}

// 	return user
// }

// func DesEncrypt(origData, Key []byte) ([]byte, error) {
// 	var err error
// 	if err != nil {
// 		return nil, err
// 	}
// 	origData = PKCS5Padding(origData, block.BlockSize())
// 	blockMode := cipher.NewCBCEncrypter(block, key)
// 	crypted := make([]byte, len(origData))
// 	blockMode.CryptBlocks(crypted, origData)
// 	return crypted, nil
// }

// func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
// 	padding := blockSize - len(ciphertext)%blockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(ciphertext, padtext...)
// }

// func DesDecrypt(crypted, key []byte) {
// 	var err error
// 	block, error := des.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	blockMode := cipher.NewCBCDecrypter(block, key)
// 	origData := make([]byte, len(origData))
// 	blockMode.CryptBlocks(origData, crypted)
// 	origData = PKCS5UnPadding(origData)
// 	return origData, nil
// }

// func PKCS5UnPadding(origData []byte) []byte {
// 	length := len(origData)
// 	unpadding := int(origData[length-1])
// 	return origData[:(length - unpadding)]
// }
