package service

import (
	"go-server/repository"
	"go-server/types"
)

type User struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}

func (u *User) Update(name string, newAge int64) error {
	return u.userRepository.Update(name, newAge)
}

func (u *User) Delete(user *types.User) error {
	//user라는 스트럭처를 어디에서 사용할지 모른다. repository를 통할 때는 필요한 필드만 넘겨주는게 좋다.
	return u.userRepository.Delete(user.Name)
}
