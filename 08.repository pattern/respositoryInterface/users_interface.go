package respositoryinterface

import (
	"context"
	"repository_pattern/entity"
)

type UsersRepositoryInterface interface {
	FindAll(ctx context.Context) ([]entity.User, error)
}
