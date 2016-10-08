package gotutorial

import (
  "fmt"
  "net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
  fmt.Printf("Request from: %s \n", r.RemoteAddr)
}
