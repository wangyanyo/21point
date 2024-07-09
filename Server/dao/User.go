package dao

import "github.com/wangyanyo/21point/Server/models"

type UserDao struct {
	id   int64
	name string
}

type UserDaoInterface interface {
	WhereId(id int64) *UserDao
	WhereName(name string) *UserDao
	Create(user models.User) error
	IsHave(username string) (bool, error)
	Get() (models.User, error)
}

func User() UserDaoInterface {
	return &UserDao{}
}
