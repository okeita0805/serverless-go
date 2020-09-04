package controller

import (
	"go-project/internal/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (Controller) GetUserIndex(ctx *gin.Context) {
	user := entity.User{
		ID:   1,
		Name: "aaa",
	}
	users := make(entity.Users, 0)
	users = append(users, user)
	ctx.JSON(http.StatusOK, users)
}
