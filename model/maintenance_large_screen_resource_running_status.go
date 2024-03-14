package model

type MaintenanceLargeScreenResourceRunningStatus struct {
	Id           int64  `gorm:"column:id" json:"id"`                      //主键id
	ResourceId   string `gorm:"column:resource_id" json:"resourceId"`     //资源id
	ResourceName string `gorm:"column:resource_name" json:"resourceName"` //资源名称
	Status       int    `gorm:"column:status" json:"status"`              //状态，1：正常，0：异常
	Type         string `gorm:"column:type" json:"type"`                  //类型，类型，Cluster：集群，Server：云服务，Storage：存储，ECS，XGW，SLB，EIP，RDB
	Region       string `gorm:"column:region" json:"region"`              //区域
	FailureTime  string `gorm:"column:failure_time" json:"failureTime"`   //故障时间
	CreateTime   string `gorm:"column:create_time" json:"createTime"`     //创建时间
}

func (model *MaintenanceLargeScreenResourceRunningStatus) TableName() string {
	return "maintenance_large_screen_resource_running_status"
}
