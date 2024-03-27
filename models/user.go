package models

import "GinRanking/dao"

type User struct {
	Id   int
	Name string
}

func (User) TableName() string {
	return "user"
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func AddUser(username string) (int, error) {
	user := User{Name: username}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}
