package main

import(
	"jwt/lib"
	"net/http"
	"log"
	
)

func main(){
	router := lib.Router()
	log.Fatal(http.ListenAndServe(":"+lib.Services_config.Port ,router))
}

