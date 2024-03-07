package model

type OperationsLargeScreenResourceRunningStatus struct {
	Id          int64  `gorm:"column:id" json:"id"`                    //主键id
	ResourceId  string `gorm:"column:resource_id" json:"resourceId"`   //资源id
	Status      int    `gorm:"column:status" json:"status"`            //状态，1：正常，0：异常
	Type        string `gorm:"column:type" json:"type"`                //类型，PhysicalServer：物理服务器，NetworkDevice：网络设备，ECS：ECS，RDB：数据库，BMS：裸金属
	Region      string `gorm:"column:region" json:"region"`            //区域
	FailureTime string `gorm:"column:failure_time" json:"failureTime"` //故障时间
	CreateTime  string `gorm:"column:create_time" json:"createTime"`   //创建时间
}

func (model *OperationsLargeScreenResourceRunningStatus) TableName() string {
	return "operations_large_screen_resource_running_status"
}
