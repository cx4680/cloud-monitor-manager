package form

type MonitorItemParam struct {
	ProductBizId    string   `form:"productBizId"`
	OsType          string   `form:"osType"`
	Display         string   `form:"display"`
	MonitorItemList []string `form:"monitorItemList"`
}

type MonitorItem struct {
	ProductAbbreviation string `gorm:"column:product_abbreviation" json:"productAbbreviation"`
	Host                string `gorm:"column:host" json:"host"`
	Expression          string `gorm:"column:metric_linux" json:"expression"`
	Labels              string `gorm:"column:labels" json:"labels"`
	Code                string `gorm:"column:metric_name" json:"itemCode"`
}
