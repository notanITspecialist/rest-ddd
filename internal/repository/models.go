package repository

type (
	User struct {
		Id        string `json:"id"         db:"id"`
		FirstName string `json:"first_name" db:"first_name"`
		LastName  string `json:"last_name"  db:"last_name"`
		Mobile    string `json:"mobile"      db:"mobile"`
	}
)
