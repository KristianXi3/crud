package service

import (
	"fmt"

	"github.com/KristianXi3/crud/entity1"
)

type UserServiceIface interface {
	Register(user *entity1.User) *entity1.User
}

type UserSvc struct {
	ListUser map[string]entity1.User
}

func NewUserService() UserServiceIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *entity1.User) *entity1.User {
	if _, ok := u.ListUser[user.Username]; ok {
		fmt.Println("Gagal, silahkan gunakan username lainnya")
	} else {
		fmt.Println("Berhasil")
	}

	fmt.Println(user)
	return user
}
