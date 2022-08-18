package controller

import (
	"ToDoProject/model"
	"ToDoProject/view"
	"encoding/json"
	"fmt"
	"net/http"
)

func Execute() http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodPost {
			fmt.Println("Insert new")
			data := view.PostRequest{}
			json.NewDecoder(request.Body).Decode(&data)

			error := model.Create(data.Name, data.Todo)

			if buildError(writer, error) {
				return
			}

			writer.WriteHeader(http.StatusCreated)
		} else if request.Method == http.MethodGet && len(request.URL.Query()) > 0 {
			fmt.Println("Read by name")
			name := request.URL.Query().Get("name")

			data, error := model.ReadByName(name)

			if buildError(writer, error) {
				return
			}

			json.NewEncoder(writer).Encode(data)
		} else if request.Method == http.MethodGet {
			fmt.Println("Read All")
			data, error := model.ReadAll()

			if buildError(writer, error) {
				return
			}

			json.NewEncoder(writer).Encode(data)
		} else if request.Method == http.MethodDelete {
			fmt.Println("Delete by name")

			name := request.URL.Path[1:]

			error := model.Delete(name)

			if buildError(writer, error) {
				return
			}

			writer.WriteHeader(http.StatusNoContent)
			json.NewEncoder(writer).Encode(struct {
				Status string `json:status`
			}{"Item deleted"})
		}
	}
}

func buildError(writer http.ResponseWriter, error error) bool {
	if error != nil {
		writer.Write([]byte(error.Error()))
		writer.WriteHeader(http.StatusBadRequest)
		return true
	}
	return false
}
