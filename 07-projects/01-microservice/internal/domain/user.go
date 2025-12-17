package domain

// User 代表用户领域模型
type User struct {
	ID    string
	Name  string
	Email string
}

// NewUser 创建一个新的用户
func NewUser(id, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
