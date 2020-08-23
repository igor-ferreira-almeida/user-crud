package usersvc

import "github.com/igor-ferreira-almeida/user-crud/domain/usermd"

var service Service

type Service interface {
	Find(id int64) (usermd.User, error)
	Add(user usermd.User) (usermd.User, error)
	Update(id int64, user usermd.User) (usermd.User, error)
}

func init() {
	service = newServiceImpl()
}

func Inject() Service {
	return service
}