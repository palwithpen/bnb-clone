package repo

import (
	"palwithpen.github.com/airbnb/src/config"
	"palwithpen.github.com/airbnb/src/entity"
)

func CreateUser(user *entity.User) error {
	_, err := config.DB.Model(user).Insert()
	return err
}

func GetUserById(id string) (*entity.User, error) {
	user := new(entity.User)
	err := config.DB.Model(user).Where("user_id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *entity.User) error {
	_, err := config.DB.Model(user).Where("user_id = ?", user.UserID).Update()
	return err
}

func DeleteUser(id string) error {
	user := &entity.User{UserID: id}
	_, err := config.DB.Model(user).Where("user_id = ?", user.UserID).Delete()
	return err
}
