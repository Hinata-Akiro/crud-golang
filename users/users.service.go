package users

import (
	"crud-app/database"
	"crud-app/utils"
	"errors"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

func checkUserExists(email string) bool {
	var user User
	if err := database.Database.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}

func createNewUser(userInput *User) (*User, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	userExist := checkUserExists(userInput.Email)
	if userExist {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(userInput.PassWord)

	if err != nil {
		return nil, err
	}

	userInput.PassWord = hashedPassword
	userInput.ID = id

	if err := database.Database.Create(&userInput).Error; err != nil {
		return nil, err
	}

	return userInput, nil
}

func loginUser(email string, password string) (string, error) {
    var user User
    if err := database.Database.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return "", errors.New("user not found")
        }
        return "", err
    }

    if err := utils.VerifyPassword(user.PassWord, password); err != nil {
        return "", errors.New("incorrect password") 
    }

    token, err := utils.GenerateJwtToken(user.ID.String(), user.Email)

    if err != nil {
        return "", err
    }

    return token, nil
}

func FindUserByEmail(email string) (*User, error) {
	var user User
    if err := database.Database.Where("email =?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}