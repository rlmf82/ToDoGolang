package controller

import (
	"ToDoProject/view"
	"encoding/json"
	"fmt"
	"net/http"
)

func Ping() http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {
			data := view.Response{
				Code: http.StatusOK,
				Body: "Pong",
			}
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(data)
		}

		fmt.Println(request.Method)
	}
}
