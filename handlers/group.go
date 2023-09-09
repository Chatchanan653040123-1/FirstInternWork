package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"sustain/services"

	"github.com/gofiber/fiber/v2"
)

type groupHandler struct {
	groupSrv services.GroupService
}

func NewGroupHandler(groupSrv services.GroupService) groupHandler {
	return groupHandler{groupSrv: groupSrv}
}

func (h groupHandler) CreateGroup(c *fiber.Ctx) error {

	request := services.GroupRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}

	response, err := h.groupSrv.CreateGroup(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid Credentials",
		})
	}

	return c.JSON(response)

}

func (h groupHandler) AddUserToGroup(c *fiber.Ctx) error {
	request := services.GroupUserRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}

	if request.GroupId == 0 || request.UserId == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": "invalid credentials (Group ID or User ID)",
		})
	}

	response, err := h.groupSrv.AddUserToGroup(request)
	if err != nil {
		fmt.Println("request", request)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Can't Add User To Group",
		})
	}
	return c.JSON(response)
}

func (h groupHandler) DeleteGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("group_id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to parse group ID",
		})
	}

	err = h.groupSrv.DeleteGroup(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to delete group",
		})
	}

	return c.SendStatus(http.StatusOK)
}

func (h groupHandler) DeleteUserGroup(c *fiber.Ctx) error {
	request := services.GroupUserRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}

	if request.GroupId == 0 || request.UserId == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": "invalid credentials (Group ID or User ID)",
		})
	}

	response, err := h.groupSrv.DeleteUserGroup(request)
	if err != nil {
		fmt.Println("request", request)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Can't Delete User In Group",
		})
	}
	return c.JSON(response)
}

func (h groupHandler) UpdateGroup(c *fiber.Ctx) error {
	group := services.UpdateGroupRequest{}

	if err := c.BodyParser(&group); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}

	response, err := h.groupSrv.UpdateGroup(group)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": "Can't Update Group",
		})
	}

	return c.JSON(response)
}
