package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerRecipeGet(w http.ResponseWriter, r *http.Request) {
	recipeIDString := r.PathValue("recipeID")
	recipeID, err := strconv.Atoi(recipeIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	dbRecipe, err := cfg.DB.GetRecipe(recipeID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get recipe")
		return
	}

	respondWithJson(w, http.StatusOK, Recipe{
		ID:          dbRecipe.ID,
		Name:        dbRecipe.Name,
		Base:        dbRecipe.Base,
		Temp:        dbRecipe.Temp,
		Ingredients: dbRecipe.Ingredients,
	})
}

func (cfg *apiConfig) handlerGetAllRecipes(w http.ResponseWriter, r *http.Request) {
	dbRecipes, err := cfg.DB.GetAllRecipes()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve recipes")
		return
	}

	recipes := []Recipe{}

	for _, dbRecipe := range dbRecipes {
		recipes = append(recipes, Recipe{
			ID:          dbRecipe.ID,
			Name:        dbRecipe.Name,
			Base:        dbRecipe.Base,
			Temp:        dbRecipe.Temp,
			Ingredients: dbRecipe.Ingredients,
		})
	}

	respondWithJson(w, http.StatusOK, recipes)
}

func (cfg *apiConfig) renderHTML(w http.ResponseWriter, path string) {
	tmpl, err := template.New(path).ParseFiles(path)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create template")
	}

	err = tmpl.Execute(w, tmpl)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't execute template")
	}
}
