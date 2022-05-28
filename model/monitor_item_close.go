package model

type MonitorItemClose struct {
	UserId    string `gorm:"column:user_id" json:"userId"`
	ItemBizId string `gorm:"column:item_biz_id" json:"itemBizId"`
}

func (m *MonitorItemClose) TableName() string {
	return "t_monitor_item_close"
}
