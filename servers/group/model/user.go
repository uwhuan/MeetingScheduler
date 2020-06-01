package model

// User struct is to store information of user
// This struct will not be used to signin
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
