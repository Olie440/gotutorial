package gotutorial

import (
  "fmt"
  "math"
  "net/http"
  "net/url"
  "strconv"
)

func parseQuery(query url.Values) (x float64, err error) {
  x, err = strconv.ParseFloat(query.Get("x"), 64)
  return
}

func sqrt(x float64) string {
  if (x < 0) {
    return sqrt(-x) + "i"
  }

  return strconv.FormatFloat(math.Sqrt(x), 'f', 0, 64)
}

func SquareRoot(w http.ResponseWriter, r *http.Request) {
  x, err := parseQuery(r.URL.Query())

  if (err != nil) {
    fmt.Fprintln(w, err)
    return
  }

  fmt.Fprintf(w, "The square root of %.0f is %s", x, sqrt(x))
}
