package staffworkingstorage

import (
	"DemoProject/modules/staff_working/staffworkingmodel"
	"context"
	"time"
)

func (s *sqlStore) Find(ctx context.Context, staffID int, startTime,
	endTime *time.Time) (*staffworkingmodel.StaffWorking, error) {
	db := s.db
	var result staffworkingmodel.StaffWorking

	if err := db.Where("staff_id = ?", staffID).
		Where("start_time <= ? AND ? <= end_time", startTime, startTime).
		Where("start_time <= ? AND ? <= end_time", endTime, endTime).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
