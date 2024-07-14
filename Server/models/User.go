package models

import (
	"log"
	"time"

	"github.com/wangyanyo/21point/common/db"
)

type User struct {
	Id         int       `grom:"primary_key;auto_increment" json:"id"`
	UserName   string    `grom:"column:user_name;unique" json:"user_name"`
	Password   string    `grom:"column:password" json:"password"`
	Score      string    `grom:"column:score" json:"score"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"` // 更新时间
}

func (user *User) TableName() string {
	return "tbl_user"
}

func (user *User) CreateTable() {
	if !db.MysqlDB.HasTable(user.TableName()) {
		log.Println("CreateTable User")
		db.MysqlDB.CreateTable(&User{})
	}
}
