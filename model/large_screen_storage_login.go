package model

type LargeScreenStorageLogin struct {
	Region    string `gorm:"column:region" json:"region"`        //区域
	Vendor    string `gorm:"column:vendor" json:"vendor"`        //厂商
	Type      string `gorm:"column:type" json:"type"`            //类型
	Username  string `gorm:"column:username" json:"username"`    //账号
	Password  string `gorm:"column:password" json:"password"`    //密码
	ManageUrl string `gorm:"column:manage_url" json:"manageUrl"` //管理地址
}

func (m *LargeScreenStorageLogin) TableName() string {
	return "large_screen_storage_login"
}
