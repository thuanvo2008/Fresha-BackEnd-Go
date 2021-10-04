package closeddatestorage

import (
	"DemoProject/modules/closed_date/closeddatemodel"
	"context"
	"time"
)

func (s *sqlStore) Find(ctx context.Context, storeID int, startDateTime,
	endDateTime *time.Time) (*closeddatemodel.ClosedDate, error) {
	db := s.db
	var result closeddatemodel.ClosedDate

	if err := db.Where("store_id = ?", storeID).
		Where("start_date_time <= ? AND ? <= end_date_time", startDateTime, startDateTime).
		Where("start_date_time <= ? AND ? <= end_date_time", endDateTime, endDateTime).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
