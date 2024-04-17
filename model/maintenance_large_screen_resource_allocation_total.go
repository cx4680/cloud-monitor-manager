package model

type MaintenanceLargeScreenResourceAllocationTotal struct {
	Id         int64   `gorm:"column:id" json:"id"`                 //主键id
	Type       string  `gorm:"column:type" json:"type"`             //类型，ECS：ECS，RDB：数据库，Storage：存储，EIP：网络EIP，middleware：中间件
	Name       string  `gorm:"column:name" json:"name"`             //名称
	Allocation float64 `gorm:"column:allocation" json:"allocation"` //分配量
	Total      float64 `gorm:"column:total" json:"total"`           //总量
	Unit       string  `gorm:"column:unit" json:"unit"`             //单位
	Time       string  `gorm:"column:time" json:"time"`             //时间
	Region     string  `gorm:"column:region" json:"region"`         //区域
}

func (model *MaintenanceLargeScreenResourceAllocationTotal) TableName() string {
	return "maintenance_large_screen_resource_allocation_total"
}
