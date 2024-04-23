package users

import(
	"crud-app/database"
	"errors"
	"crud-app/utils"

	"gorm.io/gorm"
)


func checkUserExists( email string) bool {
	var user User
    if err := database.Database.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false
        }
    }
    return true
}


func createNewUser(userInput *User) (*User, error) {

   userExist := checkUserExists(userInput.Email)
    if userExist {
        return nil, errors.New("user already exists")
    }


	hashedPassword , err := utils.HashPassword(userInput.PassWord)

	if err != nil {
        return nil, err
    }

	userInput.PassWord = hashedPassword


    if err := database.Database.Create(&userInput).Error; err != nil {
        return nil, err
    }

    return userInput, nil
}