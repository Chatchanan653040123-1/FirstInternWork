package repositories

import (
	"log"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRespositoryDB(db *gorm.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAllUser() ([]Users, error) {
	users := []Users{}
	err := r.db.Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}

	log.Printf("Get All Users Successfully")
	return users, nil
}

func (r userRepositoryDB) Register(user Users) (*Users, error) {

	err := r.db.Create(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	log.Printf("Created Successfully with ID %d", user.ID)
	return &user, nil

}

func (r userRepositoryDB) Login(user Users, username string, email string) (*Users, error) {
	if username != "" {
		err := r.db.First(&user, "username = ?", username)
		if err.Error != nil {
			return nil, err.Error
		}
	}
	if email != "" {
		err := r.db.First(&user, "email = ?", email)
		if err.Error != nil {
			return nil, err.Error
		}
	}

	log.Println("LOGIN : ", user.Email)
	return &user, nil
}
