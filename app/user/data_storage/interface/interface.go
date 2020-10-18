package IUserDataStorage

import (
	m "github.com/og/goclub/app/model"
)

type ReqCreateUser struct {
	Name string
	Mobile string
	Password string
}
type Interface interface {
	UserByMobile(mobile string) (user m.User, hasUser bool)
	UserByName(name string)  (user m.User, hasUser bool)
	CreateUser(data ReqCreateUser) (user m.User, reject error)
}
