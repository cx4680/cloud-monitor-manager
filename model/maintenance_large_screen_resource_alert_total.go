package model

type MaintenanceLargeScreenResourceAlertTotal struct {
	Id          int64  `gorm:"column:id" json:"id"`                    //主键id
	AlertLevel  string `gorm:"column:alert_level" json:"alertLevel"`   //告警等级
	AlertNumber int    `gorm:"column:alert_number" json:"alertNumber"` //告警数量
	Region      string `gorm:"column:region" json:"region"`            //区域
}

func (model *MaintenanceLargeScreenResourceAlertTotal) TableName() string {
	return "maintenance_large_screen_resource_alert_total"
}
