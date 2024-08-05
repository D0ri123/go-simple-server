package repository

import "go-server/types"

type UserRepository struct {
	userMap []*types.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Get() []*types.User {
	return nil
}

func (u *UserRepository) Update(beforeUser, updateUser *types.User) error {
	return nil
}

func (u *UserRepository) Delete(newUser *types.User) error {
	return nil
}
