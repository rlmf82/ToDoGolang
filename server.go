package main

import (
	"awesomeProject/controller"
	"awesomeProject/model"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening localhost:3000")

	db := model.Connect()
	defer db.Close()
	http.ListenAndServe("localhost:3000", controller.Register())
}
