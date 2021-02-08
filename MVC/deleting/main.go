package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jahanbabu-developer/gomodels.git/controller"
	"github.com/jahanbabu-developer/gomodels.git/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
