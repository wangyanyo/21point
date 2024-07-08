package dao

import "github.com/wangyanyo/21point/Server/models"

func InitTable() {
	new(models.User).CreateTable()
}
