package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jahanbabu-developer/gomodels.git/view"
)

func ping() http.HandlerFunc {

	return func(Wri http.ResponseWriter, Rea *http.Request) {

		if Rea.Method == http.MethodGet {
			data := view.Response{
				Code: http.StatusOK,
				Body: "Able to connect in server and test API",
			}
			json.NewEncoder(Wri).Encode(data)
		}

	}

}
