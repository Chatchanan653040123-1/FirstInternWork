package services

type GroupsResponse struct {
	GroupId   int    `gorm:"primaryKey:group_id" json:"id"`
	GroupName string `gorm:"group_name" json:"group_name"`
	CreatedAt string `gorm:"created_at" json:"created_at"`
	UpdateAt  string `db:"update_at" json:"update_at"`
}

type GroupsUsersResponse struct {
	GroupId   int    `gorm:"group_id" json:"id"`
	GroupName string `gorm:"group_name" json:"group_name"`
	UserId    int    `gorm:"group_id" json:"UserId"`
	Username  string `json:"username"`
}

type GroupRequest struct {
	GroupName string `db:"group_name" json:"group_name"`
	Username  string `db:"username" json:"username"`
}

type GroupUserRequest struct {
	GroupId int `gorm:"group_id" json:"group_id"`
	UserId  int `gorm:"user_id" json:"user_id"`
}

type UpdateGroupRequest struct {
	GroupId   int    `gorm:"group_id" json:"group_id"`
	GroupName string `gorm:"group_name" json:"group_name"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdateAt  string `gorm:"update_at" json:"update_at"`
}

type GroupService interface {
	CreateGroup(GroupRequest) (*GroupsResponse, error)
	AddUserToGroup(GroupUserRequest) (*GroupsUsersResponse, error)
	DeleteGroup(int) error
	DeleteUserGroup(GroupUserRequest) (*GroupsUsersResponse, error)
	UpdateGroup(UpdateGroupRequest) (*GroupsResponse, error)
}
