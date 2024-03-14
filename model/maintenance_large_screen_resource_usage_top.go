package model

type MaintenanceLargeScreenResourceUsageTop struct {
	Id         int64   `gorm:"column:id" json:"id"`                  //主键id
	ResourceId string  `gorm:"column:resource_id" json:"resourceId"` //资源id
	Type       string  `gorm:"column:type" json:"type"`              //类型，ECS
	Attribute  string  `gorm:"column:attribute" json:"attribute"`    //属性，属性，CPU使用率，内存使用率
	Number     float64 `gorm:"column:number" json:"number"`          //数值
	Unit       string  `gorm:"column:unit" json:"unit"`              //单位
	Region     string  `gorm:"column:region" json:"region"`          //区域
}

func (model *MaintenanceLargeScreenResourceUsageTop) TableName() string {
	return "maintenance_large_screen_resource_usage_top"
}
