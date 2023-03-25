package main

import (
	"log"
	"net/http"

	"github.com/Kdsingh333/HousewareHQ/routes"
)



func main() {
	router := routes.Router();
	err:=http.ListenAndServe(":8080",router);

	if err!= nil{
		log.Fatal(err);
		
	}
}