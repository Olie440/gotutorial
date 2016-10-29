package gotutorial

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func RandomNumber(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	fmt.Fprintf(w, "My favorite number is %d", rand.Intn(100))
}
