package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"hero0926-api-test/config"

	"github.com/gorilla/mux"
)

// Recipe ...
// Database schema
type Recipe struct {
	UniqueID   int    `json:"UNIQUEID"`
	Name       string `json:"NAME"`
	PrepTime   string `json:"PREPTIME"`
	Difficulty int    `json:"DIFFICULTY"`
	Vegetarian bool   `json:"VEGETARIAN"`

	Recipes []Recipe
}

// Rate ...
// Another Database schema for recipe ratings
type Rate struct {
	Index    int `json:"INDEX"`
	UniqueID int `json:"UNIQUEID"`
	Rate     int `json:"RATE"`
}

const PageLimit = 5

func GetRecipeEndpoint(req *http.Request) (Recipe, error) {

	params := mux.Vars(req)
	id := params["id"]

	var r Recipe
	err := config.SQL.QueryRow(`SELECT * FROM "RECIPE" WHERE "UNIQUEID"=$1;`, id).Scan(&r.UniqueID, &r.Name, &r.PrepTime, &r.Difficulty, &r.Vegetarian)

	return r, err

}

func GetRecipesEndpoint(req *http.Request) (Recipe, error) {

	var r Recipe

	params := mux.Vars(req)
	page, err := strconv.Atoi(params["p"])

	if page == 1 || page == 0 {
		page = 1
	}
	rows, err := config.SQL.Query(`SELECT * FROM "RECIPE" ORDER BY "UNIQUEID" OFFSET ` + strconv.Itoa((page*PageLimit)-PageLimit) + ` LIMIT ` + strconv.Itoa(PageLimit))

	defer rows.Close()

	if err == nil {
		for rows.Next() {
			err = rows.Scan(&r.UniqueID, &r.Name, &r.PrepTime, &r.Difficulty, &r.Vegetarian)
			if err == nil {
				row :=
					Recipe{UniqueID: r.UniqueID,
						Name:       r.Name,
						PrepTime:   r.PrepTime,
						Difficulty: r.Difficulty,
						Vegetarian: r.Vegetarian}
				r.Recipes = append(r.Recipes, row)
			}
		}

	}

	return r, err

}

func CreateRecipeEndpoint(req *http.Request) (Recipe, error) {

	var r Recipe
	_ = json.NewDecoder(req.Body).Decode(&r)

	name := r.Name
	preptime := r.PrepTime
	difficulty := r.Difficulty
	vegetarian := r.Vegetarian

	err := config.SQL.QueryRow(`
		INSERT INTO public."RECIPE"(
			"NAME", "PREPTIME", "DIFFICULTY", "VEGETARIAN")
			VALUES ($1, $2, $3, $4) RETURNING *`,
		name, preptime, difficulty, vegetarian).Scan(&r.UniqueID, &r.Name, &r.PrepTime, &r.Difficulty, &r.Vegetarian)

	return r, err

}

func DeleteRecipeEndpoint(req *http.Request) (int, error) {

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	res, err := config.SQL.Exec(`DELETE FROM "RECIPE" WHERE "UNIQUEID"=$1;`, id)

	if err != nil {
		return 0, err
	}

	rowsDeleted := res.RowsAffected()

	return int(rowsDeleted), nil
}

func UpdateRecipeEndpoint(req *http.Request) (Recipe, error) {

	var r Recipe
	_ = json.NewDecoder(req.Body).Decode(&r)

	name := r.Name
	preptime := r.PrepTime
	difficulty := r.Difficulty
	vegetarian := r.Vegetarian

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	err := config.SQL.QueryRow(`
		UPDATE public."RECIPE"
		SET "NAME"=$2, "PREPTIME"=$3, "DIFFICULTY"=$4, "VEGETARIAN"=$5
		WHERE "UNIQUEID"=$1;`,
		id, name, preptime, difficulty, vegetarian).Scan(&r.UniqueID, &r.Name, &r.PrepTime, &r.Difficulty, &r.Vegetarian)

	return r, err
}

