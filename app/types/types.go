package types

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID        int `json:"id"`
	FirstName int `json:"firstName"`
	LastName  int `json:"lastName"`
	Email     int `json:"email"`
	Password  int `json:"-"`
	CreatedAt int `json:"createdAt"`
}
