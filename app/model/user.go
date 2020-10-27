package m

import (
	"database/sql"
	f "github.com/og/gofree"
	gcrypto "github.com/og/x/crypto"
	"time"
)

type IDUser string
type User struct {
	ID IDUser `db:"id"`
	Name string `db:"name"`
	Mobile string `db:"mobile"`
	Password string `db:"password"`
	PasswordSalt string `db:"password_salt"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
func (User) Column()( c struct{
	ID f.Column
	Name f.Column
	Mobile f.Column
	Password f.Column
	PasswordSalt f.Column
	CreatedAt f.Column
	UpdatedAt f.Column
	DeletedAt f.Column
}) {
	c.ID = "id"
	c.Name = "name"
	c.Mobile = "mobile"
	c.PasswordSalt = "mobile"
	c.PasswordSalt = "password_salt"
	c.CreatedAt = "created_at"
	c.UpdatedAt = "updated_at"
	c.DeletedAt = "deleted_at"
}
type UserList []User
func (list UserList) IDList() (idList struct {
	UserID []IDUser
}) {
	for _, item := range list {
		idList.UserID = append(idList.UserID, item.ID)
	}
	return
}
func (list UserList) FindByID(userID IDUser) (user User, hasUser bool) {
	for _, item := range list {
		if item.ID == userID {
			return item, true
		}
	}
	return User{}, false
}
func (model User) VerifyPassword(password string) {
	pass := model.Password == gcrypto.SaltSHA512(gcrypto.SaltEncode{
		Password: password,
		Salt: "",
	})
	if pass {
		
	}

}
func (User) TableName() string {
	return "user"
}
func (model *User) BeforeCreate() {
	if model.ID == "" {
		model.ID = IDUser(f.UUID())
	}
}
