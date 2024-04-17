package model

type OperationsLargeScreenResourceUsageTop struct {
	Id         int64   `gorm:"column:id" json:"id"`                  //主键id
	ResourceId string  `gorm:"column:resource_id" json:"resourceId"` //资源id
	Type       string  `gorm:"column:type" json:"type"`              //类型，ECS：应用云服务器，RDB：数据库，EIP：网络EIP
	Attribute  string  `gorm:"column:attribute" json:"attribute"`    //属性，cpu：cpu，memory：内存，connections：连接数
	Number     float64 `gorm:"column:number" json:"number"`          //数值
	Unit       string  `gorm:"column:unit" json:"unit"`              //单位
	Region     string  `gorm:"column:region" json:"region"`          //区域
}

func (model *OperationsLargeScreenResourceUsageTop) TableName() string {
	return "operations_large_screen_resource_usage_top"
}
