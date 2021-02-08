package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jahanbabu-developer/gomodels.git/model"
	"github.com/jahanbabu-developer/gomodels.git/view"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := view.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name ,data.Todo); err != nil {
				w.Write([]byte("some error"))
				return 
			}
                w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(data)
				
			} else if r.Method == http.MethodGet{
				data, err := model.ReadAll()
				if err != nil{
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			}
		}
	}

