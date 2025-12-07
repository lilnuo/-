package api

import (
	"fmt"
	"net/http"
	"time"
	"webproject/dao"
	"webproject/middleware"
	"webproject/model"
	"webproject/utils"

	"github.com/gin-gonic/gin"
)

func Ping1(c *gin.Context) {
	fmt.Println("ping1:正在处理核心逻辑")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	fmt.Println("over")
}

func Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exists",
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
			"message": "user not found",
		})
		return
	}
	token, err := utils.MakeToken(req.Username, time.Now().Add(10*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token generate error,internal error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "log in",
	})
}
func InitRouter_gin() {
	r := gin.Default()
	r.GET("/ping", middleware.Example1(), middleware.Example2(), Ping1)
	r.POST("login", Login)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
