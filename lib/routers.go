package lib

import(
	"github.com/gorilla/mux"
	"net/http"
)

func Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("GET")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/home", home).Methods("GET")
	router.HandleFunc("/alluser",alluser).Methods("GET")
	router.Handle("/",router)
	return router
}