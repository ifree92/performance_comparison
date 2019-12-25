package services

var userServiceInstance UsersService = UsersService{}

func GetUserServiceInstance() *UsersService {
	return &userServiceInstance
}
