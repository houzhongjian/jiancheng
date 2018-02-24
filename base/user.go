package base

//User .
type User struct {
	Model
	Account  string `json:"account"`
	Password string `json:"password"`
	Username string `json:"username"`
}

//TableName user.
func (User) TableName() string {
	return "user"
}
