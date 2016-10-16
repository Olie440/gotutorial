package main

import (
    "github.com/olie440/gotutorial"
    "net/http"
)

func main() {
    http.HandleFunc("/", gotutorial.HelloWorld)
    http.HandleFunc("/rand", gotutorial.RandomNumber)
    http.HandleFunc("/add", gotutorial.Add)
    http.HandleFunc("/sqrt", gotutorial.SquareRoot)
    http.ListenAndServe(":8080", nil)
}
