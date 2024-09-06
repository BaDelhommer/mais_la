package database

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	APIKey    string `json:"apikey"`
	IsManager bool   `json:"ismanager"`
}

func (db *DB) CreateUser(username, apiKey string, isManager bool) (User, error) {
	dbStructure, err := db.LoadDB()
	if err != nil {
		return User{}, err
	}

	id := len(dbStructure.Recipe) + 1
	user := User{
		ID:        id,
		Username:  username,
		APIKey:    apiKey,
		IsManager: isManager,
	}

	dbStructure.User[id] = user
	err = db.WriteDB(dbStructure)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (db *DB) GetUsersByAPIKey(apiKey string) (User, error) {
	dbStructure, err := db.LoadDB()
	if err != nil {
		return User{}, err
	}

	for _, user := range dbStructure.User {
		if user.APIKey == apiKey {
			return user, nil
		}
	}

	return User{}, ErrNotExist
}

func (db *DB) GetAllUsers() (map[int]User, error) {
	dbStructure, err := db.LoadDB()
	if err != nil {
		return map[int]User{}, err
	}

	return dbStructure.User, nil

}
