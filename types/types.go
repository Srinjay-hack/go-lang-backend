package types

type UserStore interface{
	getUserEmailId(email string)(*User,error)
	getUserById(id int)(*User,error)
	createUser(User)(*User,error)

}

type User struct{
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
}
type RegistryUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
