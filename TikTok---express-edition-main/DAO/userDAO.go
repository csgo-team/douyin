package DAO

import (
	"douyin/Init/DB"
	"douyin/pojo"
)

type userdao struct {
	User pojo.User
}

var UserDAO userdao

func (u *userdao) Getuserbyname(name string) *pojo.User {
	db := DB.GetDB()
	db.Table("user").Where("userAccount = ?", name).First(&u.User)

	return &u.User
}
func (u *userdao) GetUserById(id int64) *pojo.User {
	db := DB.GetDB()
	db.Table("user").Where("id = ?", id).First(&u.User)
	return &u.User
}
func (u *userdao) CreateUser(user pojo.User) *pojo.User {
	db := DB.GetDB()
	db.Table("user").Create(&user)
	db.Table("user").Where("userAccount = ?&&userPassword= ?", user.Name, user.Password).First(&u.User)
	return &u.User
}
