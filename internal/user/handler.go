package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(router *gin.Engine, dbConnection *gorm.DB) {
	fmt.Println("Init user routing")
}
