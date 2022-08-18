package main

import (
	"ToDoProject/controller"
	"ToDoProject/model"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening: localhost:3000")

	db := model.Connect()
	defer db.Close()

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", controller.Register()))
}
