package businesstimestorage

import (
	"DemoProject/modules/business_time/businesstimemodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *businesstimemodel.BusinessTime) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
