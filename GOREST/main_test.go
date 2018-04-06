package main_test

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var Router *mux.Router

// HealthCheckHandler ...
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

// HttpHandler...
func HttpHandler(t *testing.T, req *http.Request, err error) {

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("[ERROR] Wrong status code : got %v expected %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("[ERROR] Wrong body : got %v expected %v",
			rr.Body.String(), expected)
	}
}

// basicAuth ...
// basic auth for protected endpoint
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// executeRequest ...
// execute Request with mux handler
func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	Router = mux.NewRouter()
	rr := httptest.NewRecorder()
	Router.ServeHTTP(rr, req)
	return rr
}

func TestCreateRecipe(t *testing.T) {

	var jsonStr = []byte(`
		{
			"NAME" : "New Recipe",
			"PREPTIME" : "30 min",
			"DIFFICULTY" : 3,
			"VEGETARIAN" : true
			
		}
		`)

	req, err := http.NewRequest("POST", "recipes", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", "Basic"+basicAuth("hellofresh", "hellofresh"))

	HttpHandler(t, req, err)
}

func TestGetOneRecipe(t *testing.T) {

	req, err := http.NewRequest("GET", "recipes/1", nil)

	HttpHandler(t, req, err)

}

func TestGetAllRecipes(t *testing.T) {

	req, err := http.NewRequest("GET", "/recipes", nil)

	HttpHandler(t, req, err)

}

func TestUpdateRecipe(t *testing.T) {
	var jsonStr = []byte(`
		{
			"NAME" : "New Recipe Updated",
			"PREPTIME" : "30 min",
			"DIFFICULTY" : 3,
			"VEGETARIAN" : true
			
		}
		`)
	req, err := http.NewRequest("PUT", "recipes/1", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", "Basic"+basicAuth("hellofresh", "hellofresh"))

	HttpHandler(t, req, err)

}

func TestSearchRecipeID(t *testing.T) {
	req, err := http.NewRequest("GET", "recipies/id/search?q=1", nil)

	HttpHandler(t, req, err)

}

func TestSearchRecipeName(t *testing.T) {
	req, err := http.NewRequest("GET", "recipes/name/search?q=New Recipe", nil)
	req.Header.Add("Authorization", "Basic"+basicAuth("hellofresh", "hellofresh"))

	HttpHandler(t, req, err)

}

func TestSearchRecipeTime(t *testing.T) {
	req, err := http.NewRequest("GET", "recipes/time/search?q=30 min", nil)
	HttpHandler(t, req, err)

}

func TestSearchRecipeDifficulty(t *testing.T) {
	req, err := http.NewRequest("GET", "recipes/difficulty/search?q=3", nil)
	HttpHandler(t, req, err)

}

func TestSearchRecipeVege(t *testing.T) {
	req, err := http.NewRequest("GET", "recipes/vegetarian/search?q=true", nil)
	HttpHandler(t, req, err)

}

func TestRateOneRecipe(t *testing.T) {

	var jsonStr = []byte(`
		{
			"RATE" : 1
			
		}
		`)

	req, err := http.NewRequest("POST", "recipes/1/rating", bytes.NewBuffer(jsonStr))

	HttpHandler(t, req, err)
}

func TestDeleteRecipe(t *testing.T) {

	req, err := http.NewRequest("DELETE", "recipies/1", nil)
	HttpHandler(t, req, err)

}
