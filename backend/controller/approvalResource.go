package controller

import (
	"demo_hubu_backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// ContentCreate 创建表单
func ContentCreate(c *gin.Context, db *gorm.DB) error {
	var approval model.ApprovalResource
	userName, exists := c.Get("user_name")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取当前登录信息失败"})
	}
	approval.Creator = userName.(string)
	if err := c.ShouldBindJSON(&approval); err == nil {
		db.Create(&approval)
		c.JSON(http.StatusCreated, approval)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return nil
}

// ContentUpdate 更新表单
func ContentUpdate(c *gin.Context, db *gorm.DB) error {
	id := c.Param("id")
	var approval model.ApprovalResource
	userName, exists := c.Get("user_name")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取当前登录信息失败"})
	}
	approval.Renewal = userName.(string)
	if err := db.First(&approval, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	}

	if err := c.ShouldBindJSON(&approval); err == nil {
		db.Save(&approval)
		c.JSON(http.StatusOK, approval)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return nil
}

// GetContentList 查询表单list
func GetContentList(c *gin.Context, db *gorm.DB) {
	var approval []model.ApprovalResource
	db.Find(&approval)
	c.JSON(http.StatusOK, approval)
}
