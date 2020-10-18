package IUserDataStorage_test

import (
	userDataStorage "github.com/og/goclub/app/user/data_storage"
	IUserDataStorage "github.com/og/goclub/app/user/data_storage/interface"
	"testing"
)

func TestInterface(t *testing.T) {
	like(userDataStorage.DataStorage{})
}
func like(i IUserDataStorage.Interface) {

}
