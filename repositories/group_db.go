package repositories

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type groupRepositoryDB struct {
	db *gorm.DB
}

func NewGroupRespositoryDB(db *gorm.DB) groupRepositoryDB {
	return groupRepositoryDB{db: db}
}

func (r groupRepositoryDB) CreateGroup(group Groups) (*Groups, error) {

	err := r.db.Create(&group)
	if err.Error != nil {
		return nil, err.Error
	}

	log.Println("Created Successfully with ", group)
	return &group, nil

}

func (r groupRepositoryDB) AddUserToGroup(group GroupUserReletion, user_id int, group_id int) (*GroupUserReletion, error) {
	err := r.db.Where("group_id = ? AND user_id = ?", group_id, user_id).FirstOrCreate(&group).Find(&group)
	if err.Error != nil {
		return nil, err.Error
	}

	log.Println("Add User To Group Successfully with ", group)
	return &group, nil

}

func (r groupRepositoryDB) DeleteGroup(groupID int) error {
	err := r.db.Delete(&GroupUserReletion{}, "group_id = ?", groupID)
	if err.Error != nil {
		return err.Error
	}

	result := r.db.Delete(&Groups{}, "group_id = ?", groupID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("group not found")
		}
		return result.Error
	}

	return nil
}

func (r groupRepositoryDB) DeleteUserGroup(group GroupUserReletion, user_id int, group_id int) (*GroupUserReletion, error) {
	err := r.db.Where("group_id = ? AND user_id = ?", group_id, user_id).Delete(&group).Scan(&group)
	if err.Error != nil {
		return nil, err.Error
	}

	log.Println("Add User To Group Successfully with ", group)
	return &group, nil
}

func (r groupRepositoryDB) UpdateGroup(group Groups, group_id int) (*Groups, error) {
	err := r.db.Model(&group).Where("group_id = ?", group_id).Updates(Groups{GroupName: group.GroupName, UpdateAt: group.UpdateAt})
	if err.Error != nil {
		log.Println("ERROR REPO: ", err.Error)
		return nil, err.Error
	}

	log.Printf("Updated Successfully with ID %d", group.GroupId)
	return &group, nil
}
