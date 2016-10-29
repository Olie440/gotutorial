package gotutorial

import (
	"fmt"
	"net/http"
	"strconv"
)

func OutputXTimes(w http.ResponseWriter, r *http.Request) {
	interations, err := strconv.ParseInt(r.URL.Query().Get("interations"), 10, 64)

	if err != nil || interations > 1000 {
		fmt.Fprint(w, "Invalid value for interations, please pass in a number between 0 - 1000")
		return
	}

	for interations > 0 {
		fmt.Fprintf(w, "Hello World %d \n", interations)
		interations--
	}
}
