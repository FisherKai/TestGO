package controller

import (
	"demo_hubu_backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetComprehensiveViewList(c *gin.Context, db *gorm.DB) {
	var item []model.VirtualResource
	db.Find(&item)
	c.JSON(http.StatusOK, item)
}
