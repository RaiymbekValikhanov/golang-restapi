package store

import "github.com/RaiymbekValikhanov/golang-restapi/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}