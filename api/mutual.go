package api

import (
	"net/http"
	"webproject/dao"
	"webproject/model"
	"webproject/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}
	if dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exists",
		})
		return
	}
	dao.AddUser(req.Username, req.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}
func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	if !dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user does not exists",
		})
		return
	}
	token, err := utils.GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	refreshToken, err := utils.GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
		"token":   token,
		"refresh": refreshToken,
	})
}
func ModifyPassword(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}
	var req model.ModifyPassword
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	if !dao.FindUser(username.(string), req.OldPassword) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "old password is error",
		})
		return
	}
	dao.ModifyPassword(username.(string), req.NewPassword)
	c.JSON(http.StatusOK, gin.H{
		"message": "modify password success",
	})
}
