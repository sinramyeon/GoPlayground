package model

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"testing"
)

var req *http.Request

// basicAuth ...
// basic auth for protected endpoint
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func TestGetRecipesEndpoint(t *testing.T) {

	req, err := http.NewRequest("GET", "/recipes", nil)
	_, err = GetRecipesEndpoint(req)
	if err != nil {
		t.Errorf("GetRecipesEndpoint() error = %v", err)
		return
	}

}

func TestGetRecipeEndpoint(t *testing.T) {

	req, err := http.NewRequest("GET", "/recipe/1", nil)
	_, err = GetRecipesEndpoint(req)
	if err != nil {
		t.Errorf("GetRecipeEndpoint() error = %v", err)
		return
	}

}

func TestCreateRecipeEndpoint(t *testing.T) {
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

	_, err = CreateRecipeEndpoint(req)
	if err != nil {
		t.Errorf("CreateRecipeEndpoint() error = %v", err)
		return
	}
}

func TestDeleteRecipeEndpoint(t *testing.T) {
	req, err := http.NewRequest("DELETE", "recipes/1", nil)
	req.Header.Add("Authorization", "Basic"+basicAuth("hellofresh", "hellofresh"))

	_, err = DeleteRecipeEndpoint(req)

	if err != nil {
		t.Errorf("DeleteRecipeEndpoint() error = %v", err)
		return
	}
}

func TestUpdateRecipeEndpoint(t *testing.T) {

	var r Recipe
	var jsonStr = []byte(`
		{
			"NAME" : "Updated",
			"PREPTIME" : "30 min",
			"DIFFICULTY" : 3,
			"VEGETARIAN" : true
			
		}
		`)
	req, err := http.NewRequest("PUT", "recipes/1", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", "Basic"+basicAuth("hellofresh", "hellofresh"))

	res, err := UpdateRecipeEndpoint(req)

	if err != nil {
		// If rows are unexsists
		if res.Name == r.Name {

			t.Errorf("UpdateRecipeEndpoint() error = %v", err)
			return

		}

	}
}

func TestSearchRecipeNameEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/recipes/name/search?q=New Recipe", nil)
	_, err = SearchRecipeNameEndpoint(req)
	if err != nil {
		t.Errorf("SearchRecipeNameEndpoint() error = %v", err)
		return
	}
}

func TestSearchRecipeTimeEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/recipes/time/search?q=1 min", nil)
	_, err = SearchRecipeTimeEndpoint(req)
	if err != nil {
		t.Errorf("SearchRecipeTimeEndpoint() error = %v", err)
		return
	}
}

func TestSearchRecipeDifficultyEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/recipes/difficulty/search?q=1", nil)
	_, err = SearchRecipeDifficultyEndpoint(req)
	if err != nil {
		t.Errorf("SearchRecipeDifficultyEndpoint() error = %v", err)
		return
	}
}

func TestSearchRecipeVegeEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/recipes/vege/search?q=true", nil)
	_, err = SearchRecipeVegeEndpoint(req)
	if err != nil {
		t.Errorf("SearchRecipeVegeEndpoint() error = %v", err)
		return
	}
}

func TestRateRecipeEndpoint(t *testing.T) {
	var jsonStr = []byte(`
	{
		"RATE" : 1
		
	}
	`)
	req, err := http.NewRequest("POST", "/recipes/1/rating", bytes.NewBuffer(jsonStr))
	_, err = RateRecipeEndpoint(req)
	if err != nil {
		t.Errorf("RateRecipeEndpoint() error = %v", err)
		return
	}
}
