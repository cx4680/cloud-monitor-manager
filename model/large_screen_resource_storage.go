package model

import "time"

type LargeScreenResourceStorage struct {
	ResourceId string    `gorm:"column:resource_id" json:"resource_id"` //资源id
	Type       string    `gorm:"column:type" json:"type"`               //资源类型
	Time       string    `gorm:"column:time" json:"time"`               //时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` //创建时间
	Value      int64     `gorm:"column:value" json:"value"`             //数值
}

func (m *LargeScreenResourceStorage) TableName() string {
	return "large_screen_resource_storage"
}
