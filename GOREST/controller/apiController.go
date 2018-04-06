package controller

import (
	"net/http"

	c "hero0926-api-test/config"
	"hero0926-api-test/model"
	"hero0926-api-test/util"
)

// Pagination ...
type Pagination struct {
	StartListNum uint64
	MaxListCount uint64
	TotalCnt     int `db:"TotalCnt"`
}

// GetAllRecipes ...
// return all recipes
func GetAllRecipes(w http.ResponseWriter, req *http.Request) {

	r, err := model.GetRecipesEndpoint(req)

	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r.Recipes)

}

// CreateRecipe ...
// Create Recipe [auth needed]
func CreateRecipe(w http.ResponseWriter, req *http.Request) {

	r, err := model.CreateRecipeEndpoint(req)

	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r)

}

// GetOneRecipe
// Get specific one recipe
func GetOneRecipe(w http.ResponseWriter, req *http.Request) {

	r, err := model.GetRecipeEndpoint(req)

	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r)

}

// UpdateRecipe ...
// Update one recipe [auth needed]
func UpdateRecipe(w http.ResponseWriter, req *http.Request) {

	r, err := model.UpdateRecipeEndpoint(req)

	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r)

}

// DeleteRecipe ...
// Delete one recipe [auth needed]
func DeleteRecipe(w http.ResponseWriter, req *http.Request) {

	r, err := model.DeleteRecipeEndpoint(req)

	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r)
}

// SearchRecipeName ...
// Search recipe by name
func SearchRecipeName(w http.ResponseWriter, req *http.Request) {
	r, err := model.SearchRecipeNameEndpoint(req)

	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r.Recipes)

}

// SearchRecipeTime ...
// Search recipe by time
func SearchRecipeTime(w http.ResponseWriter, req *http.Request) {
	r, err := model.SearchRecipeTimeEndpoint(req)

	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r.Recipes)

}

// SearchRecipeDifficulty ...
// Search recipe by difficulty
func SearchRecipeDifficulty(w http.ResponseWriter, req *http.Request) {
	r, err := model.SearchRecipeDifficultyEndpoint(req)

	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r.Recipes)

}

// SearchRecipeVege ...
// Search Recipe by its vegetarian or not
func SearchRecipeVege(w http.ResponseWriter, req *http.Request) {
	r, err := model.SearchRecipeVegeEndpoint(req)
	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r.Recipes)

}

// RateOneRecipe ...
// RateOneRecipe rating recipe and can be rated many times from 1-5 and a rating is never overwritten.
func RateOneRecipe(w http.ResponseWriter, req *http.Request) {

	r, err := model.RateRecipeEndpoint(req)
	if err != nil {

		util.WriteJSON(w, http.StatusInternalServerError, nil)
		c.Panic(err)
	}

	util.WriteJSON(w, http.StatusOK, r)

}
