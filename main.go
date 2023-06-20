package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/maazmaahi/mongoapi/routes"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening at port 8080 ...")
}
