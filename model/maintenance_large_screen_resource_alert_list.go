package model

type MaintenanceLargeScreenResourceAlertList struct {
	Id           int64  `gorm:"column:id" json:"id"`                      //主键id
	AlertLevel   string `gorm:"column:alert_level" json:"alertLevel"`     //告警等级
	RuleName     string `gorm:"column:rule_name" json:"ruleName"`         //规则名称
	ResourceId   string `gorm:"column:resource_id" json:"resourceId"`     //资源id
	ResourceType string `gorm:"column:resource_type" json:"resourceType"` //资源名称
	AlertTime    string `gorm:"column:alert_time" json:"alertTime"`       //告警时间
	Region       string `gorm:"column:region" json:"region"`              //区域
}

func (model *MaintenanceLargeScreenResourceAlertList) TableName() string {
	return "maintenance_large_screen_resource_alert_list"
}
