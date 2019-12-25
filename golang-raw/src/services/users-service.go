package services

import (
	"../models"
	"log"
)

type UsersService struct {
	user  models.UserModel
	users []models.UserModel
}

func (us *UsersService) Init() {
	us.user = models.UserModel{
		Id:        1,
		FirstName: "Hello",
		LastName:  "World",
		Age:       44,
		Hobbies:   []string{"football", "basketball", "guitar", "biking"},
	}
	for i := 0; i < 1000; i++ {
		us.users = append(us.users, models.UserModel{
			Id:        1,
			FirstName: "Hello",
			LastName:  "World",
			Age:       44,
			Hobbies:   []string{"football", "basketball", "guitar", "biking"},
		})
	}
	log.Print("UsersService Init")
}

func (us *UsersService) GetUsers() *[]models.UserModel {
	return &us.users
}

func (us *UsersService) GetUser() *models.UserModel {
	return &us.user
}
