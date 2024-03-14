package service

import "code.cestc.cn/ccos-ops/cloud-monitor-manager/form"

type LargeScreenService struct{}

func NewLargeScreenService() *LargeScreenService {
	return &LargeScreenService{}
}

func (s *LargeScreenService) Tags() (*form.LargeScreenResourceOverview, error) {
	return nil, nil
}

func (s *LargeScreenService) ResourceOverview(tag string) (*form.LargeScreenResourceOverview, error) {
	return nil, nil
}

func (s *LargeScreenService) ResourceAlert(tag string) (*form.LargeScreenResourceOverview, error) {
	return nil, nil
}

func (s *LargeScreenService) ResourceEcs(tag string) (*form.LargeScreenResourceOverview, error) {
	return nil, nil
}

func (s *LargeScreenService) ResourceEip(tag string) (*form.LargeScreenResourceOverview, error) {
	return nil, nil
}

func (s *LargeScreenService) ResourceRdb(tag string) (*form.LargeScreenResourceOverview, error) {
	return nil, nil
}

func (s *LargeScreenService) ResourceStorage(tag string) (*form.LargeScreenResourceOverview, error) {
	return nil, nil
}
