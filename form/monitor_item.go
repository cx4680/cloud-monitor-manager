package form

type MonitorItemParam struct {
	BizIdList    []string `form:"bizIdList"`
	ProductBizId string   `form:"productBizId"`
	OsType       string   `form:"osType"`
	Display      string   `form:"display"`
}

type MonitorItem struct {
	ProductAbbreviation string `gorm:"column:product_abbreviation" json:"ProductAbbreviation"`
	Metric              string `gorm:"column:metric" json:"Metric"`
	Labels              string `gorm:"column:labels" json:"Labels"`
}
