package usersvc

import "github.com/igor-ferreira-almeida/user-crud/domain/usermd"

type serviceImpl struct {

}

func newServiceImpl() Service {
	return serviceImpl{}
}

func (service serviceImpl) Find(id int64) (usermd.User, error) {
	user := usermd.User{ID: 1, Name: "Tales", Age: 23, Gender: "male"}
	return user, nil
}

func (service serviceImpl) Add(user usermd.User) (usermd.User, error) {
	return user, nil
}

func (service serviceImpl) Update(id int64, user usermd.User) (usermd.User, error) {
	foundUser, err := service.Find(id)

	if err != nil {
		return usermd.User{}, err
	}
	user.ID = id
	foundUser = user

	return foundUser, nil
}