func SearchRecipeNameEndpoint(req *http.Request) (Recipe, error) {

	var r Recipe

	params := mux.Vars(req)
	name := params["q"]
	page, err := strconv.Atoi(params["p"])

	if page == 1 || page == 0 {
		page = 1
	}

	rows, err := config.SQL.Query((`SELECT * FROM "RECIPE" 
		WHERE "NAME" LIKE '%` + name + `%' ORDER BY "UNIQUEID" OFFSET ` + strconv.Itoa((page*PageLimit)-PageLimit) + ` LIMIT ` + strconv.Itoa(PageLimit)))

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&r.UniqueID, &r.Name, &r.PrepTime, &r.Difficulty, &r.Vegetarian)
		if err == nil {
			row :=
				Recipe{UniqueID: r.UniqueID,
					Name:       r.Name,
					PrepTime:   r.PrepTime,
					Difficulty: r.Difficulty,
					Vegetarian: r.Vegetarian}
			r.Recipes = append(r.Recipes, row)
		}
	}

	return r, err

}

func SearchRecipeTimeEndpoint(req *http.Request) (Recipe, error) {

	var r Recipe

	params := mux.Vars(req)
	time := params["q"]
	page, err := strconv.Atoi(params["p"])

	if page == 1 || page == 0 {
		page = 1
	}

	rows, err := config.SQL.Query((`SELECT * FROM "RECIPE" 
		WHERE "PREPTIME" LIKE '%` + time + `%' ORDER BY "UNIQUEID" OFFSET ` + strconv.Itoa((page*PageLimit)-PageLimit) + ` LIMIT ` + strconv.Itoa(PageLimit)))

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&r.UniqueID, &r.Name, &r.PrepTime, &r.Difficulty, &r.Vegetarian)
		if err == nil {
			row :=
				Recipe{UniqueID: r.UniqueID,
					Name:       r.Name,
					PrepTime:   r.PrepTime,
					Difficulty: r.Difficulty,
					Vegetarian: r.Vegetarian}
			r.Recipes = append(r.Recipes, row)
		}
	}

	return r, err

}

func SearchRecipeDifficultyEndpoint(req *http.Request) (Recipe, error) {

	var r Recipe

	params := mux.Vars(req)
	diff := params["q"]
	page, err := strconv.Atoi(params["p"])

	if page == 1 || page == 0 {
		page = 1
	}

	rows, err := config.SQL.Query((`SELECT * FROM "RECIPE" 
		WHERE "DIFFICULTY" =$1 ORDER BY "UNIQUEID" OFFSET ` + strconv.Itoa((page*PageLimit)-PageLimit) + ` LIMIT ` + strconv.Itoa(PageLimit)), diff)

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&r.UniqueID, &r.Name, &r.PrepTime, &r.Difficulty, &r.Vegetarian)
		if err == nil {
			row :=
				Recipe{UniqueID: r.UniqueID,
					Name:       r.Name,
					PrepTime:   r.PrepTime,
					Difficulty: r.Difficulty,
					Vegetarian: r.Vegetarian}
			r.Recipes = append(r.Recipes, row)
		}
	}

	return r, err

}

func SearchRecipeVegeEndpoint(req *http.Request) (Recipe, error) {

	var r Recipe

	params := mux.Vars(req)
	vege := params["q"]
	page, err := strconv.Atoi(params["p"])

	if page == 1 || page == 0 {
		page = 1
	}

	rows, err := config.SQL.Query((`SELECT * FROM "RECIPE" 
		WHERE "VEGETARIAN" =$1 ORDER BY "UNIQUEID" OFFSET ` + strconv.Itoa((page*PageLimit)-PageLimit) + ` LIMIT ` + strconv.Itoa(PageLimit)), vege)

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&r.UniqueID, &r.Name, &r.PrepTime, &r.Difficulty, &r.Vegetarian)
		if err == nil {
			row :=
				Recipe{UniqueID: r.UniqueID,
					Name:       r.Name,
					PrepTime:   r.PrepTime,
					Difficulty: r.Difficulty,
					Vegetarian: r.Vegetarian}
			r.Recipes = append(r.Recipes, row)
		}
	}

	return r, err

}

func RateRecipeEndpoint(req *http.Request) (Rate, error) {

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	var r Rate
	_ = json.NewDecoder(req.Body).Decode(&r)
	rate := r.Rate

	var rateAlreadyExsist int
	err := config.SQL.QueryRow(`SELECT count(*) FROM "RATE" WHERE "UNIQUEID"=$1`, id).Scan(&rateAlreadyExsist)

	if err == nil {

		if rateAlreadyExsist == 0 {
			err = config.SQL.QueryRow(`
				INSERT INTO public."RATE"(
					"UNIQUEID", "RATE")
					VALUES ($1, $2) RETURNING *`,
				id, rate).Scan(&r.Index, &r.UniqueID, &r.Rate)

		}

	}

	return r, err
}
