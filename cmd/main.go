package main

import (
  "github.com/olie440/gotutorial"
  "net/http"
)

func main() {
  http.HandleFunc("/", gotutorial.HelloWorld)
  http.ListenAndServe(":8080", nil)
}
