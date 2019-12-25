package models

type UserModel struct {
	Id        int      `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Age       int      `json:"age"`
	Hobbies   []string `json:"hobbies"`
}
