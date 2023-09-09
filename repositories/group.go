package repositories

type Groups struct {
	GroupId   int    `gorm:"primaryKey:group_id"`
	GroupName string `gorm:"group_name"`
	CreatedAt string `gorm:"created_at"`
	UpdateAt  string `gorm:"update_at"`
}

type GroupUserReletion struct {
	Group   Groups `gorm:"foreignKey:group_id;constraint:OnDelete:CASCADE"`
	User    Users  `gorm:"foreignKey:user_id;constraint:OnDelete:CASCADE"`
	GroupId int    `gorm:"group_id"`
	UserId  int    `gorm:"user_id"`
}

type GroupRepository interface {
	CreateGroup(Groups) (*Groups, error)
	AddUserToGroup(GroupUserReletion, int, int) (*GroupUserReletion, error)
	DeleteGroup(int) error
	DeleteUserGroup(GroupUserReletion, int, int) (*GroupUserReletion, error)
	UpdateGroup(Groups, int) (*Groups, error)
}
