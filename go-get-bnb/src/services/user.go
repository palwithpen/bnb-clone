package services

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"palwithpen.github.com/airbnb/src/entity"
	"palwithpen.github.com/airbnb/src/repo"
	"palwithpen.github.com/airbnb/src/ro"
)

func SaveUser(user ro.User) (string, error) {
	var userEntity entity.User
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	userEntity.UserID = uuid.NewString()
	userEntity.Username = user.Username
	userEntity.Email = user.Email
	userEntity.PasswordHash = string(hashed)
	userEntity.CreatedAt = time.Now()
	userEntity.UpdatedAt = time.Now()

	if err := repo.CreateUser(&userEntity); err != nil {
		return "", err
	}

	return "", nil
}
