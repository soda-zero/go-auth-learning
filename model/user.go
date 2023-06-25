package model

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) error {
	query := `INSERT INTO users(name, email, password) VALUES($1, $2, $3)`
	_, err := db.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func GetUser(id string) (User, error) {
	var user User
	query := `SELECT id, name, email FROM users WHERE id = $1;`
	rows, err := db.Query(query, id)
	if err != nil {
		return User{}, err
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func CheckEmail(email string, user *User) bool {
	query := "SELECT id, name, email, password FROM users WHERE email = $1 limit 1"

	rows, err := db.Query(query, email)
	if err != nil {
		return false
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return false
		}
	}
	return true
}
