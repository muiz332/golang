package repositoryimplement

import (
	"context"
	"database/sql"
	"errors"
	"repository_pattern/entity"
	respositoryinterface "repository_pattern/respositoryInterface"
)

type usersRepositoryImp struct {
	DB *sql.DB
}


func (repository *usersRepositoryImp) FindAll(ctx context.Context) ([]entity.User, error) {
	sintaxSql := "SELECT * FROM users"
	rows, err := repository.DB.QueryContext(ctx, sintaxSql)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	users := []entity.User{}
	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Gender, &user.Password, &user.CreateAt)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}


func NewUsersRepository(db *sql.DB) respositoryinterface.UsersRepositoryInterface{
	return &usersRepositoryImp{
		DB : db,
	}
}
