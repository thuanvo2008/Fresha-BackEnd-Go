package servicestorage

import (
	"DemoProject/common"
	"DemoProject/modules/service/servicemodel"
	"context"
)

func (s *sqlStore) CreateService(ctx context.Context, data *servicemodel.ServiceCreate) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
