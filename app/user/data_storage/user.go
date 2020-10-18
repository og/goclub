package userDataStorage

import (
	f "github.com/og/gofree"
	m "github.com/og/goclub/app/model"
	IUserDataStorage "github.com/og/goclub/app/user/data_storage/interface"
	gcrypto "github.com/og/x/crypto"
)

func (dep DataStorage) UserByMobile(mobile string) (user m.User, hasUser bool) {
	userCol := user.Column()
	dep.db.OneQB(&user, &hasUser, f.QB{
		Where: f.And(userCol.Mobile, mobile),
	})
	return
}
func (dep DataStorage) UserByName(name string)  (user m.User, hasUser bool) {
	userCol := user.Column()
	dep.db.OneQB(&user, &hasUser, f.QB{
		Where: f.And(userCol.Name, name),
	})
	return
}
func (dep DataStorage) CreateUser(data IUserDataStorage.ReqCreateUser) (user m.User, reject error) {
	saltEncode := gcrypto.SaltEncode{
		Password: data.Password,
		Salt: f.UUID(),
	}
	user = m.User{
		Name: data.Name,
		Mobile: data.Mobile,
		Password: gcrypto.SaltSHA512(saltEncode),
		PasswordSalt: saltEncode.Salt,
	}
	dep.db.Create(&user)
	return user, nil
}