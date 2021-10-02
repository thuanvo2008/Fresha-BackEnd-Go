package businesstimestorage

import (
	"DemoProject/modules/business_time/businesstimemodel"
	"context"
	"time"
)

func (s *sqlStore) Find(ctx context.Context, storeID int, startTime,
	endTime *time.Time, dayOfWeek string) (*businesstimemodel.BusinessTime, error) {
	db := s.db
	var result businesstimemodel.BusinessTime
	starttime := startTime.Format("15:04:05")
	endtime := endTime.Format("15:04:05")

	if err := db.Where("store_id = ?", storeID).
		Where("start_time <= ? AND ? <= end_time", starttime, starttime).
		Where("start_time <= ? AND ? <= end_time", endtime, endtime).
		Where("day_of_week = ?", dayOfWeek).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
