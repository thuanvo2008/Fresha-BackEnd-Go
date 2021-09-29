package servicestorage

import (
	"DemoProject/common"
	"DemoProject/modules/service/servicemodel"
	"context"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *servicemodel.ServiceUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
