package usersvc

import "github.com/igor-ferreira-almeida/user-crud/domain/usermd"

var service Service

type Service interface {
	Add(user usermd.User) usermd.User
	Update(id int64, user usermd.User) usermd.User
}

func init() {
	service = newServiceImpl()
}

func Inject() Service {
	return service
}