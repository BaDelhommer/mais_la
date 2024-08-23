package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Base        string   `json:"base"`
	Temp        string   `json:"temp"`
	Ingredients []string `json:"ingredients"`
}

func (cfg *apiConfig) handlerRecipeCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name        string   `json:"name"`
		Base        string   `json:"base"`
		Temp        string   `json:"temp"`
		Ingredients []string `json:"ingredients"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	recipe, err := cfg.DB.CreateRecipe(params.Name, params.Base, params.Temp, params.Ingredients)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create recipe")
		return
	}

	respondWithJson(w, http.StatusCreated, Recipe{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Base:        recipe.Base,
		Temp:        recipe.Temp,
		Ingredients: recipe.Ingredients,
	})

}
