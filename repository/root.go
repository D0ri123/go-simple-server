package repository

import (
	"sync"
)

// DB를 연결한다 했을 대 3티어 아키텍처
var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	// DB 연결이 필요할 것임
	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: NewUserRepository(),
		}
	})

	return repositoryInstance
}
