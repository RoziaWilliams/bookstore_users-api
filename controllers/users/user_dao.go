package users

import (
	"fmt"
	"github.com/roziawilliams/bookstore_users-api/domain/users"
	"github.com/roziawilliams/bookstore_users-api/utils/errors"
)

func (user User) Get() *errors.RestErr {
	i := users.User{Id: 125}
	fmt.Println(i)
	return nil
}

func (user User) Save() *errors.RestErr {

	return nil
}
