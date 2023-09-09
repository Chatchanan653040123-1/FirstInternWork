package repositories

type Users struct {
	ID        int    `gorm:"primaryKey:id" json:"id"`
	Username  string `db:"username" json:"username"`
	Password  string `json:"-"`
	Email     string `gorm:"unique" json:"email"`
	IsAdmin   bool   `db:"is_admin" json:"is_admin"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type UserRepository interface {
	GetAllUser() ([]Users, error)
	Register(Users) (*Users, error)
	Login(Users, string, string) (*Users, error)
}
