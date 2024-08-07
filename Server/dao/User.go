package dao

import (
	"errors"

	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/db"
)

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

func (d *UserDao) WhereId(id int64) *UserDao {
	clone := d.Clone()
	clone.id = id
	return clone
}

func (d *UserDao) WhereName(name string) *UserDao {
	clone := d.Clone()
	clone.name = name
	return clone
}

func (d *UserDao) Get() (models.User, error) {
	var (
		data   models.User
		err    error
		ok     = false
		dbConn = db.MysqlDB
	)
	dbConn = dbConn.Table(data.TableName())

	if d.id != 0 {
		ok = true
		dbConn = dbConn.Where("id=?", d.id)
	}

	if len(d.name) > 0 {
		ok = true
		dbConn = dbConn.Where("user_name=?", d.name)
	}

	if ok {
		err = dbConn.First(&data).Error
	} else {
		err = errors.New("参数不全")
	}

	d.Clear()
	return data, err
}

func (d *UserDao) Create(user models.User) error {
	dbConn := db.MysqlDB
	return dbConn.Table(user.TableName()).Create(user).Error
}

func (d *UserDao) IsHave(username string) (bool, error) {
	d.WhereName(username)
	user, err := d.Get()
	d.Clear()
	if err != nil || user.Id == 0 {
		return false, err
	}
	return true, nil
}

func (d *UserDao) Delete() error {
	var (
		data   models.User
		err    error
		ok     = false
		dbConn = db.MysqlDB
	)
	dbConn = dbConn.Table(data.TableName())

	if d.id > 0 {
		ok = true
		dbConn = dbConn.Where("id=?", d.id)
	}

	if len(d.name) != 0 {
		ok = true
		dbConn = dbConn.Where("user_name=?", d.name)
	}

	if ok {
		err = dbConn.Delete(&data).Error
	} else {
		return errors.New("参数不全")
	}

	d.Clear()
	return err
}

func (d *UserDao) Clear() {
	d.id = 0
	d.name = ""
}

func (d *UserDao) Clone() *UserDao {
	return &UserDao{
		id:   d.id,
		name: d.name,
	}
}
