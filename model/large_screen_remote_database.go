package model

type LargeScreenRemoteDatabase struct {
	Region string `gorm:"column:region" json:"region"` //区域
	Host   string `gorm:"column:host" json:"host"`     //主机
	Port   string `gorm:"column:port" json:"port"`     //端口
	Db     string `gorm:"column:db" json:"db"`         //数据库
	User   string `gorm:"column:user" json:"user"`     //账号
	Pass   string `gorm:"column:pass" json:"pass"`     //密码
}

func (m *LargeScreenRemoteDatabase) TableName() string {
	return "large_screen_remote_database"
}
