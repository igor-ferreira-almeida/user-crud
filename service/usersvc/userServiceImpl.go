package usersvc

import "github.com/igor-ferreira-almeida/user-crud/domain/usermd"

type serviceImpl struct {

}

func newServiceImpl() Service {
	return serviceImpl{}
}

func (service serviceImpl) Add(user usermd.User) usermd.User {
	return user
}

func (service serviceImpl) Update(id int64, user usermd.User) usermd.User {
	user.ID = id
	return user
}