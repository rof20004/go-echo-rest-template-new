package user

import (
	"net/http"

	"github.com/labstack/echo"
)

// ListUsers - list all users
func ListUsers(c echo.Context) error {
	users, err := List()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, users)
}

// InsertUser - create a user
func InsertUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := Insert(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, "User saved successfuly")
}

// DeleteUser - delete a user
func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.String(http.StatusOK, "User deleted successfuly")
}
