package users

import (
	"github.com/gin-gonic/gin"
	"github.com/roziawilliams/bookstore_users-api/domain/users"
	"github.com/roziawilliams/bookstore_users-api/services"
	"github.com/roziawilliams/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	// Returning a JSON result, with a http statusCreated and the results of the CreateUser
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")

}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "implement me!")
//}
