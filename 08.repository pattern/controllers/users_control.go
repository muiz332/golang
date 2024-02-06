package controllers

import (
	"context"
	"database/sql"
	"errors"
	"repository_pattern/entity"
	repositoryimplement "repository_pattern/repositoryImplement"
)

func GetAllUsers(db *sql.DB) ([]entity.User, error) {
	user := repositoryimplement.NewUsersRepository(db)
	var ctx = context.Background()
	data, err := user.FindAll(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return data, nil
}
