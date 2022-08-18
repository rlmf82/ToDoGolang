package controller

import (
	"awesomeProject/model"
	"awesomeProject/view"
	"encoding/json"
	"net/http"
)

func Create() http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodPost {
			data := view.PostRequest{}
			json.NewDecoder(request.Body).Decode(&data)

			error := model.Create(data.Name, data.Todo)

			if error != nil {
				writer.Write([]byte(error.Error()))
				writer.WriteHeader(http.StatusBadRequest)
				return
			}

			writer.WriteHeader(http.StatusCreated)
		}
	}
}
