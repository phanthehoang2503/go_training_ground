package main

import (
	"fmt"
	"net/http"

	"github.com/phanthehoang2503/go_training_ground/go_web/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Listening on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
