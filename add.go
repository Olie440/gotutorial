package gotutorial

import (
	"fmt"
	"net/http"
	"strconv"
)

func Add(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.ParseInt(r.URL.Query().Get("a"), 10, 64)
	b, err := strconv.ParseInt(r.URL.Query().Get("b"), 10, 64)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, a+b)
}
