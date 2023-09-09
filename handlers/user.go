package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"sustain/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userSrv.GetAllUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Can't Get All of Users",
		})
	}
	return c.JSON(users)
}

func (h userHandler) TokenCheck(c *fiber.Ctx) error {

	header := c.Get("Authorization")
	tokenString := strings.Replace(header, "Bearer ", "", 1)
	hmacSampleSecret := []byte("TCCT_GROUP_TEST")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.JSON(fiber.Map{
			"users": claims,
		})

	} else {
		tokenError := fmt.Sprintln(err)
		c.JSON(fiber.Map{
			"Error": tokenError,
		})
	}
	return nil

}

func (h userHandler) Registers(c *fiber.Ctx) error {

	request := services.RegisterRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}

	response, err := h.userSrv.Register(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid Signup Credentials",
		})
	}

	return c.JSON(response)

}

func (h userHandler) Login(c *fiber.Ctx) error {
	request := services.LoginRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}

	if (request.Username == "" && request.Email == "") || request.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": "invalid login credentials (Email or Password)",
		})
	}

	response, err := h.userSrv.Login(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Can't Create New Product",
		})
	}

	compare := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(request.Password))
	fmt.Println(compare)
	if compare != nil {

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": "invalid login credentials",
		})
	}

	token, exp, err := createJWTToken(*response)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": token,
		"exp":   exp,
		// "user":  response,
	})
}

func createJWTToken(user services.UsersResponse) (string, int64, error) {
	exp := time.Now().Add(time.Second * 10).Unix()
	hmacSampleSecret := []byte("TCCT_GROUP_TEST")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"username":  user.Username,
		"isAdmin":   user.IsAdmin,
		"createdAt": user.CreatedAt,
		"exp":       exp,
	})

	tk, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", 0, err
	}

	return tk, exp, nil
}
