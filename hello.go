package gotutorial

import (
  "fmt"
  "net/http"
  "os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
  fmt.Fprintf(os.Stdout ,"Request from: %s \n", r.RemoteAddr)
}
