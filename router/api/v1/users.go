package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/akshara-devs/tekd-be/models"
	"github.com/akshara-devs/tekd-be/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// `json:` -> Coming from go std lib
// `binding:` -> Coming from github.com/go-playground/validator/v10 
// `gorm:` -> Coming from gorm
type createUserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
}

type updateUserRequest struct {
	Email string `json:"email" binding:"omitempty,email"`
	Name  string `json:"name"`
}

func CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.JSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user := models.User{
		Email: req.Email,
		Name:  req.Name,
	}

	if err := pkg.DB.Create(&user).Error; err != nil {
		pkg.JSON(c, http.StatusInternalServerError, "failed to create user", nil)
		return
	}

	pkg.JSON(c, http.StatusCreated, "user created", user)
}

func ListUsers(c *gin.Context) {
	var users []models.User

	if err := pkg.DB.Find(&users).Error; err != nil {
		pkg.JSON(c, http.StatusInternalServerError, "failed to list users", nil)
		return
	}

	pkg.JSON(c, http.StatusOK, "users fetched", users)
}

func GetUser(c *gin.Context) {
	user, ok := findUserByID(c)
	if !ok {
		return
	}

	pkg.JSON(c, http.StatusOK, "user fetched", user)
}

func UpdateUser(c *gin.Context) {
	user, ok := findUserByID(c)
	if !ok {
		return
	}

	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.JSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Name != "" {
		user.Name = req.Name
	}

	if err := pkg.DB.Save(&user).Error; err != nil {
		pkg.JSON(c, http.StatusInternalServerError, "failed to update user", nil)
		return
	}

	pkg.JSON(c, http.StatusOK, "user updated", user)
}

func DeleteUser(c *gin.Context) {
	user, ok := findUserByID(c)
	if !ok {
		return
	}

	if err := pkg.DB.Delete(&user).Error; err != nil {
		pkg.JSON(c, http.StatusInternalServerError, "failed to delete user", nil)
		return
	}

	pkg.JSON(c, http.StatusOK, "user deleted", nil)
}

func findUserByID(c *gin.Context) (models.User, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		pkg.JSON(c, http.StatusBadRequest, "invalid user id", nil)
		return models.User{}, false
	}

	var user models.User
	if err := pkg.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkg.JSON(c, http.StatusNotFound, "user not found", nil)
			return models.User{}, false
		}

		pkg.JSON(c, http.StatusInternalServerError, "failed to fetch user", nil)
		return models.User{}, false
	}

	return user, true
}
