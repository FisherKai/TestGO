package model

import (
	"gorm.io/gorm"
	"time"
)

// ApprovalResource 表示一个需要审批的资源表单模型
type ApprovalResource struct {
	gorm.Model
	ID           int       `json:"id"`            // 资源ID
	Name         string    `json:"name"`          // 资源名称
	Content      string    `json:"content"`       // 内容
	ResourceType string    `json:"resource_type"` // 资源类型
	Creator      string    `json:"creator"`       // 创建人
	Renewal      string    `json:"renewal"`       // 当前更新人
	Status       string    `json:"status"`        // 当前状态 (例如: "待审批", "已批准", "已拒绝")
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`    // 更新时间
}
