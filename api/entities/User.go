package entities

type User struct {
	ID uint `json:"user_id" gorm:"column:user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// custom tablename
func (e *User) TableName() string {
	return "m_users"
}