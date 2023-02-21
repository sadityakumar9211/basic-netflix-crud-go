package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sadityakumar9211/mongoapi/router"
)

const port = ":4000"

func main() {
	fmt.Println("Creating MongoDB APIs in Golang")

	fmt.Println(fmt.Sprintf("Server is Getting started at port %v", port))

	log.Fatal(http.ListenAndServe(port, router.Router()))
}
