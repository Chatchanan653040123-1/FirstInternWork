package services

import (
	"fmt"
	"sustain/errs"
	"sustain/logs"
	"sustain/repositories"
	"time"
)

type groupService struct {
	groupRepo repositories.GroupRepository
}

func NewGroupService(groupRepo repositories.GroupRepository) groupService {
	return groupService{groupRepo: groupRepo}
}

func (s groupService) CreateGroup(req GroupRequest) (*GroupsResponse, error) {
	group := repositories.Groups{
		GroupName: req.GroupName,
		CreatedAt: time.Now().Format("2006-1-2 15:04:05"),
	}

	newGroup, err := s.groupRepo.CreateGroup(group)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	groupResponse := GroupsResponse{
		GroupId:   newGroup.GroupId,
		GroupName: newGroup.GroupName,
		CreatedAt: newGroup.CreatedAt,
	}
	return &groupResponse, nil
}

func (s groupService) AddUserToGroup(req GroupUserRequest) (*GroupsUsersResponse, error) {
	group := repositories.GroupUserReletion{

		UserId:  req.UserId,
		GroupId: req.GroupId,
	}

	newGroup, err := s.groupRepo.AddUserToGroup(group, group.UserId, group.GroupId)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	groupResponse := GroupsUsersResponse{
		GroupId: newGroup.GroupId,
		UserId:  newGroup.UserId,
	}
	fmt.Println("groupResponse", groupResponse.Username)
	return &groupResponse, nil
}

func (s groupService) DeleteGroup(groupID int) error {
	err := s.groupRepo.DeleteGroup(groupID)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}
	return nil
}

func (s groupService) DeleteUserGroup(req GroupUserRequest) (*GroupsUsersResponse, error) {
	group := repositories.GroupUserReletion{
		GroupId: req.GroupId,
		UserId:  req.UserId,
	}

	deleteGroup, err := s.groupRepo.DeleteUserGroup(group, group.UserId, group.GroupId)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	groupResponse := GroupsUsersResponse{
		GroupId: deleteGroup.GroupId,
		UserId:  deleteGroup.UserId,
	}
	return &groupResponse, nil
}

func (s groupService) UpdateGroup(req UpdateGroupRequest) (*GroupsResponse, error) {

	group := repositories.Groups{
		GroupId:   req.GroupId,
		GroupName: req.GroupName,
		UpdateAt:  time.Now().Format("2006-1-2 15:04:05"),
	}

	updateGroup, err := s.groupRepo.UpdateGroup(group, req.GroupId)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	groupResponse := GroupsResponse{
		GroupId:   updateGroup.GroupId,
		GroupName: updateGroup.GroupName,
		CreatedAt: updateGroup.CreatedAt,
		UpdateAt:  updateGroup.UpdateAt,
	}
	return &groupResponse, nil
}
