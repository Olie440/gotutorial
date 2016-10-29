package gotutorial

import (
	"net/http/httptest"
	"testing"
)

func TestAddReturnsErrorsWhenAIsMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/add?b=1", nil)
	w := httptest.NewRecorder()

	Add(w, req)

	if w.Body.String() != "Invalid input, please pass in two integers: a and b" {
		t.Fail()
	}
}

func TestAddReturnsErrorsWhenBIsMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/add?a=1", nil)
	w := httptest.NewRecorder()

	Add(w, req)

	if w.Body.String() != "Invalid input, please pass in two integers: a and b" {
		t.Fail()
	}
}

func TestAddReturnsErrorsWhenAAndBAreMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/add", nil)
	w := httptest.NewRecorder()

	Add(w, req)

	if w.Body.String() != "Invalid input, please pass in two integers: a and b" {
		t.Fail()
	}
}

func TestAddReturnsSumOfAAndB(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/add?a=1&b=2", nil)
	w := httptest.NewRecorder()

	Add(w, req)

	if w.Body.String() != "3" {
		t.Fail()
	}
}
