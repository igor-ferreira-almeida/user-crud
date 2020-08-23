package usersvc

import "github.com/igor-ferreira-almeida/user-crud/domain/usermd"

type UserServiceMock struct {
	HandleFindUserFn func() (usermd.User, error)
	HandleCreateUserFn func() (usermd.User, error)
	HandleUpdateUserFn func() (usermd.User, error)
}

func (mock UserServiceMock) Find(id int64) (usermd.User, error) {
	return mock.HandleFindUserFn()
}

func (mock UserServiceMock) Add(user usermd.User) (usermd.User, error) {
	return mock.HandleCreateUserFn()
}

func (mock UserServiceMock) Update(id int64, user usermd.User) (usermd.User, error) {
	return mock.HandleUpdateUserFn()
}