package main

import (
	"hero0926-api-test/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var Router *mux.Router

func main() {

	Router = mux.NewRouter()

	Router.HandleFunc("/recipes", controller.GetAllRecipes).Methods("GET")
	Router.HandleFunc("/recipes", controller.GetAllRecipes).Methods("GET").Queries("p", "{p}")
	Router.HandleFunc("/recipes", controller.IsAuthenticated(controller.CreateRecipe)).Methods("POST")
	Router.HandleFunc("/recipes/{id}", controller.GetOneRecipe).Methods("GET")
	Router.HandleFunc("/recipes/{id}", controller.IsAuthenticated(controller.UpdateRecipe)).Methods("PUT")
	Router.HandleFunc("/recipes/{id}", controller.IsAuthenticated(controller.DeleteRecipe)).Methods("DELETE")
	Router.HandleFunc("/recipes/{id}/rating", controller.RateOneRecipe).Methods("POST")

	Router.HandleFunc("/recipes/name/search", controller.SearchRecipeName).Methods("GET").Queries("q", "{q}")
	Router.HandleFunc("/recipes/time/search", controller.SearchRecipeTime).Methods("GET").Queries("q", "{q}")
	Router.HandleFunc("/recipes/difficulty/search", controller.SearchRecipeDifficulty).Methods("GET").Queries("q", "{q}")
	Router.HandleFunc("/recipes/vegetarian/search", controller.SearchRecipeVege).Methods("GET").Queries("q", "{q}")

	Router.HandleFunc("/recipes/name/search", controller.SearchRecipeName).Methods("GET").Queries("q", "{q}", "p", "{p}")
	Router.HandleFunc("/recipes/time/search", controller.SearchRecipeTime).Methods("GET").Queries("q", "{q}", "p", "{p}")
	Router.HandleFunc("/recipes/difficulty/search", controller.SearchRecipeDifficulty).Methods("GET").Queries("q", "{q}", "p", "{p}")
	Router.HandleFunc("/recipes/vegetarian/search", controller.SearchRecipeVege).Methods("GET").Queries("q", "{q}", "p", "{p}")

	log.Fatal(http.ListenAndServe(":8090", Router))

}
