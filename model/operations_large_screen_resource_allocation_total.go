package model

type OperationsLargeScreenResourceAllocationTotal struct {
	Id         int64   `gorm:"column:id" json:"id"`                 //主键id
	Type       string  `gorm:"column:type" json:"type"`             //类型，ECS：ECS，RDB：数据库，Storage：存储，EIP：网络EIP，middleware：中间件
	Attribute  string  `gorm:"column:attribute" json:"attribute"`   //属性，upstream：出网带宽，downstream：入网带宽
	Allocation float64 `gorm:"column:allocation" json:"allocation"` //分配量
	Total      float64 `gorm:"column:total" json:"total"`           //总量
	Unit       string  `gorm:"column:unit" json:"unit"`             //单位
	Region     string  `gorm:"column:region" json:"region"`         //区域
}

func (model *OperationsLargeScreenResourceAllocationTotal) TableName() string {
	return "operations_large_screen_resource_allocation_total"
}
