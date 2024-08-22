package database

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Base        string   `json:"base"`
	Temp        string   `json:"temp"`
	Ingredients []string `json:"ingredients"`
}

func (db *DB) CreateRecipe(name, base, temp string, ingredients []string) (Recipe, error) {
	dbStructure, err := db.LoadDB()
	if err != nil {
		return Recipe{}, err
	}

	id := len(dbStructure.Recipe) + 1
	recipe := Recipe{
		ID:          id,
		Name:        name,
		Base:        base,
		Temp:        temp,
		Ingredients: ingredients,
	}

	dbStructure.Recipe[id] = recipe
	db.WriteDB(dbStructure)

	return recipe, nil
}

func (db *DB) GetRecipe(id int) (Recipe, error) {
	dbStructure, err := db.LoadDB()
	if err != nil {
		return Recipe{}, err
	}

	recipe, ok := dbStructure.Recipe[id]
	if !ok {
		return Recipe{}, err
	}

	return recipe, nil
}

func (db *DB) GetAllRecipes() (map[int]Recipe, error) {
	dbStructure, err := db.LoadDB()
	if err != nil {
		return map[int]Recipe{}, err
	}

	return dbStructure.Recipe, nil

}
