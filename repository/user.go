package repository

import (
	"go-server/types"
	"go-server/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	u.userMap = append(u.userMap, newUser)
	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}

func (u *UserRepository) Update(name string, newAge int64) error {
	isExist := false

	for _, user := range u.userMap {
		if user.Name == name {
			user.Age = newAge
			isExist = true
			continue
		}
	}

	if !isExist {
		// 에러코드 작성 필요
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}

	return nil
}

func (u *UserRepository) Delete(userName string) error {
	isExist := false

	for index, user := range u.userMap {
		if user.Name == userName {
			u.userMap = append(u.userMap[:index], u.userMap[index+1:]...)
			isExist = true
			continue
		}
	}

	if !isExist {
		// 에러코드 작성 필요
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}

	return nil
}
