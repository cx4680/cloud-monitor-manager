package form

type MonitorProductParam struct {
	BizIdList []string `form:"bizIdList"`
	Status    uint8    `form:"status"`
}
