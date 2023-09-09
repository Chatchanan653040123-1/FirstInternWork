package services

import (
	"strconv"
	"strings"
	"sustain/errs"
	"sustain/logs"
	"sustain/repositories"
	"time"
	"unicode"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) GetAllUsers() ([]UsersResponse, error) {

	users, err := s.userRepo.GetAllUser()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	userResponses := []UsersResponse{}
	for _, user := range users {
		userResponse := UsersResponse{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

// do not edit boolean operation
func checkPasswordStrength(password string) bool {
	checkPasswordFunc := !(viper.GetBool("passwordStrength.checkPasswordStrength"))
	hasUppercase := viper.GetBool("passwordStrength.hasUppercase")
	hasLowercase := viper.GetBool("passwordStrength.hasLowercase")
	hasNumber := viper.GetBool("passwordStrength.hasNumber")
	hasSpecialCharacter := viper.GetBool("passwordStrength.hasSpecialCharacter")
	length := viper.GetBool("passwordStrength.length")
	if len(password) > 8 {
		length = false
	}
	for _, character := range password {
		if unicode.IsUpper(character) {
			hasUppercase = false
		} else if unicode.IsLower(character) {
			hasLowercase = false
		} else if unicode.IsNumber(character) {
			hasNumber = false
		} else if !unicode.IsLetter(character) && !unicode.IsNumber(character) {
			hasSpecialCharacter = false
		}
	}
	return !(hasUppercase || hasLowercase || hasNumber || hasSpecialCharacter || length || checkPasswordFunc)
}

func (s userService) Register(req RegisterRequest) (*UsersResponse, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if req.Username == "" || req.Password == "" || req.Email == "" {
		logs.Error("Username, Password, and Email cannot be empty")
		return nil, errs.NewValidationError("Username, Password, and Email cannot be empty")
	}
	keyWord := []string{
		"com",
		"net",
		"th",
		"co.th",
		"org",
		"io",
		"ac.th",
		"go.th",
		"or.th",
		"in.th",
		"me",
		"biz",
	}

	validDomain := false
	for _, keyword := range keyWord {
		if strings.Contains(req.Email, keyword) {
			validDomain = true
			break
		}
	}

	if !validDomain {
		logs.Error("Email must be in the format of domain name")
		return nil, errs.NewValidationError("Email must be in the format of domain name")
	}

	if !checkPasswordStrength(req.Password) {
		logs.Error("Password must contain at least 8 characters, 1 uppercase, 1 lowercase, 1 number, and 1 special character")
		return nil, errs.NewValidationError("Password must contain at least 8 characters, 1 uppercase, 1 lowercase, 1 number, and 1 special character")
	}
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	user := repositories.Users{
		Username:  req.Username,
		Password:  string(password),
		Email:     req.Email,
		IsAdmin:   false,
		CreatedAt: time.Now().Format("2006-1-2 15:04:05"),
	}

	newUser, err := s.userRepo.Register(user)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	userResponse := UsersResponse{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}
	return &userResponse, nil
}

func (s userService) Login(req LoginRequest) (*UsersResponse, error) {

	user := repositories.Users{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	loginUser, err := s.userRepo.Login(user, user.Username, user.Email)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	userResponse := UsersResponse{
		ID:        loginUser.ID,
		Username:  loginUser.Username,
		Password:  loginUser.Password,
		Email:     loginUser.Email,
		IsAdmin:   strconv.FormatBool(loginUser.IsAdmin),
		CreatedAt: loginUser.CreatedAt,
	}
	return &userResponse, nil
}
