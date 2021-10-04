package businesstimestorage

import (
	"DemoProject/modules/business_time/businesstimemodel"
	"context"
)

func (s *sqlStore) Find(ctx context.Context, storeID int, dayOfWeek string) (*businesstimemodel.BusinessTime, error) {
	db := s.db
	var result businesstimemodel.BusinessTime

	if err := db.Where("store_id = ?", storeID).
		Where("day_of_week = ?", dayOfWeek).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
