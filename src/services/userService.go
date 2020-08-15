package services

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserList() []User {
	return []User{
		{
			Username: "user 1",
			Password: "123456",
		},
		{
			Username: "user 2",
			Password: "123456",
		},
	}
}
