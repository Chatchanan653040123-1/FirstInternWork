package services

type UsersResponse struct {
	ID        int    `gorm:"primaryKey:id" json:"id"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	Email     string `db:"email" json:"email"`
	IsAdmin   string `db:"is_admin" json:"is_admin"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type LoginRequest struct {
	Email    string `gorm:"unique" json:"email"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type RegisterRequest struct {
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Email    string `gorm:"unique" json:"email"`
	Role     string `db:"role" json:"role"`
	Group    string `db:"group" json:"group"`
}

type UserService interface {
	GetAllUsers() ([]UsersResponse, error)
	Register(RegisterRequest) (*UsersResponse, error)
	Login(LoginRequest) (*UsersResponse, error)
}
