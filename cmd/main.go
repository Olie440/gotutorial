package main

import (
    "github.com/olie440/gotutorial"
    "net/http"
)

func main() {
    http.HandleFunc("/", gotutorial.HelloWorld)
    http.HandleFunc("/rand", gotutorial.RandomNumber)
    http.ListenAndServe(":8080", nil)
}
