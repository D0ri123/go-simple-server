package service

import (
	"go-server/repository"
	"sync"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

// 네트워크와 레포지터리의 브릿지 역할.
type Service struct {
	repository *repository.Repository

	User *User
}

func NewService(repo *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: repo,
		}

		serviceInstance.User = newUserService(repo.User)
	})

	return serviceInstance
}
