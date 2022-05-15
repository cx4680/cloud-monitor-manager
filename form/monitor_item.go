package form

type MonitorItemParam struct {
	ProductBizId string `form:"productBizId"`
	OsType       string `form:"osType"`
	Display      string `form:"display"`
}

type MonitorItem struct {
	ProductAbbreviation string `gorm:"column:product_abbreviation" json:"ProductAbbreviation"`
	Expression          string `gorm:"column:metric_linux" json:"Expression"`
	Labels              string `gorm:"column:labels" json:"Labels"`
	Code                string `gorm:"column:metric_name" json:"ItemCode"`
}
