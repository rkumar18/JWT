package lib

import(
	"encoding/json"
	"net/http"
	"jwt/helpers"
	"golang.org/x/crypto/bcrypt"
)


func signup(w http.ResponseWriter, r *http.Request){
	var input userData
	json.NewDecoder(r.Body).Decode(&input)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)
	result := dbsignup(&input)
	if result.Email ==""{
		helpers.WriteResponse(w, "successfully register", http.StatusOK)
	}else{
		helpers.WriteResponse(w, "on this email already an account", http.StatusBadRequest)
	}
}

func login(w http.ResponseWriter, r *http.Request){
	var input userData
	json.NewDecoder(r.Body).Decode(&input)
	result := dblogin(&input)
	compare := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(input.Password))
	if result.Email != "" && compare == nil{
		helpers.WriteResponse(w, "login successfully", http.StatusOK)
		token := helpers.Generatetoken(input.Email)
		helpers.WriteResponse(w, token, http.StatusOK)
	}else{
		helpers.WriteResponse(w, "wrong email and password", http.StatusBadRequest)
	}
}

func home(w http.ResponseWriter, r *http.Request){
	token := r.Header.Get("Authorization")
	validtoken := helpers.Validatetoken(token)
	if validtoken == true{
	 	helpers.WriteResponse(w, "hello world" , http.StatusOK)
	}else{
		helpers.WriteResponse(w," you cannot access this " , http.StatusBadRequest)
	}
}

func alluser(w http.ResponseWriter, r *http.Request){
	token := r.Header.Get("Authorization")
	validtoken := helpers.Validatetoken(token)
	if validtoken == true{
		alluser := dballuser()
		helpers.WriteResponse(w, alluser , http.StatusOK)
   }else{
	   helpers.WriteResponse(w," you cannot access this " , http.StatusBadRequest)
   }
}