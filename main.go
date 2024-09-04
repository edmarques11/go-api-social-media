package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("api is running")
	r := router.GenerateRoutes()

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
