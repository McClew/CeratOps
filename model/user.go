package model

type User struct {
	Name         string
	Email        string
	PasswordHash string
	Role         string
}

func GetUser() User {
	hash, _ := HashPassword("admin123")
	return User{
		Name:         "Rhys",
		Email:        "rhys@example.com",
		PasswordHash: hash,
		Role:         "Admin",
	}
}
