package routes

import (
	"github.com/Kdsingh333/HousewareHQ/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Start).Methods("POST")
	r.HandleFunc("/home", controller.Home).Methods("GET")
	r.HandleFunc("/Refresh", controller.Refresh).Methods("GET")
	r.HandleFunc("/Login", controller.Login).Methods("POST")
	r.HandleFunc("/Logout", controller.Logout).Methods("GET")
	r.HandleFunc("/Admin/add", controller.AdminAdd).Methods("POST")
	r.HandleFunc("/Admin/delete", controller.AdminDelete).Methods("DELETE")
	r.HandleFunc("/getUser", controller.GetUser).Methods("GET")
	return r
}
