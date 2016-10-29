package gotutorial

import (
	"fmt"
	"net/http"
	"strconv"
)

func Add(w http.ResponseWriter, r *http.Request) {
	a, errA := strconv.ParseInt(r.URL.Query().Get("a"), 10, 64)
	b, errB := strconv.ParseInt(r.URL.Query().Get("b"), 10, 64)

	if errA != nil || errB != nil {
		fmt.Fprint(w, "Invalid input, please pass in two integers: a and b")
		return
	}

	fmt.Fprint(w, a+b)
}
