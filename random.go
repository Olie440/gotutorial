package gotutorial

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
)

func RandomNumber(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    fmt.Fprintf(w, "My favorite number is %d", rand.Intn(100))
}
