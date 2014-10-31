package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ralreegorganon/trivia"
)

func main() {
	var db trivia.DB
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}
	server := trivia.NewHTTPServer(&db)
	router, err := trivia.CreateRouter(server)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Handle("/", router)

	u := "0.0.0.0:8989"
	log.Printf("Trivia server started at http://%s\n", u)
	err = http.ListenAndServe(u, nil)
	if err != nil {
		fmt.Println(err)
	}
}
